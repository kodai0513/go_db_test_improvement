-- +goose Up
CREATE TABLE vendor_bank_accounts (
    id UUID PRIMARY KEY,
    vendor_id UUID NOT NULL,
    account_number TEXT NOT NULL
);

-- +goose Down
DROP TABLE vendor_bank_accounts;
