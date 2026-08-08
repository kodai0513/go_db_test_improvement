package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// MenuRepository は content.MenuRepository の実装。
type MenuRepository struct {
	q *sqlc.Queries
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{q: sqlc.New(db)}
}

var _ content.MenuRepository = (*MenuRepository)(nil)

func (r *MenuRepository) Create(ctx context.Context, m *content.Menu) error {
	row, err := r.q.CreateMenu(ctx, sqlc.CreateMenuParams{
		ID:   m.ID,
		Name: m.Name,
	})
	if err != nil {
		return err
	}
	*m = toMenuDomain(row)
	return nil
}

func (r *MenuRepository) Get(ctx context.Context, id uuid.UUID) (*content.Menu, error) {
	row, err := r.q.GetMenu(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toMenuDomain(row)
	return &m, nil
}

func (r *MenuRepository) List(ctx context.Context) ([]*content.Menu, error) {
	rows, err := r.q.ListMenus(ctx)
	if err != nil {
		return nil, err
	}
	menus := make([]*content.Menu, 0, len(rows))
	for _, row := range rows {
		m := toMenuDomain(row)
		menus = append(menus, &m)
	}
	return menus, nil
}

func (r *MenuRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMenu(ctx, id)
}

func toMenuDomain(row sqlc.Menu) content.Menu {
	return content.Menu{
		ID:   row.ID,
		Name: row.Name,
	}
}
