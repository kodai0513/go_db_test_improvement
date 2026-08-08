-- +goose Up
CREATE TABLE vendor_commissions (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    rate_percent INTEGER NOT NULL
);

-- +goose Down
DROP TABLE vendor_commissions;
