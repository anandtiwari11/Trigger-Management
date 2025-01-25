package service

import (
	daointerface "github.com/anandtiwari11/event-trigger/daoInterface"
	"github.com/anandtiwari11/event-trigger/models"
)


type TriggerService struct {
	TriggerDao daointerface.ITriggerDao
}

func NewUserService(TriggerDao daointerface.ITriggerDao) *TriggerService {
	return &TriggerService{
		TriggerDao : TriggerDao,
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
	return triggerService.UpdateTrigger(updatedTrigger)
}