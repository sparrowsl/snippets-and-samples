package handlers

import (
	"context"

	"text-sharing/internal/database"
)

type Application struct {
	DB  *database.Queries
	ctx context.Context
}

func NewApp(db *database.Queries) *Application {
	return &Application{
		DB:  db,
		ctx: context.Background(),
	}
}
