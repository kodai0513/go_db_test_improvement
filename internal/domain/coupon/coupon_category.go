package coupon

import (
	"context"

	"github.com/google/uuid"
)

// CouponCategory はクーポンの適用対象カテゴリを表す。
type CouponCategory struct {
	ID         uuid.UUID
	CouponID   uuid.UUID
	CategoryID uuid.UUID
}

// CouponCategoryRepository は coupon_categories テーブルへのアクセスを抽象化する。
type CouponCategoryRepository interface {
	Create(ctx context.Context, c *CouponCategory) error
	Get(ctx context.Context, id uuid.UUID) (*CouponCategory, error)
	List(ctx context.Context) ([]*CouponCategory, error)
	ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*CouponCategory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
