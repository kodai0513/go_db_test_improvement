-- +goose Up
CREATE TABLE bundle_offers (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    price_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE bundle_offers;
