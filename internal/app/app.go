package app

import (
	"net/http"
	"time"
)

func NewApp() *http.Server {
	tmpl := ParseTemplates()
	mux := NewRouter(tmpl)
	return &http.Server{
		Addr:         ":3001",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
}
