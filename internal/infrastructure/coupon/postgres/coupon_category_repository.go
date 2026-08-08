package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CouponCategoryRepository は coupon.CouponCategoryRepository の実装。
type CouponCategoryRepository struct {
	q *sqlc.Queries
}

func NewCouponCategoryRepository(db *sql.DB) *CouponCategoryRepository {
	return &CouponCategoryRepository{q: sqlc.New(db)}
}

var _ coupon.CouponCategoryRepository = (*CouponCategoryRepository)(nil)

func (r *CouponCategoryRepository) Create(ctx context.Context, c *coupon.CouponCategory) error {
	row, err := r.q.CreateCouponCategory(ctx, sqlc.CreateCouponCategoryParams{
		ID:         c.ID,
		CouponID:   c.CouponID,
		CategoryID: c.CategoryID,
	})
	if err != nil {
		return err
	}
	*c = toCouponCategoryDomain(row)
	return nil
}

func (r *CouponCategoryRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.CouponCategory, error) {
	row, err := r.q.GetCouponCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCouponCategoryDomain(row)
	return &c, nil
}

func (r *CouponCategoryRepository) List(ctx context.Context) ([]*coupon.CouponCategory, error) {
	rows, err := r.q.ListCouponCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*coupon.CouponCategory, 0, len(rows))
	for _, row := range rows {
		c := toCouponCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *CouponCategoryRepository) ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*coupon.CouponCategory, error) {
	rows, err := r.q.ListCouponCategoriesByCouponID(ctx, couponID)
	if err != nil {
		return nil, err
	}
	categories := make([]*coupon.CouponCategory, 0, len(rows))
	for _, row := range rows {
		c := toCouponCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *CouponCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCouponCategory(ctx, id)
}

func toCouponCategoryDomain(row sqlc.CouponCategory) coupon.CouponCategory {
	return coupon.CouponCategory{
		ID:         row.ID,
		CouponID:   row.CouponID,
		CategoryID: row.CategoryID,
	}
}
