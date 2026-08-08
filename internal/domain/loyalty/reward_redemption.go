package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// RewardRedemption は会員による報酬アイテムの交換を表す。
type RewardRedemption struct {
	ID           uuid.UUID
	MemberID     uuid.UUID
	RewardItemID uuid.UUID
	RedeemedAt   time.Time
}

// RewardRedemptionRepository は reward_redemptions テーブルへのアクセスを抽象化する。
type RewardRedemptionRepository interface {
	Create(ctx context.Context, r *RewardRedemption) error
	Get(ctx context.Context, id uuid.UUID) (*RewardRedemption, error)
	List(ctx context.Context) ([]*RewardRedemption, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*RewardRedemption, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
