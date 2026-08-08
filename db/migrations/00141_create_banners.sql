-- +goose Up
CREATE TABLE banners (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    image_url TEXT NOT NULL,
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE banners;
