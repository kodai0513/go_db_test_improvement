package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// TicketCategoryRepository は support.TicketCategoryRepository の実装。
type TicketCategoryRepository struct {
	q *sqlc.Queries
}

func NewTicketCategoryRepository(db *sql.DB) *TicketCategoryRepository {
	return &TicketCategoryRepository{q: sqlc.New(db)}
}

var _ support.TicketCategoryRepository = (*TicketCategoryRepository)(nil)

func (r *TicketCategoryRepository) Create(ctx context.Context, c *support.TicketCategory) error {
	row, err := r.q.CreateTicketCategory(ctx, sqlc.CreateTicketCategoryParams{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return err
	}
	*c = toTicketCategoryDomain(row)
	return nil
}

func (r *TicketCategoryRepository) Get(ctx context.Context, id uuid.UUID) (*support.TicketCategory, error) {
	row, err := r.q.GetTicketCategory(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toTicketCategoryDomain(row)
	return &c, nil
}

func (r *TicketCategoryRepository) List(ctx context.Context) ([]*support.TicketCategory, error) {
	rows, err := r.q.ListTicketCategories(ctx)
	if err != nil {
		return nil, err
	}
	categories := make([]*support.TicketCategory, 0, len(rows))
	for _, row := range rows {
		c := toTicketCategoryDomain(row)
		categories = append(categories, &c)
	}
	return categories, nil
}

func (r *TicketCategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTicketCategory(ctx, id)
}

func toTicketCategoryDomain(row sqlc.TicketCategory) support.TicketCategory {
	return support.TicketCategory{
		ID:   row.ID,
		Name: row.Name,
	}
}
