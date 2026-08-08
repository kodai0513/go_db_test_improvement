-- +goose Up
CREATE TABLE ticket_categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE ticket_categories;
