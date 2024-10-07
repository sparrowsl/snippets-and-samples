package main

import (
	"fmt"
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

	router.Get("/urls", app.getURLs)
	router.Get("/urls/{id}", app.getURL)
	router.Post("/urls", app.createURL)

	router.Get("/urls/{id}/stats", func(w http.ResponseWriter, r *http.Request) {
		message := fmt.Sprintf("stats for %s", chi.URLParam(r, "id"))
		toJSON(w, http.StatusOK, map[string]any{"stats": message})
	})

	return router
}
