package main

import (
	"log"

	"github.com/nitin-sharma-7991/aihub-backend/internal/config"
	"github.com/nitin-sharma-7991/aihub-backend/internal/logger"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logg, err := logger.New(cfg.App.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	defer logg.Sync()

	logg.Info("Application Started")
	logg.Info("Configuration Loaded")
}
