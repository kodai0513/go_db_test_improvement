package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// ProductCategoryRepository は product.ProductCategoryRepository の実装。
type ProductCategoryRepository struct {
	q *sqlc.Queries
}

func NewProductCategoryRepository(db *sql.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{q: sqlc.New(db)}
}

var _ product.ProductCategoryRepository = (*ProductCategoryRepository)(nil)

func (r *ProductCategoryRepository) Create(ctx context.Context, pc *product.ProductCategory) error {
	row, err := r.q.CreateProductCategory(ctx, sqlc.CreateProductCategoryParams{
		ID:         pc.ID,
		ProductID:  pc.ProductID,
		CategoryID: pc.CategoryID,
		CreatedAt:  pc.CreatedAt,
	})
	if err != nil {
		return err
	}
	*pc = toProductCategoryDomain(row)
	return nil
}

func (r *ProductCategoryRepository) Get(ctx context.Context, id uuid.UUID) (*product.ProductCategory, error) {
	row, err := r.q.GetProductCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	pc := toProductCategoryDomain(row)
	return &pc, nil
}

func (r *ProductCategoryRepository) List(ctx context.Context) ([]*product.ProductCategory, error) {
	rows, err := r.q.ListProductCategories(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductCategory, 0, len(rows))
	for _, row := range rows {
		pc := toProductCategoryDomain(row)
		items = append(items, &pc)
	}
	return items, nil
}

func (r *ProductCategoryRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*product.ProductCategory, error) {
	rows, err := r.q.ListProductCategoriesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductCategory, 0, len(rows))
	for _, row := range rows {
		pc := toProductCategoryDomain(row)
		items = append(items, &pc)
	}
	return items, nil
}

func (r *ProductCategoryRepository) ListByCategoryID(ctx context.Context, categoryID uuid.UUID) ([]*product.ProductCategory, error) {
	rows, err := r.q.ListProductCategoriesByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}
	items := make([]*product.ProductCategory, 0, len(rows))
	for _, row := range rows {
		pc := toProductCategoryDomain(row)
		items = append(items, &pc)
	}
	return items, nil
}

func (r *ProductCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteProductCategory(ctx, id)
}

func toProductCategoryDomain(row sqlc.ProductCategory) product.ProductCategory {
	return product.ProductCategory{
		ID:         row.ID,
		ProductID:  row.ProductID,
		CategoryID: row.CategoryID,
		CreatedAt:  row.CreatedAt,
	}
}
