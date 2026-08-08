package coupon

import (
	"context"

	"github.com/google/uuid"
)

// CouponStackingRule はクーポン同士の併用可否を表す。
type CouponStackingRule struct {
	ID                    uuid.UUID
	CouponID              uuid.UUID
	StackableWithCouponID uuid.UUID
}

// CouponStackingRuleRepository は coupon_stacking_rules テーブルへのアクセスを抽象化する。
type CouponStackingRuleRepository interface {
	Create(ctx context.Context, r *CouponStackingRule) error
	Get(ctx context.Context, id uuid.UUID) (*CouponStackingRule, error)
	List(ctx context.Context) ([]*CouponStackingRule, error)
	ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*CouponStackingRule, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
