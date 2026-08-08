package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberPreference は会員の設定情報を表す。
type MemberPreference struct {
	ID              uuid.UUID
	MemberID        uuid.UUID
	Language        string
	NewsletterOptIn bool
	UpdatedAt       time.Time
}

// MemberPreferenceRepository は member_preferences テーブルへのアクセスを抽象化する。
type MemberPreferenceRepository interface {
	Create(ctx context.Context, p *MemberPreference) error
	Get(ctx context.Context, id uuid.UUID) (*MemberPreference, error)
	List(ctx context.Context) ([]*MemberPreference, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberPreference, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
