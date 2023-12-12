package app

import (
	"go-fullstack/internal/handlers"
	mw "go-fullstack/internal/middleware"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Returns a completed *chi.Mux with all middleware and routes registered
func NewRouter(tmpl *template.Template) *chi.Mux {
	// Templates are injected here for use by the http HandlerFuncs

	app := chi.NewRouter()

	// Sensible middleware included by default:
	// Check https://github.com/go-chi/chi/tree/master/middleware for other useful middleware
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	// Register your custom middlewares:
	app.Use(mw.Cors)

	// Serve static assets like images, js, and css in ./public:
	app.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	// Register your routes & handler functions here:
	app.Handle("/", handlers.NewIndexHandler(tmpl))
	app.Get("/ping", handlers.Healthcheck)

	return app
}
