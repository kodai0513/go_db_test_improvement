package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberTagAssignment は会員とタグの割り当てを表す。
type MemberTagAssignment struct {
	ID          uuid.UUID
	MemberID    uuid.UUID
	MemberTagID uuid.UUID
	AssignedAt  time.Time
}

// MemberTagAssignmentRepository は member_tag_assignments テーブルへのアクセスを抽象化する。
type MemberTagAssignmentRepository interface {
	Create(ctx context.Context, a *MemberTagAssignment) error
	Get(ctx context.Context, id uuid.UUID) (*MemberTagAssignment, error)
	List(ctx context.Context) ([]*MemberTagAssignment, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberTagAssignment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
