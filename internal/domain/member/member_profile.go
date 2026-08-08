package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberProfile は会員プロフィールを表す。
type MemberProfile struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	BirthDate time.Time
	Gender    string
	Bio       string
	AvatarURL string
	UpdatedAt time.Time
}

// MemberProfileRepository は member_profiles テーブルへのアクセスを抽象化する。
type MemberProfileRepository interface {
	Create(ctx context.Context, p *MemberProfile) error
	Get(ctx context.Context, id uuid.UUID) (*MemberProfile, error)
	List(ctx context.Context) ([]*MemberProfile, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberProfile, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
