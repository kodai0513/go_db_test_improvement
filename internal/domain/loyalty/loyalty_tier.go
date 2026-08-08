package loyalty

import (
	"context"

	"github.com/google/uuid"
)

// LoyaltyTier はロイヤリティプログラム内の会員ランクを表す。
type LoyaltyTier struct {
	ID               uuid.UUID
	LoyaltyProgramID uuid.UUID
	Name             string
	MinPoints        int32
}

// LoyaltyTierRepository は loyalty_tiers テーブルへのアクセスを抽象化する。
type LoyaltyTierRepository interface {
	Create(ctx context.Context, t *LoyaltyTier) error
	Get(ctx context.Context, id uuid.UUID) (*LoyaltyTier, error)
	List(ctx context.Context) ([]*LoyaltyTier, error)
	ListByLoyaltyProgramID(ctx context.Context, loyaltyProgramID uuid.UUID) ([]*LoyaltyTier, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
