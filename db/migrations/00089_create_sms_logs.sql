-- +goose Up
CREATE TABLE sms_logs (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    body TEXT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE sms_logs;
