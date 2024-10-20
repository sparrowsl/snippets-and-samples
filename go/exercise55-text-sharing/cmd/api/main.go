package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"text-sharing/internal/api"
	"text-sharing/internal/api/handlers"
	"text-sharing/internal/database"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./data-sharing.db")
	if err != nil {
		slog.Default().Error(err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Default().Error(err.Error())
		return
	}
	slog.Default().Info("connected to db successfully...")

	app := handlers.NewApp(database.New(db))
	router := api.NewRouter(app)

	serv := http.Server{
		Addr:        ":5000",
		Handler:     router,
		ReadTimeout: time.Second * 5,
	}

	if err := serv.ListenAndServe(); err != nil {
		slog.Default().Error(err.Error())
	}
}
