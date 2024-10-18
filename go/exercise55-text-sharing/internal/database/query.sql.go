// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package database

import (
	"context"
)

const getAllSnippets = `-- name: GetAllSnippets :many

SELECT id, text FROM snippets
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
		if err := rows.Scan(&i.ID, &i.Text); err != nil {
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
