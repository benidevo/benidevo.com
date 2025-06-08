package main

import (
	"log"

	"github.com/benidevo/website/internal/app"
	"github.com/benidevo/website/internal/config"
)

func main() {
	cfg := config.NewSettings()
	appInstance := app.New(cfg)

	if err := appInstance.Run(); err != nil {
		log.Fatalf("Failed to run the application: %v", err)
	}

	appInstance.WaitForShutdown()
}
