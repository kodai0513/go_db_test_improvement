-- +goose Up
CREATE TABLE email_logs (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    subject TEXT NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE email_logs;
