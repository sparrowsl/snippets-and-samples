-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS snippets(
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  text TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS snippets;
-- +goose StatementEnd
