-- +goose Up
CREATE TABLE carriers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE carriers;
