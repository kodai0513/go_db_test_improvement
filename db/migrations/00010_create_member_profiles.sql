-- +goose Up
CREATE TABLE member_profiles (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    birth_date DATE,
    gender TEXT,
    bio TEXT,
    avatar_url TEXT,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_profiles;
