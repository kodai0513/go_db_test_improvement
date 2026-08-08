-- +goose Up
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    total_amount_yen INTEGER NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE orders;
