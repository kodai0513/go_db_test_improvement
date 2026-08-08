-- +goose Up
CREATE TABLE return_requests (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    reason TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE return_requests;
