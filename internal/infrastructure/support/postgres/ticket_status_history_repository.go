package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// TicketStatusHistoryRepository は support.TicketStatusHistoryRepository の実装。
type TicketStatusHistoryRepository struct {
	q *sqlc.Queries
}

func NewTicketStatusHistoryRepository(db *sql.DB) *TicketStatusHistoryRepository {
	return &TicketStatusHistoryRepository{q: sqlc.New(db)}
}

var _ support.TicketStatusHistoryRepository = (*TicketStatusHistoryRepository)(nil)

func (r *TicketStatusHistoryRepository) Create(ctx context.Context, h *support.TicketStatusHistory) error {
	row, err := r.q.CreateTicketStatusHistory(ctx, sqlc.CreateTicketStatusHistoryParams{
		ID:        h.ID,
		TicketID:  h.TicketID,
		Status:    h.Status,
		ChangedAt: h.ChangedAt,
	})
	if err != nil {
		return err
	}
	*h = toTicketStatusHistoryDomain(row)
	return nil
}

func (r *TicketStatusHistoryRepository) Get(ctx context.Context, id uuid.UUID) (*support.TicketStatusHistory, error) {
	row, err := r.q.GetTicketStatusHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	h := toTicketStatusHistoryDomain(row)
	return &h, nil
}

func (r *TicketStatusHistoryRepository) List(ctx context.Context) ([]*support.TicketStatusHistory, error) {
	rows, err := r.q.ListTicketStatusHistories(ctx)
	if err != nil {
		return nil, err
	}
	histories := make([]*support.TicketStatusHistory, 0, len(rows))
	for _, row := range rows {
		h := toTicketStatusHistoryDomain(row)
		histories = append(histories, &h)
	}
	return histories, nil
}

func (r *TicketStatusHistoryRepository) ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*support.TicketStatusHistory, error) {
	rows, err := r.q.ListTicketStatusHistoriesByTicketID(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	histories := make([]*support.TicketStatusHistory, 0, len(rows))
	for _, row := range rows {
		h := toTicketStatusHistoryDomain(row)
		histories = append(histories, &h)
	}
	return histories, nil
}

func (r *TicketStatusHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTicketStatusHistory(ctx, id)
}

func toTicketStatusHistoryDomain(row sqlc.TicketStatusHistory) support.TicketStatusHistory {
	return support.TicketStatusHistory{
		ID:        row.ID,
		TicketID:  row.TicketID,
		Status:    row.Status,
		ChangedAt: row.ChangedAt,
	}
}
