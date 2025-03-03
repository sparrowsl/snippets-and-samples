package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	router.Use(middleware.StripSlashes)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		AllowedMethods:   []string{"GET", "POST", "PATCH"},
		AllowCredentials: false,
		MaxAge:           int(time.Second * 5),
	}))

	router.Get("/urls", app.getURLs)
	router.Get("/urls/{id}", app.getURL)
	router.Post("/urls", app.createURL)
	router.Get("/urls/{id}/stats", app.getURL)
	router.Patch("/urls/{id}", app.updateVisited)

	return router
}
