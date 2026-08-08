package order

import (
	"context"

	"github.com/google/uuid"
)

// OrderTax は注文に課される税額情報を表す。
type OrderTax struct {
	ID           uuid.UUID
	OrderID      uuid.UUID
	TaxAmountYen int32
	TaxRate      int32
}

// OrderTaxRepository は order_taxes テーブルへのアクセスを抽象化する。
type OrderTaxRepository interface {
	Create(ctx context.Context, t *OrderTax) error
	Get(ctx context.Context, id uuid.UUID) (*OrderTax, error)
	List(ctx context.Context) ([]*OrderTax, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderTax, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
