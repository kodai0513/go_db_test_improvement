-- +goose Up
CREATE TABLE media_assets (
    id UUID PRIMARY KEY,
    url TEXT NOT NULL,
    media_type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE media_assets;
