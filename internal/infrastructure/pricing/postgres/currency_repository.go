package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// CurrencyRepository は pricing.CurrencyRepository の実装。
type CurrencyRepository struct {
	q *sqlc.Queries
}

func NewCurrencyRepository(db *sql.DB) *CurrencyRepository {
	return &CurrencyRepository{q: sqlc.New(db)}
}

var _ pricing.CurrencyRepository = (*CurrencyRepository)(nil)

func (r *CurrencyRepository) Create(ctx context.Context, c *pricing.Currency) error {
	row, err := r.q.CreateCurrency(ctx, sqlc.CreateCurrencyParams{
		ID:   c.ID,
		Code: c.Code,
	})
	if err != nil {
		return err
	}
	*c = toCurrencyDomain(row)
	return nil
}

func (r *CurrencyRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.Currency, error) {
	row, err := r.q.GetCurrency(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCurrencyDomain(row)
	return &c, nil
}

func (r *CurrencyRepository) List(ctx context.Context) ([]*pricing.Currency, error) {
	rows, err := r.q.ListCurrencies(ctx)
	if err != nil {
		return nil, err
	}
	currencies := make([]*pricing.Currency, 0, len(rows))
	for _, row := range rows {
		c := toCurrencyDomain(row)
		currencies = append(currencies, &c)
	}
	return currencies, nil
}

func (r *CurrencyRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCurrency(ctx, id)
}

func toCurrencyDomain(row sqlc.Currency) pricing.Currency {
	return pricing.Currency{
		ID:   row.ID,
		Code: row.Code,
	}
}
