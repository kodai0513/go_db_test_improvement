package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberCredential は会員の認証情報を表す。
type MemberCredential struct {
	ID            uuid.UUID
	MemberID      uuid.UUID
	PasswordHash  string
	LastChangedAt time.Time
}

// MemberCredentialRepository は member_credentials テーブルへのアクセスを抽象化する。
type MemberCredentialRepository interface {
	Create(ctx context.Context, c *MemberCredential) error
	Get(ctx context.Context, id uuid.UUID) (*MemberCredential, error)
	List(ctx context.Context) ([]*MemberCredential, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberCredential, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
