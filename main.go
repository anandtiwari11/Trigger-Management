package main

import (
	"context"
	"log"

	jobs "github.com/anandtiwari11/event-trigger/cronJobs"
	"github.com/anandtiwari11/event-trigger/dao"
	dependencyinjection "github.com/anandtiwari11/event-trigger/dependencyInjection"
	"github.com/anandtiwari11/event-trigger/initializers"
)

func main() {
	initializers.ConnectDB()
	app := dependencyinjection.LoadDependencies()
	ctx := context.Background()
	triggerDAO := &dao.TriggerDaoImpl{}
	eventLifecycleJob := &jobs.EventLifecycleJob{
		TriggerDAO: triggerDAO,
	}
	scheduledTriggerJob := &jobs.ScheduledTriggerJob{
		TriggerDAO: triggerDAO,
	}
	go eventLifecycleJob.Run()
	go scheduledTriggerJob.Run()
	startErr := app.Start(ctx)
	if startErr != nil {
		log.Fatalf("Error starting application: %v", startErr)
	}
	defer func() {
		if err := app.Stop(ctx); err != nil {
			log.Fatalf("Error stopping application: %v", err)
		}
	}()
	select {}
}
