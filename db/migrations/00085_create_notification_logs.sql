-- +goose Up
CREATE TABLE notification_logs (
    id UUID PRIMARY KEY,
    notification_id UUID NOT NULL,
    status TEXT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE notification_logs;
