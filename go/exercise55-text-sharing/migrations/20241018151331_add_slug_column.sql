-- +goose Up
-- +goose StatementBegin
ALTER TABLE snippets
ADD COLUMN slug TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE snippets
DROP COLUMN IF EXISTS slug;
-- +goose StatementEnd
