-- +goose Up
CREATE TABLE order_status_histories (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    status TEXT NOT NULL,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE order_status_histories;
