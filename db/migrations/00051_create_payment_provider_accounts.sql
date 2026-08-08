-- +goose Up
CREATE TABLE payment_provider_accounts (
    id UUID PRIMARY KEY,
    payment_provider_id UUID NOT NULL,
    member_id UUID NOT NULL,
    external_account_id TEXT NOT NULL
);

-- +goose Down
DROP TABLE payment_provider_accounts;
