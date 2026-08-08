package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// Repository は pricing.PriceListRepository の実装。
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ pricing.PriceListRepository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, p *pricing.PriceList) error {
	row, err := r.q.CreatePriceList(ctx, sqlc.CreatePriceListParams{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
	})
	if err != nil {
		return err
	}
	*p = toPriceListDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*pricing.PriceList, error) {
	row, err := r.q.GetPriceList(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPriceListDomain(row)
	return &p, nil
}

func (r *Repository) List(ctx context.Context) ([]*pricing.PriceList, error) {
	rows, err := r.q.ListPriceLists(ctx)
	if err != nil {
		return nil, err
	}
	lists := make([]*pricing.PriceList, 0, len(rows))
	for _, row := range rows {
		p := toPriceListDomain(row)
		lists = append(lists, &p)
	}
	return lists, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePriceList(ctx, id)
}

func toPriceListDomain(row sqlc.PriceList) pricing.PriceList {
	return pricing.PriceList{
		ID:        row.ID,
		Name:      row.Name,
		CreatedAt: row.CreatedAt,
	}
}
