-- +goose Up
CREATE TABLE tax_rates (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    rate INTEGER NOT NULL
);

-- +goose Down
DROP TABLE tax_rates;
