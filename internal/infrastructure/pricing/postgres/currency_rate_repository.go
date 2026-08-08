package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// CurrencyRateRepository は pricing.CurrencyRateRepository の実装。
type CurrencyRateRepository struct {
	q *sqlc.Queries
}

func NewCurrencyRateRepository(db *sql.DB) *CurrencyRateRepository {
	return &CurrencyRateRepository{q: sqlc.New(db)}
}

var _ pricing.CurrencyRateRepository = (*CurrencyRateRepository)(nil)

func (r *CurrencyRateRepository) Create(ctx context.Context, c *pricing.CurrencyRate) error {
	row, err := r.q.CreateCurrencyRate(ctx, sqlc.CreateCurrencyRateParams{
		ID:          c.ID,
		CurrencyID:  c.CurrencyID,
		RateToJpy:   c.RateToJPY,
		EffectiveAt: c.EffectiveAt,
	})
	if err != nil {
		return err
	}
	*c = toCurrencyRateDomain(row)
	return nil
}

func (r *CurrencyRateRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.CurrencyRate, error) {
	row, err := r.q.GetCurrencyRate(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCurrencyRateDomain(row)
	return &c, nil
}

func (r *CurrencyRateRepository) List(ctx context.Context) ([]*pricing.CurrencyRate, error) {
	rows, err := r.q.ListCurrencyRates(ctx)
	if err != nil {
		return nil, err
	}
	rates := make([]*pricing.CurrencyRate, 0, len(rows))
	for _, row := range rows {
		c := toCurrencyRateDomain(row)
		rates = append(rates, &c)
	}
	return rates, nil
}

func (r *CurrencyRateRepository) ListByCurrencyID(ctx context.Context, currencyID uuid.UUID) ([]*pricing.CurrencyRate, error) {
	rows, err := r.q.ListCurrencyRatesByCurrencyID(ctx, currencyID)
	if err != nil {
		return nil, err
	}
	rates := make([]*pricing.CurrencyRate, 0, len(rows))
	for _, row := range rows {
		c := toCurrencyRateDomain(row)
		rates = append(rates, &c)
	}
	return rates, nil
}

func (r *CurrencyRateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCurrencyRate(ctx, id)
}

func toCurrencyRateDomain(row sqlc.CurrencyRate) pricing.CurrencyRate {
	return pricing.CurrencyRate{
		ID:          row.ID,
		CurrencyID:  row.CurrencyID,
		RateToJPY:   row.RateToJpy,
		EffectiveAt: row.EffectiveAt,
	}
}
