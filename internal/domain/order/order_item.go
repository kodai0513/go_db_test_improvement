package order

import (
	"context"

	"github.com/google/uuid"
)

// OrderItem は注文明細(注文に含まれる商品)を表す。
type OrderItem struct {
	ID           uuid.UUID
	OrderID      uuid.UUID
	ProductID    uuid.UUID
	Quantity     int32
	UnitPriceYen int32
}

// OrderItemRepository は order_items テーブルへのアクセスを抽象化する。
type OrderItemRepository interface {
	Create(ctx context.Context, i *OrderItem) error
	Get(ctx context.Context, id uuid.UUID) (*OrderItem, error)
	List(ctx context.Context) ([]*OrderItem, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
