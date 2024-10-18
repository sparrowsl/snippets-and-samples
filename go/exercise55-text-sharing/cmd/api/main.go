package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"text-sharing/internal/database"

	_ "modernc.org/sqlite"
)

type application struct {
	db *database.Queries
}

func main() {
	db, err := sql.Open("sqlite", "./data-sharing.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("connected to db successfully...")

	app := application{
		db: database.New(db),
	}

	serv := http.Server{
		Addr:        ":5000",
		Handler:     app.routes(),
		ReadTimeout: time.Second * 5,
	}

	if err := serv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
