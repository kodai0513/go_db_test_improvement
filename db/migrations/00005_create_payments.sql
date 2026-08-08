-- +goose Up
CREATE TABLE payments (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    amount_yen INTEGER NOT NULL,
    method TEXT NOT NULL,
    paid_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE payments;
