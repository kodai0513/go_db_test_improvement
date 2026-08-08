-- +goose Up
CREATE TABLE price_histories (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    price_yen INTEGER NOT NULL,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE price_histories;
