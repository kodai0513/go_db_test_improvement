package loyalty

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReferralBonus は紹介による特典ポイント付与を表す。
type ReferralBonus struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Points    int32
	GrantedAt time.Time
}

// ReferralBonusRepository は referral_bonuses テーブルへのアクセスを抽象化する。
type ReferralBonusRepository interface {
	Create(ctx context.Context, b *ReferralBonus) error
	Get(ctx context.Context, id uuid.UUID) (*ReferralBonus, error)
	List(ctx context.Context) ([]*ReferralBonus, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*ReferralBonus, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
