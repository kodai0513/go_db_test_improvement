-- +goose Up
CREATE TABLE suppliers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    contact_email TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE suppliers;
