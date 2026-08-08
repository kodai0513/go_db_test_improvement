package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres/sqlc"
)

// CategoryRepository は product.CategoryRepository の実装。
type CategoryRepository struct {
	q *sqlc.Queries
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{q: sqlc.New(db)}
}

var _ product.CategoryRepository = (*CategoryRepository)(nil)

func (r *CategoryRepository) Create(ctx context.Context, c *product.Category) error {
	row, err := r.q.CreateCategory(ctx, sqlc.CreateCategoryParams{
		ID:               c.ID,
		Name:             c.Name,
		ParentCategoryID: c.ParentCategoryID,
		CreatedAt:        c.CreatedAt,
	})
	if err != nil {
		return err
	}
	*c = toCategoryDomain(row)
	return nil
}

func (r *CategoryRepository) Get(ctx context.Context, id uuid.UUID) (*product.Category, error) {
	row, err := r.q.GetCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCategoryDomain(row)
	return &c, nil
}

func (r *CategoryRepository) List(ctx context.Context) ([]*product.Category, error) {
	rows, err := r.q.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*product.Category, 0, len(rows))
	for _, row := range rows {
		c := toCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *CategoryRepository) ListByParentCategoryID(ctx context.Context, parentCategoryID uuid.UUID) ([]*product.Category, error) {
	rows, err := r.q.ListCategoriesByParentCategoryID(ctx, parentCategoryID)
	if err != nil {
		return nil, err
	}
	categories := make([]*product.Category, 0, len(rows))
	for _, row := range rows {
		c := toCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCategory(ctx, id)
}

func toCategoryDomain(row sqlc.Category) product.Category {
	return product.Category{
		ID:               row.ID,
		Name:             row.Name,
		ParentCategoryID: row.ParentCategoryID,
		CreatedAt:        row.CreatedAt,
	}
}
