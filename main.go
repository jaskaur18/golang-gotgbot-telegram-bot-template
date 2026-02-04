package main

import (
	"log"

	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/app"
	"github.com/jaskaur18/golang-gotgbot-telegram-bot-template/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it, relying on system env")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app.Run(cfg)
}
