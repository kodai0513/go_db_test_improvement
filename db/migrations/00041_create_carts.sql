-- +goose Up
CREATE TABLE carts (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE carts;
