package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewSummaryRepository は review.ReviewSummaryRepository の実装。
type ReviewSummaryRepository struct {
	q *sqlc.Queries
}

func NewReviewSummaryRepository(db *sql.DB) *ReviewSummaryRepository {
	return &ReviewSummaryRepository{q: sqlc.New(db)}
}

var _ review.ReviewSummaryRepository = (*ReviewSummaryRepository)(nil)

func (r *ReviewSummaryRepository) Create(ctx context.Context, s *review.ReviewSummary) error {
	row, err := r.q.CreateReviewSummary(ctx, sqlc.CreateReviewSummaryParams{
		ID:            s.ID,
		ProductID:     s.ProductID,
		AverageRating: s.AverageRating,
		ReviewCount:   s.ReviewCount,
	})
	if err != nil {
		return err
	}
	*s = toReviewSummaryDomain(row)
	return nil
}

func (r *ReviewSummaryRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewSummary, error) {
	row, err := r.q.GetReviewSummary(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toReviewSummaryDomain(row)
	return &s, nil
}

func (r *ReviewSummaryRepository) List(ctx context.Context) ([]*review.ReviewSummary, error) {
	rows, err := r.q.ListReviewSummaries(ctx)
	if err != nil {
		return nil, err
	}
	summaries := make([]*review.ReviewSummary, 0, len(rows))
	for _, row := range rows {
		s := toReviewSummaryDomain(row)
		summaries = append(summaries, &s)
	}
	return summaries, nil
}

func (r *ReviewSummaryRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*review.ReviewSummary, error) {
	rows, err := r.q.ListReviewSummariesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	summaries := make([]*review.ReviewSummary, 0, len(rows))
	for _, row := range rows {
		s := toReviewSummaryDomain(row)
		summaries = append(summaries, &s)
	}
	return summaries, nil
}

func (r *ReviewSummaryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewSummary(ctx, id)
}

func toReviewSummaryDomain(row sqlc.ReviewSummary) review.ReviewSummary {
	return review.ReviewSummary{
		ID:            row.ID,
		ProductID:     row.ProductID,
		AverageRating: row.AverageRating,
		ReviewCount:   row.ReviewCount,
	}
}
