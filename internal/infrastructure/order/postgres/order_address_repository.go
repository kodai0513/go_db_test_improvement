package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres/sqlc"
)

// OrderAddressRepository は order.OrderAddressRepository の実装。
type OrderAddressRepository struct {
	q *sqlc.Queries
}

func NewOrderAddressRepository(db *sql.DB) *OrderAddressRepository {
	return &OrderAddressRepository{q: sqlc.New(db)}
}

var _ order.OrderAddressRepository = (*OrderAddressRepository)(nil)

func (r *OrderAddressRepository) Create(ctx context.Context, a *order.OrderAddress) error {
	row, err := r.q.CreateOrderAddress(ctx, sqlc.CreateOrderAddressParams{
		ID:         a.ID,
		OrderID:    a.OrderID,
		PostalCode: a.PostalCode,
		Address:    a.Address,
	})
	if err != nil {
		return err
	}
	*a = toOrderAddressDomain(row)
	return nil
}

func (r *OrderAddressRepository) Get(ctx context.Context, id uuid.UUID) (*order.OrderAddress, error) {
	row, err := r.q.GetOrderAddress(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toOrderAddressDomain(row)
	return &a, nil
}

func (r *OrderAddressRepository) List(ctx context.Context) ([]*order.OrderAddress, error) {
	rows, err := r.q.ListOrderAddresses(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderAddress, 0, len(rows))
	for _, row := range rows {
		a := toOrderAddressDomain(row)
		items = append(items, &a)
	}
	return items, nil
}

func (r *OrderAddressRepository) ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*order.OrderAddress, error) {
	rows, err := r.q.ListOrderAddressesByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	items := make([]*order.OrderAddress, 0, len(rows))
	for _, row := range rows {
		a := toOrderAddressDomain(row)
		items = append(items, &a)
	}
	return items, nil
}

func (r *OrderAddressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteOrderAddress(ctx, id)
}

func toOrderAddressDomain(row sqlc.OrderAddress) order.OrderAddress {
	return order.OrderAddress{
		ID:         row.ID,
		OrderID:    row.OrderID,
		PostalCode: row.PostalCode,
		Address:    row.Address,
	}
}
