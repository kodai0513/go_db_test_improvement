package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductAttributeRepository は product.ProductAttributeRepository の実装。
type ProductAttributeRepository struct {
	q *sqlc.Queries
}

func NewProductAttributeRepository(db *sql.DB) *ProductAttributeRepository {
	return &ProductAttributeRepository{q: sqlc.New(db)}
}

var _ product.ProductAttributeRepository = (*ProductAttributeRepository)(nil)

func (r *ProductAttributeRepository) Create(ctx context.Context, a *product.ProductAttribute) error {
	row, err := r.q.CreateProductAttribute(ctx, sqlc.CreateProductAttributeParams{
		ID:        a.ID,
		Name:      a.Name,
		CreatedAt: a.CreatedAt,
	})
	if err != nil {
		return err
	}
	*a = toProductAttributeDomain(row)
	return nil
}

func (r *ProductAttributeRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductAttribute, error) {
	row, err := r.q.GetProductAttribute(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toProductAttributeDomain(row)
	return &a, nil
}

func (r *ProductAttributeRepository) List(ctx context.Context) ([]*product.ProductAttribute, error) {
	rows, err := r.q.ListProductAttributes(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductAttribute, 0, len(rows))
	for _, row := range rows {
		a := toProductAttributeDomain(row)
		items = append(items, &a)
	}
	return items, nil
}

func (r *ProductAttributeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductAttribute(ctx, id)
}

func toProductAttributeDomain(row sqlc.ProductAttribute) product.ProductAttribute {
	return product.ProductAttribute{
		ID:        row.ID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
	}
}
