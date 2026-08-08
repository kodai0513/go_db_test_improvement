package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorReviewRepository は merchant.VendorReviewRepository の実装。
type VendorReviewRepository struct {
	q *sqlc.Queries
}

func NewVendorReviewRepository(db *sql.DB) *VendorReviewRepository {
	return &VendorReviewRepository{q: sqlc.New(db)}
}

var _ merchant.VendorReviewRepository = (*VendorReviewRepository)(nil)

func (r *VendorReviewRepository) Create(ctx context.Context, rv *merchant.VendorReview) error {
	row, err := r.q.CreateVendorReview(ctx, sqlc.CreateVendorReviewParams{
		ID:       rv.ID,
		VendorID: rv.VendorID,
		MemberID: rv.MemberID,
		Rating:   rv.Rating,
		Comment:  rv.Comment,
	})
	if err != nil {
		return err
	}
	*rv = toVendorReviewDomain(row)
	return nil
}

func (r *VendorReviewRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorReview, error) {
	row, err := r.q.GetVendorReview(ctx, id)
	if err != nil {
		return nil, err
	}
	rv := toVendorReviewDomain(row)
	return &rv, nil
}

func (r *VendorReviewRepository) List(ctx context.Context) ([]*merchant.VendorReview, error) {
	rows, err := r.q.ListVendorReviews(ctx)
	if err != nil {
		return nil, err
	}
	reviews := make([]*merchant.VendorReview, 0, len(rows))
	for _, row := range rows {
		rv := toVendorReviewDomain(row)
		reviews = append(reviews, &rv)
	}
	return reviews, nil
}

func (r *VendorReviewRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorReview, error) {
	rows, err := r.q.ListVendorReviewsByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	reviews := make([]*merchant.VendorReview, 0, len(rows))
	for _, row := range rows {
		rv := toVendorReviewDomain(row)
		reviews = append(reviews, &rv)
	}
	return reviews, nil
}

func (r *VendorReviewRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorReview(ctx, id)
}

func toVendorReviewDomain(row sqlc.VendorReview) merchant.VendorReview {
	return merchant.VendorReview{
		ID:       row.ID,
		VendorID: row.VendorID,
		MemberID: row.MemberID,
		Rating:   row.Rating,
		Comment:  row.Comment,
	}
}
