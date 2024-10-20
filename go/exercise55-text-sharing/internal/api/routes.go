package api

import (
	"net/http"

	"text-sharing/internal/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(app *handlers.Application) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)

	router.Get("/", app.GetAllSnippets)
	router.Get("/{slug}", app.GetSnippet)
	router.Post("/", app.CreateSnippet)
	router.Patch("/{slug}", app.UpdateSnippet)

	return router
}
