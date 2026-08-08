-- +goose Up
CREATE TABLE member_point_accounts (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL UNIQUE,
    balance INTEGER NOT NULL DEFAULT 0,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_point_accounts;
