package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewTagAssignmentRepository は review.ReviewTagAssignmentRepository の実装。
type ReviewTagAssignmentRepository struct {
	q *sqlc.Queries
}

func NewReviewTagAssignmentRepository(db *sql.DB) *ReviewTagAssignmentRepository {
	return &ReviewTagAssignmentRepository{q: sqlc.New(db)}
}

var _ review.ReviewTagAssignmentRepository = (*ReviewTagAssignmentRepository)(nil)

func (r *ReviewTagAssignmentRepository) Create(ctx context.Context, a *review.ReviewTagAssignment) error {
	row, err := r.q.CreateReviewTagAssignment(ctx, sqlc.CreateReviewTagAssignmentParams{
		ID:          a.ID,
		ReviewID:    a.ReviewID,
		ReviewTagID: a.ReviewTagID,
	})
	if err != nil {
		return err
	}
	*a = toReviewTagAssignmentDomain(row)
	return nil
}

func (r *ReviewTagAssignmentRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewTagAssignment, error) {
	row, err := r.q.GetReviewTagAssignment(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toReviewTagAssignmentDomain(row)
	return &a, nil
}

func (r *ReviewTagAssignmentRepository) List(ctx context.Context) ([]*review.ReviewTagAssignment, error) {
	rows, err := r.q.ListReviewTagAssignments(ctx)
	if err != nil {
		return nil, err
	}
	assignments := make([]*review.ReviewTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toReviewTagAssignmentDomain(row)
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *ReviewTagAssignmentRepository) ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*review.ReviewTagAssignment, error) {
	rows, err := r.q.ListReviewTagAssignmentsByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	assignments := make([]*review.ReviewTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toReviewTagAssignmentDomain(row)
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *ReviewTagAssignmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewTagAssignment(ctx, id)
}

func toReviewTagAssignmentDomain(row sqlc.ReviewTagAssignment) review.ReviewTagAssignment {
	return review.ReviewTagAssignment{
		ID:          row.ID,
		ReviewID:    row.ReviewID,
		ReviewTagID: row.ReviewTagID,
	}
}
