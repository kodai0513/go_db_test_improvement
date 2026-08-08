-- +goose Up
CREATE TABLE member_tag_assignments (
    id UUID PRIMARY KEY,
    member_id UUID NOT NULL,
    member_tag_id UUID NOT NULL,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE member_tag_assignments;
