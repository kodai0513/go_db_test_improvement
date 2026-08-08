package review

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReviewReport はレビューに対する通報を表す。
type ReviewReport struct {
	ID               uuid.UUID
	ReviewID         uuid.UUID
	ReporterMemberID uuid.UUID
	Reason           string
	CreatedAt        time.Time
}

// ReviewReportRepository は review_reports テーブルへのアクセスを抽象化する。
type ReviewReportRepository interface {
	Create(ctx context.Context, r *ReviewReport) error
	Get(ctx context.Context, id uuid.UUID) (*ReviewReport, error)
	List(ctx context.Context) ([]*ReviewReport, error)
	ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*ReviewReport, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
