-- +goose Up
CREATE TABLE tickets (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    subject TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE tickets;
