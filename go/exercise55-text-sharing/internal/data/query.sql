-- name: GetSnippet :one
SELECT * FROM snippets
WHERE slug = ?
LIMIT 1;

-- name: GetAllSnippets :many
SELECT * FROM snippets
ORDER BY created_at DESC;

-- name: CreateSnippet :one
INSERT INTO snippets (slug, text)
VALUES (?, ?)
RETURNING *;

-- name: UpdateSnippet :exec
UPDATE snippets
SET text = ?
WHERE slug = ?;
