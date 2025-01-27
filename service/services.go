package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/anandtiwari11/event-trigger/constants"
	daointerface "github.com/anandtiwari11/event-trigger/daoInterface"
	"github.com/anandtiwari11/event-trigger/initializers"
	"github.com/anandtiwari11/event-trigger/models"
	"github.com/go-co-op/gocron/v2"
)

type TriggerService struct {
	TriggerDao daointerface.ITriggerDao
}

func NewTriggerService(triggerDao daointerface.ITriggerDao) *TriggerService {
	return &TriggerService{
		TriggerDao: triggerDao,
	}
}

func (triggerService *TriggerService) CreateNewEvent(event *models.Event) error {
	return triggerService.TriggerDao.CreateNewEvent(event)
}

func (triggerService *TriggerService) MoveToArchive(event *models.Event) error {
	return triggerService.TriggerDao.MoveToArchive(event)
}

func (triggerService *TriggerService) DeleteFromArchive(event *models.Event) error {
	return triggerService.TriggerDao.DeleteFromArchive(event)
}

func (triggerService *TriggerService) CreateNewTrigger(trigger *models.Trigger) error {
	return triggerService.TriggerDao.CreateNewTrigger(trigger)
}

func (triggerService *TriggerService) DeleteTrigger(trigger *models.Trigger) error {
	return triggerService.TriggerDao.DeleteTrigger(trigger)
}
func (triggerService *TriggerService) GetAllEvents() (*[]models.Event, error) {
	return triggerService.TriggerDao.GetAllEvents()
}

func (triggerService *TriggerService) UpdateEvent(updatedEvent *models.Event) error {
	return triggerService.TriggerDao.UpdateEvent(updatedEvent)
}

func (triggerService *TriggerService) DeleteEvent(input *models.Event) error {
	return triggerService.TriggerDao.DeleteEvent(input)
}

func (triggerService *TriggerService) GetAllTriggers() (*[]models.Trigger, error) {
	return triggerService.TriggerDao.GetAllTriggers()
}
func (triggerService *TriggerService) UpdateTrigger(updatedTrigger *models.Trigger) error {
	return triggerService.TriggerDao.UpdateTrigger(updatedTrigger)
}

func (triggerService *TriggerService) ProcessScheduledTrigger(trigger *models.Trigger) error {
	if trigger.IsRecurring {
		_, err := initializers.Scheduler.NewJob(
			gocron.DurationJob(
				time.Duration(trigger.Interval)*time.Minute,
			),
			gocron.NewTask(
				func() {
					Event := models.Event{
						Message:       trigger.Message,
						ExecutionTime: time.Now(),
						State:         constants.ACTIVE,
					}
					triggerService.TriggerDao.CreateNewEvent(&Event)
					log.Println(constants.EVENT_CREATED_SUCCESFULLY_FOR_TRIGGER, trigger, Event)
				},
			),
			gocron.WithName("Event Creation"),
		)
		if err != nil {
            log.Println(constants.ERROR_WHILE_CREATING_EVENT)
		}
	} else {
         _, err := initializers.Scheduler.NewJob(
            gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(trigger.ExecutionTime)),
            gocron.NewTask(
                func() {
                    Event := models.Event {
                        Message:       trigger.Message,
						ExecutionTime: time.Now(),
						State:         constants.ACTIVE,
                    }
                    triggerService.TriggerDao.CreateNewEvent(&Event)
					log.Println(constants.EVENT_CREATED_SUCCESFULLY_FOR_TRIGGER, trigger, "where the event is ", Event)
                },
            ),
            gocron.WithName("Event Creation"),
         )
         if err != nil {
            log.Println(constants.ERROR_WHILE_CREATING_EVENT, trigger)
         }
    }
	return nil
}

func (triggerService *TriggerService) ProcessAPITrigger(input *models.Trigger) error {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	var payloadJSON interface{}
	if err := json.Unmarshal(input.Payload, &payloadJSON); err != nil {
		log.Println("error, Invalid payload format, ", "details : ", input.Payload)
		return err
	}
	request, err := http.NewRequest(http.MethodPost, input.Endpoint, strings.NewReader(string(input.Payload)))
	if err != nil {
		log.Println("error while getting response from the api where url is", input.Endpoint, " and payload is ", input.Payload)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Println("Failed to call URL ", input.Endpoint)
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read API response")
	}
	log.Printf("API Response: %s", string(body))

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		log.Println("API call returned non-2xx status code")
		return err
	}
	event := models.Event{
		ExecutionTime: time.Now(),
		Message:       string(input.Payload),
		State:         constants.ACTIVE,
	}
	log.Printf("Attempting to store API response: %+v", &event)
	if err := triggerService.TriggerDao.CreateNewEvent(&event); err != nil {
		log.Println("Failed to store API response for event ", event)
		return err
	}
	return nil
}