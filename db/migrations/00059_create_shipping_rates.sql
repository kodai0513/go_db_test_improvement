-- +goose Up
CREATE TABLE shipping_rates (
    id UUID PRIMARY KEY,
    carrier_id UUID NOT NULL,
    zone_name TEXT NOT NULL,
    price_yen INTEGER NOT NULL
);

-- +goose Down
DROP TABLE shipping_rates;
