package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderDiscountRepository は order.OrderDiscountRepository の実装。
type OrderDiscountRepository struct {
	q *sqlc.Queries
}

func NewOrderDiscountRepository(db *sql.DB) *OrderDiscountRepository {
	return &OrderDiscountRepository{q: sqlc.New(db)}
}

var _ order.OrderDiscountRepository = (*OrderDiscountRepository)(nil)

func (r *OrderDiscountRepository) Create(ctx context.Context, d *order.OrderDiscount) error {
	row, err := r.q.CreateOrderDiscount(ctx, sqlc.CreateOrderDiscountParams{
		ID:                d.ID,
		OrderID:           d.OrderID,
		CouponID:          d.CouponID,
		DiscountAmountYen: d.DiscountAmountYen,
	})
	if err != nil {
		return err
	}
	*d = toOrderDiscountDomain(row)
	return nil
}

func (r *OrderDiscountRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderDiscount, error) {
	row, err := r.q.GetOrderDiscount(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toOrderDiscountDomain(row)
	return &d, nil
}

func (r *OrderDiscountRepository) List(ctx context.Context) ([]*order.OrderDiscount, error) {
	rows, err := r.q.ListOrderDiscounts(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderDiscount, 0, len(rows))
	for _, row := range rows {
		d := toOrderDiscountDomain(row)
		items = append(items, &d)
	}
	return items, nil
}

func (r *OrderDiscountRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderDiscount, error) {
	rows, err := r.q.ListOrderDiscountsByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderDiscount, 0, len(rows))
	for _, row := range rows {
		d := toOrderDiscountDomain(row)
		items = append(items, &d)
	}
	return items, nil
}

func (r *OrderDiscountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderDiscount(ctx, id)
}

func toOrderDiscountDomain(row sqlc.OrderDiscount) order.OrderDiscount {
	return order.OrderDiscount{
		ID:                row.ID,
		OrderID:           row.OrderID,
		CouponID:          row.CouponID,
		DiscountAmountYen: row.DiscountAmountYen,
	}
}
