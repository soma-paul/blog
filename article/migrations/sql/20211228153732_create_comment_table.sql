-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments(
    id          serial PRIMARY KEY          not null,
    uid         int REFERENCES users(id)    not null,
    article_id  int REFERENCES articles(id) not null,
    comment     text,
    created_at  timestamp ,     
    updated_at  timestamp 

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
-- +goose StatementEnd
