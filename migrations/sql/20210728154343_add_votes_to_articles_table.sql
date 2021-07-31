-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF EXISTS articles
    ADD COLUMN IF NOT EXISTS up_vote int,
    ADD COLUMN IF NOT EXISTS down_vote int;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE articles 
    DROP COLUMN IF EXISTS up_vote,
    DROP COLUMN IF EXISTS down_vote;

-- +goose StatementEnd
