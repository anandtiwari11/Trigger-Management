package dao

import (
	"fmt"
	"log"
	"time"

	"github.com/anandtiwari11/event-trigger/initializers"
	"github.com/anandtiwari11/event-trigger/models"
	"gorm.io/gorm"
)

type TriggerDaoImpl struct{}

func NewTriggerDaoImpl() *TriggerDaoImpl {
	return &TriggerDaoImpl{}
}

func (dao *TriggerDaoImpl) CreateNewTrigger(trigger *models.Trigger) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&trigger).Error; err != nil {
			return fmt.Errorf("failed to create trigger: %v", err)
		}
		return nil
	})
}

func (dao *TriggerDaoImpl) DeleteTrigger(trigger *models.Trigger) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(trigger).Error; err != nil {
			return fmt.Errorf("failed to delete trigger %d: %v", trigger.ID, err)
		}
		return nil
	})
}

func (dao *TriggerDaoImpl) CreateNewEvent(event *models.Event) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&event).Error; err != nil {
			return fmt.Errorf("failed to create event: %v", err)
		}
		return nil
	})
}

func (dao *TriggerDaoImpl) UpdateExecutionTime(trigger *models.Trigger) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		newExecutionTime := time.Now().Add(time.Duration(trigger.Interval) * time.Minute)
		trigger.ExecutionTime = newExecutionTime
		if err := tx.Save(trigger).Error; err != nil {
			return fmt.Errorf("failed to update execution time for trigger %d: %v", trigger.ID, err)
		}
		log.Printf("Updated trigger %d with new execution time: %v", trigger.ID, newExecutionTime)
		return nil
	})
}

func (dao *TriggerDaoImpl) MoveToArchive(event *models.Event) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		event.State = "archived"
		if err := tx.Save(event).Error; err != nil {
			return fmt.Errorf("failed to move event %d to archive: %v", event.ID, err)
		}
		log.Printf("Event with ID %d has been moved to archive", event.ID)
		return nil
	})
}

func (dao *TriggerDaoImpl) DeleteFromArchive(event *models.Event) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(event).Error; err != nil {
			return fmt.Errorf("failed to delete event %d from archive: %v", event.ID, err)
		}
		log.Printf("Event with ID %d has been deleted from the archive", event.ID)
		return nil
	})
}

func (dao *TriggerDaoImpl) FetchAllCurrent() (*[]models.Trigger, error) {
	var triggers []models.Trigger
	err := initializers.DB.Where("execution_time <= ?", time.Now()).Find(&triggers).Error
	if err != nil {
		log.Println("Error While Fetching All Triggers")
		return nil, fmt.Errorf("error fetching current triggers: %v", err)
	}
	return &triggers, nil
}

func (dao *TriggerDaoImpl) FetchTriggerByTriggerId(id uint) (*models.Trigger, error) {
	var trigger models.Trigger
	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Where("id = ?", id).First(&trigger).Error
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching trigger %d: %v", id, err)
	}
	return &trigger, nil
}

func (dao *TriggerDaoImpl) GetAllEvents() (*[]models.Event, error) {
	var events []models.Event
	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&events).Error
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching all events: %v", err)
	}
	return &events, nil
}

func (dao *TriggerDaoImpl) GetAllTriggers() (*[]models.Trigger, error) {
	var triggers []models.Trigger
	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Find(&triggers).Error
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching all triggers: %v", err)
	}
	return &triggers, nil
}

func (dao *TriggerDaoImpl) UpdateEvent(updatedEvent *models.Event) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(updatedEvent).Error; err != nil {
			return fmt.Errorf("failed to update event %d: %v", updatedEvent.ID, err)
		}
		return nil
	})
}

func (dao *TriggerDaoImpl) DeleteEvent(input *models.Event) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", input.ID).Delete(&models.Event{}).Error; err != nil {
			return fmt.Errorf("failed to delete event %d: %v", input.ID, err)
		}
		return nil
	})
}

func (dao *TriggerDaoImpl) FetchAllActive() (*[]models.Event, error) {
	var events []models.Event
	tx := initializers.DB.Begin()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	if err := tx.Where("state = ?", "active").Find(&events).Error; err != nil {
		tx.Rollback() 
		log.Printf("Error fetching active events: %v", err)
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &events, nil
}

func (dao *TriggerDaoImpl) FetchAllArchived() (*[]models.Event, error) {
	var events []models.Event
	tx := initializers.DB.Begin()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	if err := tx.Where("state = ?", "archived").Find(&events).Error; err != nil {
		tx.Rollback() 
		log.Printf("Error fetching archived events: %v", err)
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &events, nil
}

func (dao *TriggerDaoImpl) UpdateTrigger(updatedTrigger *models.Trigger) error {
	return initializers.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(updatedTrigger).Error; err != nil {
			return fmt.Errorf("failed to update event %d: %v", updatedTrigger.ID, err)
		}
		return nil
	})
}