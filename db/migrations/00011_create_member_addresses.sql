-- +goose Up
CREATE TABLE member_addresses (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    postal_code TEXT NOT NULL,
    prefecture TEXT NOT NULL,
    city TEXT NOT NULL,
    line1 TEXT NOT NULL,
    line2 TEXT,
    is_default BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_addresses;
