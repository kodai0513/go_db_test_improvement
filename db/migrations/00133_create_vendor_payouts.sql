-- +goose Up
CREATE TABLE vendor_payouts (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    amount_yen INTEGER NOT NULL,
    paid_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE vendor_payouts;
