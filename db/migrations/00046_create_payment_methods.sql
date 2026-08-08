-- +goose Up
CREATE TABLE payment_methods (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    type TEXT NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE payment_methods;
