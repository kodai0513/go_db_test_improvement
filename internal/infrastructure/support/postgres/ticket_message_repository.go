package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// TicketMessageRepository は support.TicketMessageRepository の実装。
type TicketMessageRepository struct {
	q *sqlc.Queries
}

func NewTicketMessageRepository(db *sql.DB) *TicketMessageRepository {
	return &TicketMessageRepository{q: sqlc.New(db)}
}

var _ support.TicketMessageRepository = (*TicketMessageRepository)(nil)

func (r *TicketMessageRepository) Create(ctx context.Context, m *support.TicketMessage) error {
	row, err := r.q.CreateTicketMessage(ctx, sqlc.CreateTicketMessageParams{
		ID:             m.ID,
		TicketID:       m.TicketID,
		SenderMemberID: m.SenderMemberID,
		Body:           m.Body,
		CreatedAt:      m.CreatedAt,
	})
	if err != nil {
		return err
	}
	*m = toTicketMessageDomain(row)
	return nil
}

func (r *TicketMessageRepository) Get(ctx context.Context, id uuid.UUID) (*support.TicketMessage, error) {
	row, err := r.q.GetTicketMessage(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toTicketMessageDomain(row)
	return &m, nil
}

func (r *TicketMessageRepository) List(ctx context.Context) ([]*support.TicketMessage, error) {
	rows, err := r.q.ListTicketMessages(ctx)
	if err != nil {
		return nil, err
	}
	messages := make([]*support.TicketMessage, 0, len(rows))
	for _, row := range rows {
		m := toTicketMessageDomain(row)
		messages = append(messages, &m)
	}
	return messages, nil
}

func (r *TicketMessageRepository) ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*support.TicketMessage, error) {
	rows, err := r.q.ListTicketMessagesByTicketID(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	messages := make([]*support.TicketMessage, 0, len(rows))
	for _, row := range rows {
		m := toTicketMessageDomain(row)
		messages = append(messages, &m)
	}
	return messages, nil
}

func (r *TicketMessageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTicketMessage(ctx, id)
}

func toTicketMessageDomain(row sqlc.TicketMessage) support.TicketMessage {
	return support.TicketMessage{
		ID:             row.ID,
		TicketID:       row.TicketID,
		SenderMemberID: row.SenderMemberID,
		Body:           row.Body,
		CreatedAt:      row.CreatedAt,
	}
}
