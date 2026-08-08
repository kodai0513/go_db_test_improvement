package coupon

import (
	"context"

	"github.com/google/uuid"
)

// DiscountRule はクーポン適用に必要な最低利用金額のルールを表す。
type DiscountRule struct {
	ID           uuid.UUID
	CouponID     uuid.UUID
	MinAmountYen int32
}

// DiscountRuleRepository は discount_rules テーブルへのアクセスを抽象化する。
type DiscountRuleRepository interface {
	Create(ctx context.Context, r *DiscountRule) error
	Get(ctx context.Context, id uuid.UUID) (*DiscountRule, error)
	List(ctx context.Context) ([]*DiscountRule, error)
	ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*DiscountRule, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
