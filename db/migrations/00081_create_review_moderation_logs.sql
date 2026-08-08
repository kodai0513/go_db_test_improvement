-- +goose Up
CREATE TABLE review_moderation_logs (
    id UUID PRIMARY KEY,
    review_id UUID NOT NULL,
    action TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE review_moderation_logs;
