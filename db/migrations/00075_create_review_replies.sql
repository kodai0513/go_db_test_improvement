-- +goose Up
CREATE TABLE review_replies (
    id UUID PRIMARY KEY,
    review_id UUID NOT NULL,
    member_id UUID NOT NULL,
    comment TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE review_replies;
