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
	router.Use(middleware.StripSlashes)

	router.Get("/urls", app.allURLs)
	router.Get("/urls/{id}", app.getOneURL)
	router.Post("/urls", app.createURL)

	return router
}
