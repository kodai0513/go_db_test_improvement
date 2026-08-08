package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres/sqlc"
)

// VendorCategoryRepository は merchant.VendorCategoryRepository の実装。
type VendorCategoryRepository struct {
	q *sqlc.Queries
}

func NewVendorCategoryRepository(db *sql.DB) *VendorCategoryRepository {
	return &VendorCategoryRepository{q: sqlc.New(db)}
}

var _ merchant.VendorCategoryRepository = (*VendorCategoryRepository)(nil)

func (r *VendorCategoryRepository) Create(ctx context.Context, c *merchant.VendorCategory) error {
	row, err := r.q.CreateVendorCategory(ctx, sqlc.CreateVendorCategoryParams{
		ID:         c.ID,
		VendorID:   c.VendorID,
		CategoryID: c.CategoryID,
	})
	if err != nil {
		return err
	}
	*c = toVendorCategoryDomain(row)
	return nil
}

func (r *VendorCategoryRepository) Get(ctx context.Context, id uuid.UUID) (*merchant.VendorCategory, error) {
	row, err := r.q.GetVendorCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toVendorCategoryDomain(row)
	return &c, nil
}

func (r *VendorCategoryRepository) List(ctx context.Context) ([]*merchant.VendorCategory, error) {
	rows, err := r.q.ListVendorCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*merchant.VendorCategory, 0, len(rows))
	for _, row := range rows {
		c := toVendorCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *VendorCategoryRepository) ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*merchant.VendorCategory, error) {
	rows, err := r.q.ListVendorCategoriesByVendorID(ctx, vendorID)
	if err != nil {
		return nil, err
	}
	categories := make([]*merchant.VendorCategory, 0, len(rows))
	for _, row := range rows {
		c := toVendorCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *VendorCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteVendorCategory(ctx, id)
}

func toVendorCategoryDomain(row sqlc.VendorCategory) merchant.VendorCategory {
	return merchant.VendorCategory{
		ID:         row.ID,
		VendorID:   row.VendorID,
		CategoryID: row.CategoryID,
	}
}
