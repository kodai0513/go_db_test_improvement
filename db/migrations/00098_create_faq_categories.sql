-- +goose Up
CREATE TABLE faq_categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE faq_categories;
