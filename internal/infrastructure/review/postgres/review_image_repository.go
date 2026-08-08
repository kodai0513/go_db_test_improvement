package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// ReviewImageRepository は review.ReviewImageRepository の実装。
type ReviewImageRepository struct {
	q *sqlc.Queries
}

func NewReviewImageRepository(db *sql.DB) *ReviewImageRepository {
	return &ReviewImageRepository{q: sqlc.New(db)}
}

var _ review.ReviewImageRepository = (*ReviewImageRepository)(nil)

func (r *ReviewImageRepository) Create(ctx context.Context, i *review.ReviewImage) error {
	row, err := r.q.CreateReviewImage(ctx, sqlc.CreateReviewImageParams{
		ID:       i.ID,
		ReviewID: i.ReviewID,
		Url:      i.URL,
	})
	if err != nil {
		return err
	}
	*i = toReviewImageDomain(row)
	return nil
}

func (r *ReviewImageRepository) Get(ctx context.Context, id uuid.UUID) (*review.ReviewImage, error) {
	row, err := r.q.GetReviewImage(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toReviewImageDomain(row)
	return &i, nil
}

func (r *ReviewImageRepository) List(ctx context.Context) ([]*review.ReviewImage, error) {
	rows, err := r.q.ListReviewImages(ctx)
	if err != nil {
		return nil, err
	}
	images := make([]*review.ReviewImage, 0, len(rows))
	for _, row := range rows {
		i := toReviewImageDomain(row)
		images = append(images, &i)
	}
	return images, nil
}

func (r *ReviewImageRepository) ListByReviewID(ctx context.Context, reviewID uuid.UUID) ([]*review.ReviewImage, error) {
	rows, err := r.q.ListReviewImagesByReviewID(ctx, reviewID)
	if err != nil {
		return nil, err
	}
	images := make([]*review.ReviewImage, 0, len(rows))
	for _, row := range rows {
		i := toReviewImageDomain(row)
		images = append(images, &i)
	}
	return images, nil
}

func (r *ReviewImageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReviewImage(ctx, id)
}

func toReviewImageDomain(row sqlc.ReviewImage) review.ReviewImage {
	return review.ReviewImage{
		ID:       row.ID,
		ReviewID: row.ReviewID,
		URL:      row.Url,
	}
}
