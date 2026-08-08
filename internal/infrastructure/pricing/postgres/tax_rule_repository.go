package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// TaxRuleRepository は pricing.TaxRuleRepository の実装。
type TaxRuleRepository struct {
	q *sqlc.Queries
}

func NewTaxRuleRepository(db *sql.DB) *TaxRuleRepository {
	return &TaxRuleRepository{q: sqlc.New(db)}
}

var _ pricing.TaxRuleRepository = (*TaxRuleRepository)(nil)

func (r *TaxRuleRepository) Create(ctx context.Context, t *pricing.TaxRule) error {
	row, err := r.q.CreateTaxRule(ctx, sqlc.CreateTaxRuleParams{
		ID:        t.ID,
		TaxRateID: t.TaxRateID,
		Region:    t.Region,
	})
	if err != nil {
		return err
	}
	*t = toTaxRuleDomain(row)
	return nil
}

func (r *TaxRuleRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.TaxRule, error) {
	row, err := r.q.GetTaxRule(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toTaxRuleDomain(row)
	return &t, nil
}

func (r *TaxRuleRepository) List(ctx context.Context) ([]*pricing.TaxRule, error) {
	rows, err := r.q.ListTaxRules(ctx)
	if err != nil {
		return nil, err
	}
	rules := make([]*pricing.TaxRule, 0, len(rows))
	for _, row := range rows {
		t := toTaxRuleDomain(row)
		rules = append(rules, &t)
	}
	return rules, nil
}

func (r *TaxRuleRepository) ListByTaxRateID(ctx context.Context, taxRateID uuid.UUID) ([]*pricing.TaxRule, error) {
	rows, err := r.q.ListTaxRulesByTaxRateID(ctx, taxRateID)
	if err != nil {
		return nil, err
	}
	rules := make([]*pricing.TaxRule, 0, len(rows))
	for _, row := range rows {
		t := toTaxRuleDomain(row)
		rules = append(rules, &t)
	}
	return rules, nil
}

func (r *TaxRuleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTaxRule(ctx, id)
}

func toTaxRuleDomain(row sqlc.TaxRule) pricing.TaxRule {
	return pricing.TaxRule{
		ID:        row.ID,
		TaxRateID: row.TaxRateID,
		Region:    row.Region,
	}
}
