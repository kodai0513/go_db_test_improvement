package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// Repository implements order.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ order.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, o *order.Order) error {
	row, err := r.q.CreateOrder(ctx, sqlc.CreateOrderParams{
		ID:             o.ID,
		MemberID:       o.MemberID,
		TotalAmountYen: o.TotalAmountYen,
		Status:         o.Status,
		CreatedAt:      o.CreatedAt,
	})
	if err != nil {
		return err
	}
	*o = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*order.Order, error) {
	row, err := r.q.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	o := toDomain(row)
	return &o, nil
}

func (r *Repository) List(ctx context.Context) ([]*order.Order, error) {
	rows, err := r.q.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	orders := make([]*order.Order, 0, len(rows))
	for _, row := range rows {
		o := toDomain(row)
		orders = append(orders, &o)
	}
	return orders, nil
}

func toDomain(row sqlc.Order) order.Order {
	return order.Order{
		ID:             row.ID,
		MemberID:       row.MemberID,
		TotalAmountYen: row.TotalAmountYen,
		Status:         row.Status,
		CreatedAt:      row.CreatedAt,
	}
}
