package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewReportRepository は review.ReviewReportRepository の実装。
type ReviewReportRepository struct {
	q *sqlc.Queries
}

func NewReviewReportRepository(db *sql.DB) *ReviewReportRepository {
	return &ReviewReportRepository{q: sqlc.New(db)}
}

var _ review.ReviewReportRepository = (*ReviewReportRepository)(nil)

func (r *ReviewReportRepository) Create(ctx context.Context, rp *review.ReviewReport) error {
	row, err := r.q.CreateReviewReport(ctx, sqlc.CreateReviewReportParams{
		ID:               rp.ID,
		ReviewID:         rp.ReviewID,
		ReporterMemberID: rp.ReporterMemberID,
		Reason:           rp.Reason,
		CreatedAt:        rp.CreatedAt,
	})
	if err != nil {
		return err
	}
	*rp = toReviewReportDomain(row)
	return nil
}

func (r *ReviewReportRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewReport, error) {
	row, err := r.q.GetReviewReport(ctx, id)
	if err != nil {
		return nil, err
	}
	rp := toReviewReportDomain(row)
	return &rp, nil
}

func (r *ReviewReportRepository) List(ctx context.Context) ([]*review.ReviewReport, error) {
	rows, err := r.q.ListReviewReports(ctx)
	if err != nil {
		return nil, err
	}
	reports := make([]*review.ReviewReport, 0, len(rows))
	for _, row := range rows {
		rp := toReviewReportDomain(row)
		reports = append(reports, &rp)
	}
	return reports, nil
}

func (r *ReviewReportRepository) ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*review.ReviewReport, error) {
	rows, err := r.q.ListReviewReportsByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	reports := make([]*review.ReviewReport, 0, len(rows))
	for _, row := range rows {
		rp := toReviewReportDomain(row)
		reports = append(reports, &rp)
	}
	return reports, nil
}

func (r *ReviewReportRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewReport(ctx, id)
}

func toReviewReportDomain(row sqlc.ReviewReport) review.ReviewReport {
	return review.ReviewReport{
		ID:               row.ID,
		ReviewID:         row.ReviewID,
		ReporterMemberID: row.ReporterMemberID,
		Reason:           row.Reason,
		CreatedAt:        row.CreatedAt,
	}
}
