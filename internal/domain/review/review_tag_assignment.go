package review

import (
	"context"

	"github.com/google/uuid"
)

// ReviewTagAssignment はレビューとレビュータグの紐付けを表す。
type ReviewTagAssignment struct {
	ID          uuid.UUID
	ReviewID    uuid.UUID
	ReviewTagID uuid.UUID
}

// ReviewTagAssignmentRepository は review_tag_assignments テーブルへのアクセスを抽象化する。
type ReviewTagAssignmentRepository interface {
	Create(ctx context.Context, a *ReviewTagAssignment) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewTagAssignment, error)
	List(ctx context.Context) ([]*ReviewTagAssignment, error)
	ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*ReviewTagAssignment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
