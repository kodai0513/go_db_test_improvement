-- +goose Up
CREATE TABLE member_preferences (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    language TEXT NOT NULL DEFAULT 'ja',
    newsletter_opt_in BOOLEAN NOT NULL DEFAULT true,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_preferences;
