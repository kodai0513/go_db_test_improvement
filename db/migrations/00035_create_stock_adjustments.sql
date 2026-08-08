-- +goose Up
CREATE TABLE stock_adjustments (
    id UUID PRIMARY KEY,
    inventory_id UUID NOT NULL,
    adjusted_quantity INTEGER NOT NULL,
    note TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE stock_adjustments;
