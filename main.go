package main

import (
	"context"
	"log"
	dependencyinjection "github.com/anandtiwari11/event-trigger/dependencyInjection"
)

func main() {
	app := dependencyinjection.LoadDependencies()
	ctx := context.Background()
	startErr := app.Start(ctx)
	if startErr != nil {
		log.Fatalf("Error starting application: %v", startErr)
	}
	defer func() {
		if err := app.Stop(ctx); err != nil {
			log.Fatalf("Error stopping application: %v", err)
		}
	}()
	<-app.Done()
}
