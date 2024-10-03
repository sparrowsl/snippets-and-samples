package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func allURLs(writer http.ResponseWriter, request *http.Request) {
	toJSON(writer, http.StatusOK, map[string]any{"urls": db})
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
