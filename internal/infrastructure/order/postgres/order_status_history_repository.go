package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderStatusHistoryRepository は order.OrderStatusHistoryRepository の実装。
type OrderStatusHistoryRepository struct {
	q *sqlc.Queries
}

func NewOrderStatusHistoryRepository(db *sql.DB) *OrderStatusHistoryRepository {
	return &OrderStatusHistoryRepository{q: sqlc.New(db)}
}

var _ order.OrderStatusHistoryRepository = (*OrderStatusHistoryRepository)(nil)

func (r *OrderStatusHistoryRepository) Create(ctx context.Context, h *order.OrderStatusHistory) error {
	row, err := r.q.CreateOrderStatusHistory(ctx, sqlc.CreateOrderStatusHistoryParams{
		ID:        h.ID,
		OrderID:   h.OrderID,
		Status:    h.Status,
		ChangedAt: h.ChangedAt,
	})
	if err != nil {
		return err
	}
	*h = toOrderStatusHistoryDomain(row)
	return nil
}

func (r *OrderStatusHistoryRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderStatusHistory, error) {
	row, err := r.q.GetOrderStatusHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	h := toOrderStatusHistoryDomain(row)
	return &h, nil
}

func (r *OrderStatusHistoryRepository) List(ctx context.Context) ([]*order.OrderStatusHistory, error) {
	rows, err := r.q.ListOrderStatusHistories(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderStatusHistory, 0, len(rows))
	for _, row := range rows {
		h := toOrderStatusHistoryDomain(row)
		items = append(items, &h)
	}
	return items, nil
}

func (r *OrderStatusHistoryRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderStatusHistory, error) {
	rows, err := r.q.ListOrderStatusHistoriesByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderStatusHistory, 0, len(rows))
	for _, row := range rows {
		h := toOrderStatusHistoryDomain(row)
		items = append(items, &h)
	}
	return items, nil
}

func (r *OrderStatusHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderStatusHistory(ctx, id)
}

func toOrderStatusHistoryDomain(row sqlc.OrderStatusHistory) order.OrderStatusHistory {
	return order.OrderStatusHistory{
		ID:        row.ID,
		OrderID:   row.OrderID,
		Status:    row.Status,
		ChangedAt: row.ChangedAt,
	}
}
