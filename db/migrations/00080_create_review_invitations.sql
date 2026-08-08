-- +goose Up
CREATE TABLE review_invitations (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    member_id UUID NOT NULL,
    sent_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE review_invitations;
