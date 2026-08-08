-- +goose Up
CREATE TABLE coupons (
    id UUID PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    discount_rate INTEGER NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE coupons;
