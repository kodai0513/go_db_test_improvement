-- +goose Up
CREATE TABLE order_cancellations (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    reason TEXT NOT NULL,
    cancelled_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE order_cancellations;
