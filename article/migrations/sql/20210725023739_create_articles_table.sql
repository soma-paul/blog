-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS articles(
    id          serial PRIMARY KEY          not null,
    title       varchar(70)                 not null,
    description text                        not null,
    uid         int REFERENCES users(id)    not null,
    username    varchar(30)                 not null,
    created_at  timestamp ,      
    updated_at  timestamp

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd
