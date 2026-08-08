package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// TaxRateRepository は pricing.TaxRateRepository の実装。
type TaxRateRepository struct {
	q *sqlc.Queries
}

func NewTaxRateRepository(db *sql.DB) *TaxRateRepository {
	return &TaxRateRepository{q: sqlc.New(db)}
}

var _ pricing.TaxRateRepository = (*TaxRateRepository)(nil)

func (r *TaxRateRepository) Create(ctx context.Context, t *pricing.TaxRate) error {
	row, err := r.q.CreateTaxRate(ctx, sqlc.CreateTaxRateParams{
		ID:   t.ID,
		Name: t.Name,
		Rate: t.Rate,
	})
	if err != nil {
		return err
	}
	*t = toTaxRateDomain(row)
	return nil
}

func (r *TaxRateRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.TaxRate, error) {
	row, err := r.q.GetTaxRate(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toTaxRateDomain(row)
	return &t, nil
}

func (r *TaxRateRepository) List(ctx context.Context) ([]*pricing.TaxRate, error) {
	rows, err := r.q.ListTaxRates(ctx)
	if err != nil {
		return nil, err
	}
	rates := make([]*pricing.TaxRate, 0, len(rows))
	for _, row := range rows {
		t := toTaxRateDomain(row)
		rates = append(rates, &t)
	}
	return rates, nil
}

func (r *TaxRateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTaxRate(ctx, id)
}

func toTaxRateDomain(row sqlc.TaxRate) pricing.TaxRate {
	return pricing.TaxRate{
		ID:   row.ID,
		Name: row.Name,
		Rate: row.Rate,
	}
}
