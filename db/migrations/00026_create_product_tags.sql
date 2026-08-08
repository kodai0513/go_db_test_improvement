-- +goose Up
CREATE TABLE product_tags (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE product_tags;
