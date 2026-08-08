-- +goose Up
CREATE TABLE vendor_contracts (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE vendor_contracts;
