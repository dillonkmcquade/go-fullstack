package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server         *http.Server
	shutdownSignal chan os.Signal // Signals a shutdown
}

func New() *App {
	tmpl := Templates()

	router := Router(tmpl)

	// Server config
	server := &http.Server{
		Addr:         ":3001",
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		// ErrorLog:     *log.Logger,  Optionally provide your own error logger
	}

	return &App{
		server:         server,
		shutdownSignal: make(chan os.Signal, 1),
	}
}

// Start the server and wait for shutdown
func (a *App) Run() {
	// Shutdown when signal received
	go a.waitForShutdown()

	a.start()
}

func (a *App) start() {
	log.Printf("Listening on port %s", a.server.Addr)

	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("server error: %s", err)
	}
}

func (a *App) waitForShutdown() {
	// Listen for interrupt or terminate signals

	signal.Notify(a.shutdownSignal, syscall.SIGINT, syscall.SIGTERM)
	tc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	select {
	case <-a.shutdownSignal:
		if err := a.server.Shutdown(tc); err != nil {
			log.Fatalf("on server shutdown: %s", err)
		} else {
			log.Println("server shutdown successfully")
		}
		// Add other signals here
	}
}
