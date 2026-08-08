-- +goose Up
CREATE TABLE refunds (
    id UUID PRIMARY KEY,
    payment_id UUID NOT NULL,
    amount_yen INTEGER NOT NULL,
    reason TEXT NOT NULL,
    refunded_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE refunds;
