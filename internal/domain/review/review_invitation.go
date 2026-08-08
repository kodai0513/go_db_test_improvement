package review

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReviewInvitation は注文に対するレビュー投稿依頼を表す。
type ReviewInvitation struct {
	ID       uuid.UUID
	OrderID  uuid.UUID
	MemberID uuid.UUID
	SentAt   time.Time
}

// ReviewInvitationRepository は review_invitations テーブルへのアクセスを抽象化する。
type ReviewInvitationRepository interface {
	Create(ctx context.Context, i *ReviewInvitation) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewInvitation, error)
	List(ctx context.Context) ([]*ReviewInvitation, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*ReviewInvitation, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
