// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
)

const createSnippet = `-- name: CreateSnippet :one
INSERT INTO snippets (slug, text)
VALUES (?, ?)
RETURNING id, text, created_at, slug
`

type CreateSnippetParams struct {
	Slug string `db:"slug" json:"slug"`
	Text string `db:"text" json:"text"`
}

func (q *Queries) CreateSnippet(ctx context.Context, arg CreateSnippetParams) (Snippet, error) {
	row := q.db.QueryRowContext(ctx, createSnippet, arg.Slug, arg.Text)
	var i Snippet
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.CreatedAt,
		&i.Slug,
	)
	return i, err
}

const getAllSnippets = `-- name: GetAllSnippets :many

SELECT id, text, created_at, slug FROM snippets
ORDER BY created_at DESC
`

// -- name: GetOneURL :one
// SELECT * FROM urls
// WHERE short_url = ?
// LIMIT 1;
func (q *Queries) GetAllSnippets(ctx context.Context) ([]Snippet, error) {
	rows, err := q.db.QueryContext(ctx, getAllSnippets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Snippet{}
	for rows.Next() {
		var i Snippet
		if err := rows.Scan(
			&i.ID,
			&i.Text,
			&i.CreatedAt,
			&i.Slug,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}