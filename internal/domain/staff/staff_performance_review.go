package staff

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StaffPerformanceReview はスタッフの人事評価を表す。
type StaffPerformanceReview struct {
	ID            uuid.UUID
	StaffMemberID uuid.UUID
	Score         int32
	ReviewedAt    time.Time
}

// StaffPerformanceReviewRepository は staff_performance_reviews テーブルへのアクセスを抽象化する。
type StaffPerformanceReviewRepository interface {
	Create(ctx context.Context, r *StaffPerformanceReview) error
	Get(ctx context.Context, id uuid.UUID) (*StaffPerformanceReview, error)
	List(ctx context.Context) ([]*StaffPerformanceReview, error)
	ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*StaffPerformanceReview, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
