package jobs

import (
	"log"
	"time"

	"github.com/anandtiwari11/event-trigger/constants"
	"github.com/anandtiwari11/event-trigger/dao"
	"github.com/anandtiwari11/event-trigger/models"
)

type ScheduledTriggerJob struct {
	TriggerDAO *dao.TriggerDaoImpl
}

func (job *ScheduledTriggerJob) Run() {
	for range time.Tick(1 * time.Second) {
		triggers, err := job.TriggerDAO.FetchAllCurrent()
		if err != nil {
			log.Printf("Error fetching active triggers: %v", err)
			continue
		}
		for _, trigger := range *triggers {
			event := &models.Event{
				Payload:   trigger.Payload,
				State:     string(constants.ACTIVE),
				Response:  string(trigger.Message),
				Timestamp: time.Now(),
			}
			err := job.TriggerDAO.CreateNewEvent(event)
			if err != nil {
				err := job.TriggerDAO.DeleteTrigger(&trigger)
				log.Printf("Failed to create event for trigger %v: %v", trigger.ID, err)
				continue
			}

			log.Printf("Successfully created event for trigger %v", trigger.ID)

			if trigger.IsRecurring {
				istLocation, err := time.LoadLocation("Asia/Kolkata")
				if err != nil {
					log.Printf("Failed to load IST timezone: %v", err)
					continue
				}
				trigger.ExecutionTime = trigger.ExecutionTime.Add(time.Duration(trigger.Interval)).In(istLocation)
				err = job.TriggerDAO.UpdateExecutionTime(&trigger)
				if err != nil {
					log.Printf("Failed to update execution time for recurring trigger %v: %v", trigger.ID, err)
					continue
				}
				log.Printf("Updated execution time for recurring trigger %v to %v", trigger.ID, trigger.ExecutionTime)
			} else {
				err := job.TriggerDAO.DeleteTrigger(&trigger)
				if err != nil {
					log.Printf("Failed to delete non-recurring trigger %v: %v", trigger.ID, err)
					continue
				}
				log.Printf("Successfully deleted non-recurring trigger %v", trigger.ID)
			}
		}
	}
}

// func (job *ScheduledTriggerJob) Run() {
// 	ctx := context.Background()
// 	for range time.Tick(1 * time.Second) {
// 		result, err := job.RedisClient.ZRangeWithScores(ctx, "active_triggers", 0, 0).Result()
// 		if err != nil {
// 			log.Printf("Error fetching triggers from Redis: %v", err)
// 			continue
// 		}
// 		if len(result) == 0 {
// 			continue
// 		}
// 		triggerID := result[0].Member.(string)
// 		executionTime := time.Unix(int64(result[0].Score), 0)
// 		if executionTime.After(time.Now()) {
// 			continue
// 		}
// 		id, _ := strconv.Atoi(triggerID)
// 		trigger, err := job.TriggerDAO.FetchTriggerByTriggerId(uint(id))
// 		if err != nil {
// 			_ = job.RedisClient.ZRem(ctx, "active_triggers", triggerID).Err()
// 			log.Printf("Error while finding trigger with trigger id %d", id)
// 			continue
// 		}
// 		newEvent := models.Event {
// 			Payload: trigger.Payload,
// 			Timestamp: time.Now(),
// 			State: constants.ACTIVE,
// 			Response: trigger.Message,
// 		}
// 		job.TriggerDAO.CreateNewEvent(&newEvent)
// 		if trigger.IsRecurring {
// 			trigger.ExecutionTime = trigger.ExecutionTime.Add(time.Duration(trigger.Interval) * time.Minute)
// 			if err := job.RedisClient.ZAdd(ctx, "active_triggers", &redis.Z{
// 				Score:  float64(trigger.ExecutionTime.Unix()),
// 				Member: trigger.ID,
// 			}).Err(); err != nil {
// 				log.Printf("Error updating recurring trigger in Redis: %v", err)
// 			}
// 			job.TriggerDAO.UpdateEvent(&newEvent)
// 		} else {
// 			job.TriggerDAO.DeleteTrigger(trigger)
// 			job.RedisClient.ZRem(ctx, "active_triggers", triggerID)
// 		}
// 	}
// }
