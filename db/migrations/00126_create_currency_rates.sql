-- +goose Up
CREATE TABLE currency_rates (
    id UUID PRIMARY KEY,
    currency_id UUID NOT NULL,
    rate_to_jpy INTEGER NOT NULL,
    effective_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE currency_rates;
