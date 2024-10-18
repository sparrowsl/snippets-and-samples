package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
		server.WriteJSON(writer, http.StatusBadRequest, map[string]any{"error": err})
		return
	}

	fmt.Println(input)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	snippets, err := app.db.GetAllSnippets(ctx)
	if err != nil {
		server.WriteJSON(writer, http.StatusInternalServerError, map[string]any{"error": err})
		return
	}

	server.WriteJSON(writer, http.StatusOK, map[string]any{"snippets": snippets})
}
