-- +goose Up
CREATE TABLE stock_movements (
    id UUID PRIMARY KEY,
    inventory_id UUID NOT NULL,
    quantity_delta INTEGER NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE stock_movements;
