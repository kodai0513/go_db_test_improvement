-- +goose Up
CREATE TABLE staff_performance_reviews (
    id UUID PRIMARY KEY,
    staff_member_id UUID NOT NULL,
    score INTEGER NOT NULL,
    reviewed_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE staff_performance_reviews;
