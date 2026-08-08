-- +goose Up
CREATE TABLE purchase_orders (
    id UUID PRIMARY KEY,
    supplier_id UUID NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE purchase_orders;
