package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberSession は会員のログインセッションを表す。
type MemberSession struct {
	ID           uuid.UUID
	MemberID     uuid.UUID
	SessionToken string
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

// MemberSessionRepository は member_sessions テーブルへのアクセスを抽象化する。
type MemberSessionRepository interface {
	Create(ctx context.Context, s *MemberSession) error
	Get(ctx context.Context, id uuid.UUID) (*MemberSession, error)
	List(ctx context.Context) ([]*MemberSession, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberSession, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
