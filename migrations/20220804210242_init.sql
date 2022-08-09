-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Movie (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    year integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Movie;
-- +goose StatementEnd
