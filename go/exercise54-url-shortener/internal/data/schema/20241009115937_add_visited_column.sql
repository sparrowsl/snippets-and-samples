-- +goose Up
-- +goose StatementBegin
ALTER TABLE urls
ADD COLUMN visited INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE urls
DROP COLUMN visited;
-- +goose StatementEnd
