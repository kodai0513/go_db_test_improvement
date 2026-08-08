package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// BlogCategoryRepository は content.BlogCategoryRepository の実装。
type BlogCategoryRepository struct {
	q *sqlc.Queries
}

func NewBlogCategoryRepository(db *sql.DB) *BlogCategoryRepository {
	return &BlogCategoryRepository{q: sqlc.New(db)}
}

var _ content.BlogCategoryRepository = (*BlogCategoryRepository)(nil)

func (r *BlogCategoryRepository) Create(ctx context.Context, c *content.BlogCategory) error {
	row, err := r.q.CreateBlogCategory(ctx, sqlc.CreateBlogCategoryParams{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return err
	}
	*c = toBlogCategoryDomain(row)
	return nil
}

func (r *BlogCategoryRepository) Get(ctx context.Context, id uuid.UUID) (*content.BlogCategory, error) {
	row, err := r.q.GetBlogCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toBlogCategoryDomain(row)
	return &c, nil
}

func (r *BlogCategoryRepository) List(ctx context.Context) ([]*content.BlogCategory, error) {
	rows, err := r.q.ListBlogCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*content.BlogCategory, 0, len(rows))
	for _, row := range rows {
		c := toBlogCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *BlogCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBlogCategory(ctx, id)
}

func toBlogCategoryDomain(row sqlc.BlogCategory) content.BlogCategory {
	return content.BlogCategory{
		ID:   row.ID,
		Name: row.Name,
	}
}
