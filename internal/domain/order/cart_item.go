package order

import (
	"context"

	"github.com/google/uuid"
)

// CartItem はカートに入れられた商品明細を表す。
type CartItem struct {
	ID        uuid.UUID
	CartID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int32
}

// CartItemRepository は cart_items テーブルへのアクセスを抽象化する。
type CartItemRepository interface {
	Create(ctx context.Context, i *CartItem) error
	Get(ctx context.Context, id uuid.UUID) (*CartItem, error)
	List(ctx context.Context) ([]*CartItem, error)
	ListByCartID(ctx context.Context, cartID uuid.UUID) ([]*CartItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
