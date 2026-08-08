-- +goose Up
CREATE TABLE ticket_status_histories (
    id UUID PRIMARY KEY,
    ticket_id UUID NOT NULL,
    status TEXT NOT NULL,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE ticket_status_histories;
