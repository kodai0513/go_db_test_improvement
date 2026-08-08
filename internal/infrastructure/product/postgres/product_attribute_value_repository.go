package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductAttributeValueRepository は product.ProductAttributeValueRepository の実装。
type ProductAttributeValueRepository struct {
	q *sqlc.Queries
}

func NewProductAttributeValueRepository(db *sql.DB) *ProductAttributeValueRepository {
	return &ProductAttributeValueRepository{q: sqlc.New(db)}
}

var _ product.ProductAttributeValueRepository = (*ProductAttributeValueRepository)(nil)

func (r *ProductAttributeValueRepository) Create(ctx context.Context, v *product.ProductAttributeValue) error {
	row, err := r.q.CreateProductAttributeValue(ctx, sqlc.CreateProductAttributeValueParams{
		ID:                  v.ID,
		ProductID:           v.ProductID,
		ProductAttributeID:  v.ProductAttributeID,
		Value:               v.Value,
	})
	if err != nil {
		return err
	}
	*v = toProductAttributeValueDomain(row)
	return nil
}

func (r *ProductAttributeValueRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductAttributeValue, error) {
	row, err := r.q.GetProductAttributeValue(ctx, id)
	if err != nil {
		return nil, err
	}
	v := toProductAttributeValueDomain(row)
	return &v, nil
}

func (r *ProductAttributeValueRepository) List(ctx context.Context) ([]*product.ProductAttributeValue, error) {
	rows, err := r.q.ListProductAttributeValues(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductAttributeValue, 0, len(rows))
	for _, row := range rows {
		v := toProductAttributeValueDomain(row)
		items = append(items, &v)
	}
	return items, nil
}

func (r *ProductAttributeValueRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*product.ProductAttributeValue, error) {
	rows, err := r.q.ListProductAttributeValuesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductAttributeValue, 0, len(rows))
	for _, row := range rows {
		v := toProductAttributeValueDomain(row)
		items = append(items, &v)
	}
	return items, nil
}

func (r *ProductAttributeValueRepository) ListByProductAttributeID(ctx context.Context, productAttributeID uuid.UUID) ([]*product.ProductAttributeValue, error) {
	rows, err := r.q.ListProductAttributeValuesByProductAttributeID(ctx, productAttributeID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductAttributeValue, 0, len(rows))
	for _, row := range rows {
		v := toProductAttributeValueDomain(row)
		items = append(items, &v)
	}
	return items, nil
}

func (r *ProductAttributeValueRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductAttributeValue(ctx, id)
}

func toProductAttributeValueDomain(row sqlc.ProductAttributeValue) product.ProductAttributeValue {
	return product.ProductAttributeValue{
		ID:                  row.ID,
		ProductID:           row.ProductID,
		ProductAttributeID:  row.ProductAttributeID,
		Value:               row.Value,
	}
}
