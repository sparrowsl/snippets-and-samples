package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"text-sharing/internal/database"
	"text-sharing/internal/server"
)

func (app *application) getAllSnippets(writer http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	snippets, err := app.db.GetAllSnippets(ctx)
	if err != nil {
		server.WriteJSON(writer, http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	server.WriteJSON(writer, http.StatusOK, map[string]any{"snippets": snippets})
}

func (app application) createSnippet(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var input struct {
		Text string `json:"text"`
	}
	if err := decoder.Decode(&input); err != nil {
		server.WriteJSON(writer, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	if strings.TrimSpace(input.Text) == "" {
		server.WriteJSON(writer, http.StatusBadRequest, map[string]any{"error": "text is empty!!"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	hash := md5.New().Sum([]byte(input.Text))
	snippet, err := app.db.CreateSnippet(ctx, database.CreateSnippetParams{
		Slug: hex.EncodeToString(hash),
		Text: input.Text,
	})
	if err != nil {
		server.WriteJSON(writer, http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	server.WriteJSON(writer, http.StatusOK, map[string]any{"snippet": snippet})
}
