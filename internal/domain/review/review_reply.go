package review

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReviewReply はレビューへの返信コメントを表す。
type ReviewReply struct {
	ID        uuid.UUID
	ReviewID  uuid.UUID
	MemberID  uuid.UUID
	Comment   string
	CreatedAt time.Time
}

// ReviewReplyRepository は review_replies テーブルへのアクセスを抽象化する。
type ReviewReplyRepository interface {
	Create(ctx context.Context, r *ReviewReply) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewReply, error)
	List(ctx context.Context) ([]*ReviewReply, error)
	ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*ReviewReply, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
