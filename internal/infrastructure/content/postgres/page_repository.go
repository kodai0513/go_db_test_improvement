package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres/sqlc"
)

// PageRepository は content.PageRepository の実装。
type PageRepository struct {
	q *sqlc.Queries
}

func NewPageRepository(db *sql.DB) *PageRepository {
	return &PageRepository{q: sqlc.New(db)}
}

var _ content.PageRepository = (*PageRepository)(nil)

func (r *PageRepository) Create(ctx context.Context, p *content.Page) error {
	row, err := r.q.CreatePage(ctx, sqlc.CreatePageParams{
		ID:    p.ID,
		Slug:  p.Slug,
		Title: p.Title,
		Body:  p.Body,
	})
	if err != nil {
		return err
	}
	*p = toPageDomain(row)
	return nil
}

func (r *PageRepository) Get(ctx context.Context, id uuid.UUID) (*content.Page, error) {
	row, err := r.q.GetPage(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPageDomain(row)
	return &p, nil
}

func (r *PageRepository) List(ctx context.Context) ([]*content.Page, error) {
	rows, err := r.q.ListPages(ctx)
	if err != nil {
		return nil, err
	}
	pages := make([]*content.Page, 0, len(rows))
	for _, row := range rows {
		p := toPageDomain(row)
		pages = append(pages, &p)
	}
	return pages, nil
}

func (r *PageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePage(ctx, id)
}

func toPageDomain(row sqlc.Page) content.Page {
	return content.Page{
		ID:    row.ID,
		Slug:  row.Slug,
		Title: row.Title,
		Body:  row.Body,
	}
}
