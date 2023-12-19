package main

import (
	"go-fullstack/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := app.New()
	app.Run()
}
