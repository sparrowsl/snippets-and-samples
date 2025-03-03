package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"breve/internal/database"
)

type URL struct {
	Id       string `json:"id"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

func (app *application) getURLs(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	urls, _ := app.db.AllURLs(ctx)

	if len(urls) == 0 {
		urls = []database.Url{}
	}

	toJSON(writer, http.StatusOK, map[string]any{"urls": urls})
}

func (app *application) getURL(writer http.ResponseWriter, request *http.Request) {
	paramId := chi.URLParam(request, "id")

	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	url, err := app.db.GetOneURL(ctx, paramId)
	if err != nil {
		toJSON(writer, http.StatusNotFound, map[string]any{"error": err.Error()})
		return
	}

	toJSON(writer, http.StatusOK, map[string]any{"url": url})
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

	// validate if url is not empty
	if strings.TrimSpace(payload.LongURL) == "" {
		toJSON(writer, http.StatusBadRequest, map[string]any{"error": "url is required"})
		return
	}

	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	url, err := app.db.CreateURL(ctx, database.CreateURLParams{
		ShortUrl: generateShortURL(),
		LongUrl:  payload.LongURL,
	})
	if err != nil {
		toJSON(writer, http.StatusInternalServerError, map[string]any{"error": err.Error()})
		return
	}

	toJSON(writer, http.StatusOK, map[string]any{"url": url})
}

func (app *application) updateVisited(writer http.ResponseWriter, request *http.Request) {
	paramId := chi.URLParam(request, "id")

	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	url, err := app.db.GetOneURL(ctx, paramId)
	if err != nil {
		toJSON(writer, http.StatusNotFound, map[string]any{"error": err.Error()})
		return
	}

	err = app.db.UpdateVisited(ctx, database.UpdateVisitedParams{
		Visited:  url.Visited + 1,
		ShortUrl: paramId,
	})
	if err != nil {
		toJSON(writer, http.StatusInternalServerError, map[string]any{"error": err})
		return
	}
}
