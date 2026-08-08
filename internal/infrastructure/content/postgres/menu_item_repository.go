package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// MenuItemRepository は content.MenuItemRepository の実装。
type MenuItemRepository struct {
	q *sqlc.Queries
}

func NewMenuItemRepository(db *sql.DB) *MenuItemRepository {
	return &MenuItemRepository{q: sqlc.New(db)}
}

var _ content.MenuItemRepository = (*MenuItemRepository)(nil)

func (r *MenuItemRepository) Create(ctx context.Context, m *content.MenuItem) error {
	row, err := r.q.CreateMenuItem(ctx, sqlc.CreateMenuItemParams{
		ID:        m.ID,
		MenuID:    m.MenuID,
		Label:     m.Label,
		Url:       m.URL,
		SortOrder: m.SortOrder,
	})
	if err != nil {
		return err
	}
	*m = toMenuItemDomain(row)
	return nil
}

func (r *MenuItemRepository) Get(ctx context.Context, id uuid.UUID) (*content.MenuItem, error) {
	row, err := r.q.GetMenuItem(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toMenuItemDomain(row)
	return &m, nil
}

func (r *MenuItemRepository) List(ctx context.Context) ([]*content.MenuItem, error) {
	rows, err := r.q.ListMenuItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*content.MenuItem, 0, len(rows))
	for _, row := range rows {
		m := toMenuItemDomain(row)
		items = append(items, &m)
	}
	return items, nil
}

func (r *MenuItemRepository) ListByMenuID(ctx context.Context, menuID uuid.UUID) ([]*content.MenuItem, error) {
	rows, err := r.q.ListMenuItemsByMenuID(ctx, menuID)
	if err != nil {
		return nil, err
	}
	items := make([]*content.MenuItem, 0, len(rows))
	for _, row := range rows {
		m := toMenuItemDomain(row)
		items = append(items, &m)
	}
	return items, nil
}

func (r *MenuItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMenuItem(ctx, id)
}

func toMenuItemDomain(row sqlc.MenuItem) content.MenuItem {
	return content.MenuItem{
		ID:        row.ID,
		MenuID:    row.MenuID,
		Label:     row.Label,
		URL:       row.Url,
		SortOrder: row.SortOrder,
	}
}
