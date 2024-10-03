package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"breve/internal/database"
	"github.com/go-chi/chi/v5"
)

type URL struct {
	Id       string `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

var db = make([]URL, 0)

func (app *application) allURLs(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	urls, _ := app.db.AllURLs(ctx)

	if len(urls) == 0 {
		urls = []database.Url{}
	}

	toJSON(writer, http.StatusOK, map[string]any{"urls": urls})
}

func (app *application) getOneURL(writer http.ResponseWriter, request *http.Request) {
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

func (app *application) createURL(writer http.ResponseWriter, request *http.Request) {
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
