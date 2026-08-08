package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderNoteRepository は order.OrderNoteRepository の実装。
type OrderNoteRepository struct {
	q *sqlc.Queries
}

func NewOrderNoteRepository(db *sql.DB) *OrderNoteRepository {
	return &OrderNoteRepository{q: sqlc.New(db)}
}

var _ order.OrderNoteRepository = (*OrderNoteRepository)(nil)

func (r *OrderNoteRepository) Create(ctx context.Context, n *order.OrderNote) error {
	row, err := r.q.CreateOrderNote(ctx, sqlc.CreateOrderNoteParams{
		ID:        n.ID,
		OrderID:   n.OrderID,
		Note:      n.Note,
		CreatedAt: n.CreatedAt,
	})
	if err != nil {
		return err
	}
	*n = toOrderNoteDomain(row)
	return nil
}

func (r *OrderNoteRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderNote, error) {
	row, err := r.q.GetOrderNote(ctx, id)
	if err != nil {
		return nil, err
	}
	n := toOrderNoteDomain(row)
	return &n, nil
}

func (r *OrderNoteRepository) List(ctx context.Context) ([]*order.OrderNote, error) {
	rows, err := r.q.ListOrderNotes(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderNote, 0, len(rows))
	for _, row := range rows {
		n := toOrderNoteDomain(row)
		items = append(items, &n)
	}
	return items, nil
}

func (r *OrderNoteRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderNote, error) {
	rows, err := r.q.ListOrderNotesByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderNote, 0, len(rows))
	for _, row := range rows {
		n := toOrderNoteDomain(row)
		items = append(items, &n)
	}
	return items, nil
}

func (r *OrderNoteRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderNote(ctx, id)
}

func toOrderNoteDomain(row sqlc.OrderNote) order.OrderNote {
	return order.OrderNote{
		ID:        row.ID,
		OrderID:   row.OrderID,
		Note:      row.Note,
		CreatedAt: row.CreatedAt,
	}
}
