package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// FaqCategoryRepository は support.FaqCategoryRepository の実装。
type FaqCategoryRepository struct {
	q *sqlc.Queries
}

func NewFaqCategoryRepository(db *sql.DB) *FaqCategoryRepository {
	return &FaqCategoryRepository{q: sqlc.New(db)}
}

var _ support.FaqCategoryRepository = (*FaqCategoryRepository)(nil)

func (r *FaqCategoryRepository) Create(ctx context.Context, c *support.FaqCategory) error {
	row, err := r.q.CreateFaqCategory(ctx, sqlc.CreateFaqCategoryParams{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return err
	}
	*c = toFaqCategoryDomain(row)
	return nil
}

func (r *FaqCategoryRepository) Get(ctx context.Context, id uuid.UUID) (*support.FaqCategory, error) {
	row, err := r.q.GetFaqCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toFaqCategoryDomain(row)
	return &c, nil
}

func (r *FaqCategoryRepository) List(ctx context.Context) ([]*support.FaqCategory, error) {
	rows, err := r.q.ListFaqCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*support.FaqCategory, 0, len(rows))
	for _, row := range rows {
		c := toFaqCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *FaqCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteFaqCategory(ctx, id)
}

func toFaqCategoryDomain(row sqlc.FaqCategory) support.FaqCategory {
	return support.FaqCategory{
		ID:   row.ID,
		Name: row.Name,
	}
}
