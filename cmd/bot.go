package main

import (
	"bot-templates-profi/internal/app"
	"bot-templates-profi/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	cfg := config.LoadConfig()

	app.MustRun(cfg)

}
