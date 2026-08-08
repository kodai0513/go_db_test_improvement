package coupon

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// CouponUsage はクーポンの利用実績を表す。
type CouponUsage struct {
	ID       uuid.UUID
	CouponID uuid.UUID
	OrderID  uuid.UUID
	UsedAt   time.Time
}

// CouponUsageRepository は coupon_usages テーブルへのアクセスを抽象化する。
type CouponUsageRepository interface {
	Create(ctx context.Context, u *CouponUsage) error
	Get(ctx context.Context, id uuid.UUID) (*CouponUsage, error)
	List(ctx context.Context) ([]*CouponUsage, error)
	ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*CouponUsage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
