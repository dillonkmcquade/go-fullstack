package main

import (
	"context"
	"go-fullstack/internal/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := app.NewApp()

	go func() {
		log.Printf("Listening on port %s", app.Addr)
		err := app.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	// Listen for interrupt or terminate signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Shutdown when signal received
	log.Printf("Received %s, commencing graceful shutdown", <-sigChan)
	tc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Shutdown(tc); err != nil {
		log.Printf("Error shutting down server: %s", err)
	}

	log.Println("Server shutdown successfully")
}
