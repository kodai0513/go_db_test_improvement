package pricing

import (
	"context"

	"github.com/google/uuid"
)

// DiscountTier は数量に応じた割引段階を表す。
type DiscountTier struct {
	ID              uuid.UUID
	MinQuantity     int32
	DiscountPercent int32
}

// DiscountTierRepository は discount_tiers テーブルへのアクセスを抽象化する。
type DiscountTierRepository interface {
	Create(ctx context.Context, d *DiscountTier) error
	Get(ctx context.Context, id uuid.UUID) (*DiscountTier, error)
	List(ctx context.Context) ([]*DiscountTier, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
