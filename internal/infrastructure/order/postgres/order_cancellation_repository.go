package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderCancellationRepository は order.OrderCancellationRepository の実装。
type OrderCancellationRepository struct {
	q *sqlc.Queries
}

func NewOrderCancellationRepository(db *sql.DB) *OrderCancellationRepository {
	return &OrderCancellationRepository{q: sqlc.New(db)}
}

var _ order.OrderCancellationRepository = (*OrderCancellationRepository)(nil)

func (r *OrderCancellationRepository) Create(ctx context.Context, c *order.OrderCancellation) error {
	row, err := r.q.CreateOrderCancellation(ctx, sqlc.CreateOrderCancellationParams{
		ID:          c.ID,
		OrderID:     c.OrderID,
		Reason:      c.Reason,
		CancelledAt: c.CancelledAt,
	})
	if err != nil {
		return err
	}
	*c = toOrderCancellationDomain(row)
	return nil
}

func (r *OrderCancellationRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderCancellation, error) {
	row, err := r.q.GetOrderCancellation(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toOrderCancellationDomain(row)
	return &c, nil
}

func (r *OrderCancellationRepository) List(ctx context.Context) ([]*order.OrderCancellation, error) {
	rows, err := r.q.ListOrderCancellations(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderCancellation, 0, len(rows))
	for _, row := range rows {
		c := toOrderCancellationDomain(row)
		items = append(items, &c)
	}
	return items, nil
}

func (r *OrderCancellationRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderCancellation, error) {
	rows, err := r.q.ListOrderCancellationsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderCancellation, 0, len(rows))
	for _, row := range rows {
		c := toOrderCancellationDomain(row)
		items = append(items, &c)
	}
	return items, nil
}

func (r *OrderCancellationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderCancellation(ctx, id)
}

func toOrderCancellationDomain(row sqlc.OrderCancellation) order.OrderCancellation {
	return order.OrderCancellation{
		ID:          row.ID,
		OrderID:     row.OrderID,
		Reason:      row.Reason,
		CancelledAt: row.CancelledAt,
	}
}
