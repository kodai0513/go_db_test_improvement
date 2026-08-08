-- +goose Up
CREATE TABLE notification_batches (
    id UUID PRIMARY KEY,
    template_id UUID NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE notification_batches;
