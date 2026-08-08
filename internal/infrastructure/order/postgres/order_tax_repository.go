package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderTaxRepository は order.OrderTaxRepository の実装。
type OrderTaxRepository struct {
	q *sqlc.Queries
}

func NewOrderTaxRepository(db *sql.DB) *OrderTaxRepository {
	return &OrderTaxRepository{q: sqlc.New(db)}
}

var _ order.OrderTaxRepository = (*OrderTaxRepository)(nil)

func (r *OrderTaxRepository) Create(ctx context.Context, t *order.OrderTax) error {
	row, err := r.q.CreateOrderTax(ctx, sqlc.CreateOrderTaxParams{
		ID:           t.ID,
		OrderID:      t.OrderID,
		TaxAmountYen: t.TaxAmountYen,
		TaxRate:      t.TaxRate,
	})
	if err != nil {
		return err
	}
	*t = toOrderTaxDomain(row)
	return nil
}

func (r *OrderTaxRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderTax, error) {
	row, err := r.q.GetOrderTax(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toOrderTaxDomain(row)
	return &t, nil
}

func (r *OrderTaxRepository) List(ctx context.Context) ([]*order.OrderTax, error) {
	rows, err := r.q.ListOrderTaxes(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderTax, 0, len(rows))
	for _, row := range rows {
		t := toOrderTaxDomain(row)
		items = append(items, &t)
	}
	return items, nil
}

func (r *OrderTaxRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderTax, error) {
	rows, err := r.q.ListOrderTaxesByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderTax, 0, len(rows))
	for _, row := range rows {
		t := toOrderTaxDomain(row)
		items = append(items, &t)
	}
	return items, nil
}

func (r *OrderTaxRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderTax(ctx, id)
}

func toOrderTaxDomain(row sqlc.OrderTax) order.OrderTax {
	return order.OrderTax{
		ID:           row.ID,
		OrderID:      row.OrderID,
		TaxAmountYen: row.TaxAmountYen,
		TaxRate:      row.TaxRate,
	}
}
