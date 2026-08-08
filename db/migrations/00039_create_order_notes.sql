-- +goose Up
CREATE TABLE order_notes (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    note TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE order_notes;
