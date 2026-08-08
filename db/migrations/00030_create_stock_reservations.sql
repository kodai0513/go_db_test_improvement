-- +goose Up
CREATE TABLE stock_reservations (
    id UUID PRIMARY KEY,
    inventory_id UUID NOT NULL,
    order_id UUID NOT NULL,
    quantity INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE stock_reservations;
