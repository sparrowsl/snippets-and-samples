package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type URL struct {
	Id       string `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

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
