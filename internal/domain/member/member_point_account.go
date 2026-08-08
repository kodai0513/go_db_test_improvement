package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberPointAccount は会員のポイント口座を表す。
type MemberPointAccount struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Balance   int32
	UpdatedAt time.Time
}

// MemberPointAccountRepository は member_point_accounts テーブルへのアクセスを抽象化する。
type MemberPointAccountRepository interface {
	Create(ctx context.Context, a *MemberPointAccount) error
	Get(ctx context.Context, id uuid.UUID) (*MemberPointAccount, error)
	List(ctx context.Context) ([]*MemberPointAccount, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberPointAccount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
