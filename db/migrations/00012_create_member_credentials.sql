-- +goose Up
CREATE TABLE member_credentials (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    password_hash TEXT NOT NULL,
    last_changed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_credentials;
