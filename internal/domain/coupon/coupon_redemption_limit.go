package coupon

import (
	"context"

	"github.com/google/uuid"
)

// CouponRedemptionLimit はクーポンの利用可能回数の上限を表す。
type CouponRedemptionLimit struct {
	ID             uuid.UUID
	CouponID       uuid.UUID
	MaxRedemptions int32
}

// CouponRedemptionLimitRepository は coupon_redemption_limits テーブルへのアクセスを抽象化する。
type CouponRedemptionLimitRepository interface {
	Create(ctx context.Context, l *CouponRedemptionLimit) error
	Get(ctx context.Context, id uuid.UUID) (*CouponRedemptionLimit, error)
	List(ctx context.Context) ([]*CouponRedemptionLimit, error)
	ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*CouponRedemptionLimit, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
