package review

import (
	"context"

	"github.com/google/uuid"
)

// ReviewSummary は商品ごとのレビュー集計(平均評価・件数)を表す。
type ReviewSummary struct {
	ID            uuid.UUID
	ProductID     uuid.UUID
	AverageRating int32
	ReviewCount   int32
}

// ReviewSummaryRepository は review_summaries テーブルへのアクセスを抽象化する。
type ReviewSummaryRepository interface {
	Create(ctx context.Context, s *ReviewSummary) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewSummary, error)
	List(ctx context.Context) ([]*ReviewSummary, error)
	ListByProductID(ctx context.Context, productID uuid.UUID) ([]*ReviewSummary, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
