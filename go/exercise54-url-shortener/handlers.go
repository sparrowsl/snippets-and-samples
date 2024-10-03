package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var db = make([]URL, 0)

func allURLs(writer http.ResponseWriter, request *http.Request) {
	toJSON(writer, http.StatusOK, map[string]any{"urls": db})
}

func getOneURL(writer http.ResponseWriter, request *http.Request) {
	paramId := chi.URLParam(request, "id")

	// search for the URL in db
	for _, u := range db {
		if u.Id == paramId {
			toJSON(writer, http.StatusOK, map[string]any{"url": u})
			return
		}
	}

	toJSON(writer, http.StatusNotFound, nil)
}

func createURL(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var payload struct {
		LongURL string `json:"url"`
	}

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