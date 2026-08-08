package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberReferral は会員による紹介実績を表す。
type MemberReferral struct {
	ID               uuid.UUID
	ReferrerMemberID uuid.UUID
	ReferredMemberID uuid.UUID
	RewardYen        int32
	CreatedAt        time.Time
}

// MemberReferralRepository は member_referrals テーブルへのアクセスを抽象化する。
type MemberReferralRepository interface {
	Create(ctx context.Context, r *MemberReferral) error
	Get(ctx context.Context, id uuid.UUID) (*MemberReferral, error)
	List(ctx context.Context) ([]*MemberReferral, error)
	ListByReferrerMemberID(ctx context.Context, referrerMemberID uuid.UUID) ([]*MemberReferral, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
