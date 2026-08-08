package coupon

import (
	"context"

	"github.com/google/uuid"
)

// CouponProduct はクーポンの適用対象商品を表す。
type CouponProduct struct {
	ID        uuid.UUID
	CouponID  uuid.UUID
	ProductID uuid.UUID
}

// CouponProductRepository は coupon_products テーブルへのアクセスを抽象化する。
type CouponProductRepository interface {
	Create(ctx context.Context, p *CouponProduct) error
	Get(ctx context.Context, id uuid.UUID) (*CouponProduct, error)
	List(ctx context.Context) ([]*CouponProduct, error)
	ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*CouponProduct, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
