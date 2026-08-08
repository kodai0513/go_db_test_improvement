package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// PriceRuleRepository は pricing.PriceRuleRepository の実装。
type PriceRuleRepository struct {
	q *sqlc.Queries
}

func NewPriceRuleRepository(db *sql.DB) *PriceRuleRepository {
	return &PriceRuleRepository{q: sqlc.New(db)}
}

var _ pricing.PriceRuleRepository = (*PriceRuleRepository)(nil)

func (r *PriceRuleRepository) Create(ctx context.Context, p *pricing.PriceRule) error {
	row, err := r.q.CreatePriceRule(ctx, sqlc.CreatePriceRuleParams{
		ID:          p.ID,
		PriceListID: p.PriceListID,
		ProductID:   p.ProductID,
		PriceYen:    p.PriceYen,
	})
	if err != nil {
		return err
	}
	*p = toPriceRuleDomain(row)
	return nil
}

func (r *PriceRuleRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.PriceRule, error) {
	row, err := r.q.GetPriceRule(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPriceRuleDomain(row)
	return &p, nil
}

func (r *PriceRuleRepository) List(ctx context.Context) ([]*pricing.PriceRule, error) {
	rows, err := r.q.ListPriceRules(ctx)
	if err != nil {
		return nil, err
	}
	rules := make([]*pricing.PriceRule, 0, len(rows))
	for _, row := range rows {
		p := toPriceRuleDomain(row)
		rules = append(rules, &p)
	}
	return rules, nil
}

func (r *PriceRuleRepository) ListByPriceListID(ctx context.Context, priceListID uuid.UUID) ([]*pricing.PriceRule, error) {
	rows, err := r.q.ListPriceRulesByPriceListID(ctx, priceListID)
	if err != nil {
		return nil, err
	}
	rules := make([]*pricing.PriceRule, 0, len(rows))
	for _, row := range rows {
		p := toPriceRuleDomain(row)
		rules = append(rules, &p)
	}
	return rules, nil
}

func (r *PriceRuleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePriceRule(ctx, id)
}

func toPriceRuleDomain(row sqlc.PriceRule) pricing.PriceRule {
	return pricing.PriceRule{
		ID:          row.ID,
		PriceListID: row.PriceListID,
		ProductID:   row.ProductID,
		PriceYen:    row.PriceYen,
	}
}
