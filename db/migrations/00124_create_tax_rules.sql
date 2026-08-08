-- +goose Up
CREATE TABLE tax_rules (
    id UUID PRIMARY KEY,
    tax_rate_id UUID NOT NULL,
    region TEXT NOT NULL
);

-- +goose Down
DROP TABLE tax_rules;
