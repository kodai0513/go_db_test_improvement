package review

import (
	"context"

	"github.com/google/uuid"
)

// ReviewVote はレビューに対する「参考になった」投票を表す。
type ReviewVote struct {
	ID        uuid.UUID
	ReviewID  uuid.UUID
	MemberID  uuid.UUID
	IsHelpful bool
}

// ReviewVoteRepository は review_votes テーブルへのアクセスを抽象化する。
type ReviewVoteRepository interface {
	Create(ctx context.Context, v *ReviewVote) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewVote, error)
	List(ctx context.Context) ([]*ReviewVote, error)
	ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*ReviewVote, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
