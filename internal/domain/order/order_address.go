package order

import (
	"context"

	"github.com/google/uuid"
)

// OrderAddress は注文の配送先住所を表す。
type OrderAddress struct {
	ID         uuid.UUID
	OrderID    uuid.UUID
	PostalCode string
	Address    string
}

// OrderAddressRepository は order_addresses テーブルへのアクセスを抽象化する。
type OrderAddressRepository interface {
	Create(ctx context.Context, a *OrderAddress) error
	Get(ctx context.Context, id uuid.UUID) (*OrderAddress, error)
	List(ctx context.Context) ([]*OrderAddress, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderAddress, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
