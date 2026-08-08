-- +goose Up
CREATE TABLE shipments (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    address TEXT NOT NULL,
    status TEXT NOT NULL,
    shipped_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE shipments;
