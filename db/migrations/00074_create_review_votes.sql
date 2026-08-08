-- +goose Up
CREATE TABLE review_votes (
    id UUID PRIMARY KEY,
    review_id UUID NOT NULL,
    member_id UUID NOT NULL,
    is_helpful BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE review_votes;
