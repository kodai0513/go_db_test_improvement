package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// BrandRepository は product.BrandRepository の実装。
type BrandRepository struct {
	q *sqlc.Queries
}

func NewBrandRepository(db *sql.DB) *BrandRepository {
	return &BrandRepository{q: sqlc.New(db)}
}

var _ product.BrandRepository = (*BrandRepository)(nil)

func (r *BrandRepository) Create(ctx context.Context, b *product.Brand) error {
	row, err := r.q.CreateBrand(ctx, sqlc.CreateBrandParams{
		ID:        b.ID,
		Name:      b.Name,
		CreatedAt: b.CreatedAt,
	})
	if err != nil {
		return err
	}
	*b = toBrandDomain(row)
	return nil
}

func (r *BrandRepository) Get(ctx context.Context, id uuid.UUID) (*product.Brand, error) {
	row, err := r.q.GetBrand(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toBrandDomain(row)
	return &b, nil
}

func (r *BrandRepository) List(ctx context.Context) ([]*product.Brand, error) {
	rows, err := r.q.ListBrands(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.Brand, 0, len(rows))
	for _, row := range rows {
		b := toBrandDomain(row)
		items = append(items, &b)
	}
	return items, nil
}

func (r *BrandRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBrand(ctx, id)
}

func toBrandDomain(row sqlc.Brand) product.Brand {
	return product.Brand{
		ID:        row.ID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
	}
}
