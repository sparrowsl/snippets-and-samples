package main

import (
	"encoding/json"
	"fmt"
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

func allURLs(writer http.ResponseWriter, request *http.Request) {
	toJSON(writer, http.StatusOK, map[string]any{"message": "Hello World"})
}

func createURL(writer http.ResponseWriter, request *http.Request) {
	type Payload struct {
		LongURL string `json:"url"`
	}

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var payload Payload
	if err := decoder.Decode(&payload); err != nil {
		toJSON(writer, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	u := URL{
		Id:       fmt.Sprintf("%d", len(db)+1),
		LongURL:  payload.LongURL,
		ShortURL: generateShortURL(0),
	}

	db = append(db, u)
	toJSON(writer, http.StatusOK, map[string]any{"url": u})
}

func generateShortURL(length int) string {
	if length == 0 {
		length = 6
	}

	return fmt.Sprintf("abc12300%d", len(db))
}

func toJSON(writer http.ResponseWriter, status int, data map[string]any) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write(result)

	return nil
}
