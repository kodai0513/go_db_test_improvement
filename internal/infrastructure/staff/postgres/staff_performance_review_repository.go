package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// StaffPerformanceReviewRepository は staff.StaffPerformanceReviewRepository の実装。
type StaffPerformanceReviewRepository struct {
	q *sqlc.Queries
}

func NewStaffPerformanceReviewRepository(db *sql.DB) *StaffPerformanceReviewRepository {
	return &StaffPerformanceReviewRepository{q: sqlc.New(db)}
}

var _ staff.StaffPerformanceReviewRepository = (*StaffPerformanceReviewRepository)(nil)

func (r *StaffPerformanceReviewRepository) Create(ctx context.Context, rv *staff.StaffPerformanceReview) error {
	row, err := r.q.CreateStaffPerformanceReview(ctx, sqlc.CreateStaffPerformanceReviewParams{
		ID:            rv.ID,
		StaffMemberID: rv.StaffMemberID,
		Score:         rv.Score,
		ReviewedAt:    rv.ReviewedAt,
	})
	if err != nil {
		return err
	}
	*rv = toStaffPerformanceReviewDomain(row)
	return nil
}

func (r *StaffPerformanceReviewRepository) Get(ctx context.Context, id uuid.UUID) (*staff.StaffPerformanceReview, error) {
	row, err := r.q.GetStaffPerformanceReview(ctx, id)
	if err != nil {
		return nil, err
	}
	rv := toStaffPerformanceReviewDomain(row)
	return &rv, nil
}

func (r *StaffPerformanceReviewRepository) List(ctx context.Context) ([]*staff.StaffPerformanceReview, error) {
	rows, err := r.q.ListStaffPerformanceReviews(ctx)
	if err != nil {
		return nil, err
	}
	reviews := make([]*staff.StaffPerformanceReview, 0, len(rows))
	for _, row := range rows {
		rv := toStaffPerformanceReviewDomain(row)
		reviews = append(reviews, &rv)
	}
	return reviews, nil
}

func (r *StaffPerformanceReviewRepository) ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*staff.StaffPerformanceReview, error) {
	rows, err := r.q.ListStaffPerformanceReviewsByStaffMemberID(ctx, staffMemberID)
	if err != nil {
		return nil, err
	}
	reviews := make([]*staff.StaffPerformanceReview, 0, len(rows))
	for _, row := range rows {
		rv := toStaffPerformanceReviewDomain(row)
		reviews = append(reviews, &rv)
	}
	return reviews, nil
}

func (r *StaffPerformanceReviewRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStaffPerformanceReview(ctx, id)
}

func toStaffPerformanceReviewDomain(row sqlc.StaffPerformanceReview) staff.StaffPerformanceReview {
	return staff.StaffPerformanceReview{
		ID:            row.ID,
		StaffMemberID: row.StaffMemberID,
		Score:         row.Score,
		ReviewedAt:    row.ReviewedAt,
	}
}
