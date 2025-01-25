package serviceinterface


import "github.com/anandtiwari11/event-trigger/models"

type IServiceInterface interface {
	DeleteEvent(input *models.Event) error
	UpdateEvent(updatedEvent *models.Event) error
	GetAllEvents() (*[]models.Event, error)
	CreateNewTrigger(trigger *models.Trigger) error
	DeleteTrigger(trigger *models.Trigger) error
	CreateNewEvent(event *models.Event) error
	MoveToArchive(event *models.Event) error
	DeleteFromArchive(event *models.Event) error
	GetAllTriggers() (*[]models.Trigger, error)
	UpdateTrigger(updatedTrigger *models.Trigger) error
}