-- +goose Up
CREATE TABLE ticket_attachments (
    id UUID PRIMARY KEY,
    ticket_message_id UUID NOT NULL,
    url TEXT NOT NULL
);

-- +goose Down
DROP TABLE ticket_attachments;
