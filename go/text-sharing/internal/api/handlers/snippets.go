package handlers

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"text-sharing/internal/database"
)

func (app *Application) GetAllSnippets(writer http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	snippets, err := app.DB.GetAllSnippets(ctx)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]any{"error": err.Error()})
		return
	}

	json.NewEncoder(writer).Encode(map[string]any{"snippets": snippets})
}

func (app Application) CreateSnippet(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var input struct {
		Text string `json:"text"`
	}
	if err := decoder.Decode(&input); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]any{"error": err.Error()})
		return
	}

	if strings.TrimSpace(input.Text) == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]any{"error": "text is empty!!"})
		return
	}

	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	hash := md5.New().Sum([]byte(input.Text))
	slug := hex.EncodeToString(hash)

	snippet, err := app.DB.CreateSnippet(ctx, database.CreateSnippetParams{
		Slug: slug,
		Text: input.Text,
	})
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]any{"error": err})
		return
	}

	json.NewEncoder(writer).Encode(map[string]any{"snippet": snippet})
}

func (app *Application) GetSnippet(writer http.ResponseWriter, request *http.Request) {
	slug := chi.URLParam(request, "slug")

	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	snippet, err := app.DB.GetSnippet(ctx, slug)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			writer.WriteHeader(http.StatusNotFound)
			json.NewEncoder(writer).Encode(map[string]any{"error": "No snippet found!!"})
		default:
			writer.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(writer).Encode(map[string]any{"error": err.Error()})
		}
		return
	}

	json.NewEncoder(writer).Encode(map[string]any{"snippet": snippet})
}

func (app *Application) UpdateSnippet(writer http.ResponseWriter, request *http.Request) {
	slug := chi.URLParam(request, "slug")

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var input struct {
		Text string `json:"text"`
	}
	if err := decoder.Decode(&input); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]any{"error": err.Error()})
		return
	}

	if strings.TrimSpace(input.Text) == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]any{"error": "text is empty!!"})
		return
	}

	ctx, cancel := context.WithTimeout(app.ctx, time.Second*5)
	defer cancel()

	exists, err := app.DB.GetSnippet(ctx, slug)
	if err != nil {
		switch {
		case exists.ID == 0:
			writer.WriteHeader(http.StatusNotFound)
			json.NewEncoder(writer).Encode(map[string]any{"error": "snippet does not exists!!"})
		default:
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(map[string]any{"error": err.Error()})
		}
		return
	}

	err = app.DB.UpdateSnippet(ctx, database.UpdateSnippetParams{
		Text: input.Text,
		Slug: slug,
	})
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]any{"error": err.Error()})
	}
}
