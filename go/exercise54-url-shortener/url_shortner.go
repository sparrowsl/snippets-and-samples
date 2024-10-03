package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type URL struct {
	Id       string `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

var db = make([]URL, 0)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Get("/urls", allURLs)
	router.Post("/urls", createURL)

	http.ListenAndServe(":5000", router)
}
