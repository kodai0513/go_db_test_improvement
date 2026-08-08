package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// LoyaltyTierHistory は会員のロイヤリティランク変更履歴を表す。
type LoyaltyTierHistory struct {
	ID            uuid.UUID
	MemberID      uuid.UUID
	LoyaltyTierID uuid.UUID
	ChangedAt     time.Time
}

// LoyaltyTierHistoryRepository は loyalty_tier_histories テーブルへのアクセスを抽象化する。
type LoyaltyTierHistoryRepository interface {
	Create(ctx context.Context, h *LoyaltyTierHistory) error
	Get(ctx context.Context, id uuid.UUID) (*LoyaltyTierHistory, error)
	List(ctx context.Context) ([]*LoyaltyTierHistory, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*LoyaltyTierHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
