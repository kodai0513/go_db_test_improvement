-- +goose Up
CREATE TABLE review_reports (
    id UUID PRIMARY KEY,
    review_id UUID NOT NULL,
    reporter_member_id UUID NOT NULL,
    reason TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE review_reports;
