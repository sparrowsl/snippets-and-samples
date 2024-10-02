package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.StripSlashes)

	router.Get("/", index)
	router.Post("/generate", generateURL)

	http.ListenAndServe(":5000", router)
}

func index(writer http.ResponseWriter, request *http.Request) {
	toJSON(writer, http.StatusOK, map[string]any{"message": "Hello World"})
}

func generateURL(writer http.ResponseWriter, request *http.Request) {
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

	fmt.Println(payload)
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
