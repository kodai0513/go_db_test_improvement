-- +goose Up
CREATE TABLE campaign_targets (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL,
    member_id UUID NOT NULL
);

-- +goose Down
DROP TABLE campaign_targets;
