-- +goose Up
CREATE TABLE content_translations (
    id UUID PRIMARY KEY,
    page_id UUID NOT NULL,
    locale TEXT NOT NULL,
    translated_body TEXT NOT NULL
);

-- +goose Down
DROP TABLE content_translations;
