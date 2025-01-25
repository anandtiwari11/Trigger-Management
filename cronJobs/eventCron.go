package jobs

import (
	"log"
	"time"

	"github.com/anandtiwari11/event-trigger/dao"
)

type EventLifecycleJob struct {
	TriggerDAO *dao.TriggerDaoImpl
}

func (job *EventLifecycleJob) Run() {
	for range time.Tick(1 * time.Minute) {
		log.Printf("Fetching all active events")
		activeEvents, err := job.TriggerDAO.FetchAllActive()
		if err != nil {
			log.Printf("Error fetching active events: %v", err)
			continue
		}
		for _, event := range *activeEvents {
			if time.Since(event.Timestamp) > 2*time.Hour {
				log.Printf("The event with event id %d has been sent to archive", event.ID)
				job.TriggerDAO.MoveToArchive(&event)
			}
		}

		archivedEvents, err := job.TriggerDAO.FetchAllArchived()
		if err != nil {
			log.Printf("Error fetching archived events: %v", err)
			continue
		}
		for _, event := range *archivedEvents {
			if time.Since(event.Timestamp) > 48*time.Hour {
				log.Printf("The event with event id %d has been deleted", event.ID)
				job.TriggerDAO.DeleteFromArchive(&event)
			}
		}
	}
}
