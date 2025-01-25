package daointerface

import "github.com/anandtiwari11/event-trigger/models"

type ITriggerDao interface {
	UpdateTrigger(updatedTrigger *models.Trigger) error
	DeleteEvent(input *models.Event) error
	UpdateEvent(updatedEvent *models.Event) error
	GetAllEvents() (*[]models.Event, error)
	CreateNewTrigger(trigger *models.Trigger) error
	DeleteTrigger(trigger *models.Trigger) error
	CreateNewEvent(event *models.Event) error
	MoveToArchive(event *models.Event) error
	DeleteFromArchive(event *models.Event) error
	FetchAllActive() (*[]models.Event, error)
	UpdateExecutionTime(trigger *models.Trigger) error
	GetAllTriggers() (*[]models.Trigger, error)
}
