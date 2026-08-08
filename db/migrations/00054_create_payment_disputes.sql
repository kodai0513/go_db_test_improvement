-- +goose Up
CREATE TABLE payment_disputes (
    id UUID PRIMARY KEY,
    payment_id UUID NOT NULL,
    reason TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE payment_disputes;
