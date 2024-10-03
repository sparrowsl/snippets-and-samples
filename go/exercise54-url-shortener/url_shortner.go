package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Get("/urls", allURLs)
	router.Get("/urls/{id}", getOneURL)
	router.Post("/urls", createURL)

	serv := http.Server{
		Handler:     router,
		Addr:        ":5000",
		ReadTimeout: time.Second * 5,
	}

	serv.ListenAndServe()
}
