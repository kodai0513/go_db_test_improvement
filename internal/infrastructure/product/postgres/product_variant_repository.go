package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductVariantRepository は product.ProductVariantRepository の実装。
type ProductVariantRepository struct {
	q *sqlc.Queries
}

func NewProductVariantRepository(db *sql.DB) *ProductVariantRepository {
	return &ProductVariantRepository{q: sqlc.New(db)}
}

var _ product.ProductVariantRepository = (*ProductVariantRepository)(nil)

func (r *ProductVariantRepository) Create(ctx context.Context, v *product.ProductVariant) error {
	row, err := r.q.CreateProductVariant(ctx, sqlc.CreateProductVariantParams{
		ID:        v.ID,
		ProductID: v.ProductID,
		SKU:       v.SKU,
		PriceYen:  v.PriceYen,
		CreatedAt: v.CreatedAt,
	})
	if err != nil {
		return err
	}
	*v = toProductVariantDomain(row)
	return nil
}

func (r *ProductVariantRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductVariant, error) {
	row, err := r.q.GetProductVariant(ctx, id)
	if err != nil {
		return nil, err
	}
	v := toProductVariantDomain(row)
	return &v, nil
}

func (r *ProductVariantRepository) List(ctx context.Context) ([]*product.ProductVariant, error) {
	rows, err := r.q.ListProductVariants(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductVariant, 0, len(rows))
	for _, row := range rows {
		v := toProductVariantDomain(row)
		items = append(items, &v)
	}
	return items, nil
}

func (r *ProductVariantRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*product.ProductVariant, error) {
	rows, err := r.q.ListProductVariantsByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductVariant, 0, len(rows))
	for _, row := range rows {
		v := toProductVariantDomain(row)
		items = append(items, &v)
	}
	return items, nil
}

func (r *ProductVariantRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductVariant(ctx, id)
}

func toProductVariantDomain(row sqlc.ProductVariant) product.ProductVariant {
	return product.ProductVariant{
		ID:        row.ID,
		ProductID: row.ProductID,
		SKU:       row.SKU,
		PriceYen:  row.PriceYen,
		CreatedAt: row.CreatedAt,
	}
}
