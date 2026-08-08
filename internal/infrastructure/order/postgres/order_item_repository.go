package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderItemRepository は order.OrderItemRepository の実装。
type OrderItemRepository struct {
	q *sqlc.Queries
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{q: sqlc.New(db)}
}

var _ order.OrderItemRepository = (*OrderItemRepository)(nil)

func (r *OrderItemRepository) Create(ctx context.Context, i *order.OrderItem) error {
	row, err := r.q.CreateOrderItem(ctx, sqlc.CreateOrderItemParams{
		ID:           i.ID,
		OrderID:      i.OrderID,
		ProductID:    i.ProductID,
		Quantity:     i.Quantity,
		UnitPriceYen: i.UnitPriceYen,
	})
	if err != nil {
		return err
	}
	*i = toOrderItemDomain(row)
	return nil
}

func (r *OrderItemRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderItem, error) {
	row, err := r.q.GetOrderItem(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toOrderItemDomain(row)
	return &i, nil
}

func (r *OrderItemRepository) List(ctx context.Context) ([]*order.OrderItem, error) {
	rows, err := r.q.ListOrderItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderItem, 0, len(rows))
	for _, row := range rows {
		i := toOrderItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *OrderItemRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderItem, error) {
	rows, err := r.q.ListOrderItemsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderItem, 0, len(rows))
	for _, row := range rows {
		i := toOrderItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *OrderItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderItem(ctx, id)
}

func toOrderItemDomain(row sqlc.OrderItem) order.OrderItem {
	return order.OrderItem{
		ID:           row.ID,
		OrderID:      row.OrderID,
		ProductID:    row.ProductID,
		Quantity:     row.Quantity,
		UnitPriceYen: row.UnitPriceYen,
	}
}
