package main

import (
	"context"
	"log"
	"sebsegura/umock/internal/api"
	"sebsegura/umock/internal/config"
)

func main() {
	start()
}

func start() {
	config.LoadConfig()

	router, err := api.SetupRouter(context.Background(), "spec.yaml")
	if err != nil {
		log.Fatalf("cannot setup router: %v", err)
	}

	if err = router.Run(); err != nil {
		log.Fatalf("cannot start router: %v", err)
	}
}
