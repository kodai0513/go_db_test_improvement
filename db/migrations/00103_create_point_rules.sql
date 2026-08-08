-- +goose Up
CREATE TABLE point_rules (
    id UUID PRIMARY KEY,
    loyalty_program_id UUID NOT NULL,
    action TEXT NOT NULL,
    points INTEGER NOT NULL
);

-- +goose Down
DROP TABLE point_rules;
