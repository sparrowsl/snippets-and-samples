-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS urls (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	short_url TEXT NOT NULL UNIQUE,
	long_url TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
