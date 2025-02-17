package main

import (
	"log"

	"github.com/khapaev/solid-octo-meme/config"
	"github.com/khapaev/solid-octo-meme/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
