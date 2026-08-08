-- +goose Up
CREATE TABLE promo_codes (
    id UUID PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    campaign_id UUID NOT NULL
);

-- +goose Down
DROP TABLE promo_codes;
