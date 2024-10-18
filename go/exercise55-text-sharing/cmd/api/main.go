package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"text-sharing/internal/server"

	_ "modernc.org/sqlite"
)

type application struct {
	routes http.Handler
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

	server := http.Server{
		Addr:        ":5000",
		Handler:     server.Routes(),
		ReadTimeout: time.Second * 5,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
