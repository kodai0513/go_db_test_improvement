-- +goose Up
CREATE TABLE payment_providers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE payment_providers;
