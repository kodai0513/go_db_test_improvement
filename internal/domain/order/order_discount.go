package order

import (
	"context"

	"github.com/google/uuid"
)

// OrderDiscount は注文に適用された割引を表す。CouponID はクーポンが
// 紐付かない場合はゼロ値(uuid.Nil)を許容する。
type OrderDiscount struct {
	ID                uuid.UUID
	OrderID           uuid.UUID
	CouponID          uuid.UUID
	DiscountAmountYen int32
}

// OrderDiscountRepository は order_discounts テーブルへのアクセスを抽象化する。
type OrderDiscountRepository interface {
	Create(ctx context.Context, d *OrderDiscount) error
	Get(ctx context.Context, id uuid.UUID) (*OrderDiscount, error)
	List(ctx context.Context) ([]*OrderDiscount, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*OrderDiscount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
