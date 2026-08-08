package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres/sqlc"
)

// Repository implements review.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ review.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, rv *review.Review) error {
	row, err := r.q.CreateReview(ctx, sqlc.CreateReviewParams{
		ID:        rv.ID,
		ProductID: rv.ProductID,
		MemberID:  rv.MemberID,
		Rating:    rv.Rating,
		Comment:   rv.Comment,
		CreatedAt: rv.CreatedAt,
	})
	if err != nil {
		return err
	}
	*rv = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*review.Review, error) {
	row, err := r.q.GetReview(ctx, id)
	if err != nil {
		return nil, err
	}
	rv := toDomain(row)
	return &rv, nil
}

func (r *Repository) List(ctx context.Context) ([]*review.Review, error) {
	rows, err := r.q.ListReviews(ctx)
	if err != nil {
		return nil, err
	}
	reviews := make([]*review.Review, 0, len(rows))
	for _, row := range rows {
		rv := toDomain(row)
		reviews = append(reviews, &rv)
	}
	return reviews, nil
}

func toDomain(row sqlc.Review) review.Review {
	return review.Review{
		ID:        row.ID,
		ProductID: row.ProductID,
		MemberID:  row.MemberID,
		Rating:    row.Rating,
		Comment:   row.Comment,
		CreatedAt: row.CreatedAt,
	}
}
