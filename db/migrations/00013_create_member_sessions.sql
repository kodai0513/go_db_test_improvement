-- +goose Up
CREATE TABLE member_sessions (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    session_token TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_sessions;
