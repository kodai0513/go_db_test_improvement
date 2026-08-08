-- +goose Up
CREATE TABLE satisfaction_surveys (
    id UUID PRIMARY KEY,
    ticket_id UUID NOT NULL,
    score INTEGER NOT NULL,
    comment TEXT,
    submitted_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE satisfaction_surveys;
