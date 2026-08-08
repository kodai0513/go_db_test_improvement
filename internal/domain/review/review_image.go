package review

import (
	"context"

	"github.com/google/uuid"
)

// ReviewImage はレビューに添付された画像を表す。
type ReviewImage struct {
	ID       uuid.UUID
	ReviewID uuid.UUID
	URL      string
}

// ReviewImageRepository は review_images テーブルへのアクセスを抽象化する。
type ReviewImageRepository interface {
	Create(ctx context.Context, i *ReviewImage) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewImage, error)
	List(ctx context.Context) ([]*ReviewImage, error)
	ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*ReviewImage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
