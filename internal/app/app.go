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
	server  *http.Server
	sigChan chan os.Signal
	done    chan bool
}

// Returns a configured instance of http.Server
func NewApp() *App {
	tmpl := ParseTemplates()

	mux := NewRouter(tmpl)

	// Server config
	server := &http.Server{
		Addr:         ":3001",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		// ErrorLog:     *log.Logger,  Optionally provide your own error logger
	}

	return &App{
		server:  server,
		sigChan: make(chan os.Signal, 1),
		done:    make(chan bool, 1),
	}
}

// Start the server and wait for shutdown
func (a *App) Run() {
	// Shutdown when signal received
	go a.waitForShutdown()

	a.start()

	if success := <-a.done; success {
		log.Println("Server shutdown successfully")
	}
}

func (a *App) start() {
	log.Printf("Listening on port %s", a.server.Addr)

	err := a.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}

func (a *App) waitForShutdown() {
	// Listen for interrupt or terminate signals

	signal.Notify(a.sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("Received %s, commencing graceful shutdown", <-a.sigChan)

	tc, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.server.Shutdown(tc); err != nil {
		log.Printf("Error shutting down server: %s", err)
		a.done <- false
		return
	}

	a.done <- true
}
