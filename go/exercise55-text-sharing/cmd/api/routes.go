package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)

	router.Get("/", app.getAllSnippets)
	router.Get("/{slug}", app.getSnippet)
	router.Post("/", app.createSnippet)
	router.Patch("/{slug}", app.updateSnippet)

	return router
}
