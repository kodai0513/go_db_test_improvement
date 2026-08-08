package loyalty

import (
	"context"

	"github.com/google/uuid"
)

// PointRule はポイント付与ルールを表す。
type PointRule struct {
	ID               uuid.UUID
	LoyaltyProgramID uuid.UUID
	Action           string
	Points           int32
}

// PointRuleRepository は point_rules テーブルへのアクセスを抽象化する。
type PointRuleRepository interface {
	Create(ctx context.Context, r *PointRule) error
	Get(ctx context.Context, id uuid.UUID) (*PointRule, error)
	List(ctx context.Context) ([]*PointRule, error)
	ListByLoyaltyProgramID(ctx context.Context, loyaltyProgramID uuid.UUID) ([]*PointRule, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
