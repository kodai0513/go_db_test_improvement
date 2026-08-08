package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// TicketAttachmentRepository は support.TicketAttachmentRepository の実装。
type TicketAttachmentRepository struct {
	q *sqlc.Queries
}

func NewTicketAttachmentRepository(db *sql.DB) *TicketAttachmentRepository {
	return &TicketAttachmentRepository{q: sqlc.New(db)}
}

var _ support.TicketAttachmentRepository = (*TicketAttachmentRepository)(nil)

func (r *TicketAttachmentRepository) Create(ctx context.Context, a *support.TicketAttachment) error {
	row, err := r.q.CreateTicketAttachment(ctx, sqlc.CreateTicketAttachmentParams{
		ID:              a.ID,
		TicketMessageID: a.TicketMessageID,
		URL:             a.URL,
	})
	if err != nil {
		return err
	}
	*a = toTicketAttachmentDomain(row)
	return nil
}

func (r *TicketAttachmentRepository) Get(ctx context.Context, id uuid.UUID) (*support.TicketAttachment, error) {
	row, err := r.q.GetTicketAttachment(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toTicketAttachmentDomain(row)
	return &a, nil
}

func (r *TicketAttachmentRepository) List(ctx context.Context) ([]*support.TicketAttachment, error) {
	rows, err := r.q.ListTicketAttachments(ctx)
	if err != nil {
		return nil, err
	}
	attachments := make([]*support.TicketAttachment, 0, len(rows))
	for _, row := range rows {
		a := toTicketAttachmentDomain(row)
		attachments = append(attachments, &a)
	}
	return attachments, nil
}

func (r *TicketAttachmentRepository) ListByTicketMessageID(ctx context.Context, ticketMessageID uuid.UUID) ([]*support.TicketAttachment, error) {
	rows, err := r.q.ListTicketAttachmentsByTicketMessageID(ctx, ticketMessageID)
	if err != nil {
		return nil, err
	}
	attachments := make([]*support.TicketAttachment, 0, len(rows))
	for _, row := range rows {
		a := toTicketAttachmentDomain(row)
		attachments = append(attachments, &a)
	}
	return attachments, nil
}

func (r *TicketAttachmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTicketAttachment(ctx, id)
}

func toTicketAttachmentDomain(row sqlc.TicketAttachment) support.TicketAttachment {
	return support.TicketAttachment{
		ID:              row.ID,
		TicketMessageID: row.TicketMessageID,
		URL:             row.URL,
	}
}
