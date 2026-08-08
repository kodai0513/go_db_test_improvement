package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewTagRepository は review.ReviewTagRepository の実装。
type ReviewTagRepository struct {
	q *sqlc.Queries
}

func NewReviewTagRepository(db *sql.DB) *ReviewTagRepository {
	return &ReviewTagRepository{q: sqlc.New(db)}
}

var _ review.ReviewTagRepository = (*ReviewTagRepository)(nil)

func (r *ReviewTagRepository) Create(ctx context.Context, t *review.ReviewTag) error {
	row, err := r.q.CreateReviewTag(ctx, sqlc.CreateReviewTagParams{
		ID:   t.ID,
		Name: t.Name,
	})
	if err != nil {
		return err
	}
	*t = toReviewTagDomain(row)
	return nil
}

func (r *ReviewTagRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewTag, error) {
	row, err := r.q.GetReviewTag(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toReviewTagDomain(row)
	return &t, nil
}

func (r *ReviewTagRepository) List(ctx context.Context) ([]*review.ReviewTag, error) {
	rows, err := r.q.ListReviewTags(ctx)
	if err != nil {
		return nil, err
	}
	tags := make([]*review.ReviewTag, 0, len(rows))
	for _, row := range rows {
		t := toReviewTagDomain(row)
		tags = append(tags, &t)
	}
	return tags, nil
}

func (r *ReviewTagRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewTag(ctx, id)
}

func toReviewTagDomain(row sqlc.ReviewTag) review.ReviewTag {
	return review.ReviewTag{
		ID:   row.ID,
		Name: row.Name,
	}
}
