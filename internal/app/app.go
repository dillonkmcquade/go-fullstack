package app

import (
	"net/http"
	"time"
)

// Returns a configured instance of http.Server
func NewApp() *http.Server {
	tmpl := ParseTemplates()
	mux := NewRouter(tmpl)

	// Server config
	return &http.Server{
		Addr:         ":3001",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		// ErrorLog:     *log.Logger,  Optionally provide your own error logger
	}
}
