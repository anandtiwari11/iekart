package main

import (
	"context"
	"log"

	dependencyinjection "github.com/anandtiwari11/IEKart-go/dependencyInjection"
	"github.com/anandtiwari11/IEKart-go/initializers"
)

func main() {
	initializers.ConnectDB()
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
	select {}
}