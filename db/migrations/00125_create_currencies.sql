-- +goose Up
CREATE TABLE currencies (
    id UUID PRIMARY KEY,
    code TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE currencies;
