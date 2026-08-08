package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// PointRuleRepository は loyalty.PointRuleRepository の実装。
type PointRuleRepository struct {
	q *sqlc.Queries
}

func NewPointRuleRepository(db *sql.DB) *PointRuleRepository {
	return &PointRuleRepository{q: sqlc.New(db)}
}

var _ loyalty.PointRuleRepository = (*PointRuleRepository)(nil)

func (r *PointRuleRepository) Create(ctx context.Context, p *loyalty.PointRule) error {
	row, err := r.q.CreatePointRule(ctx, sqlc.CreatePointRuleParams{
		ID:               p.ID,
		LoyaltyProgramID: p.LoyaltyProgramID,
		Action:           p.Action,
		Points:           p.Points,
	})
	if err != nil {
		return err
	}
	*p = toPointRuleDomain(row)
	return nil
}

func (r *PointRuleRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.PointRule, error) {
	row, err := r.q.GetPointRule(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPointRuleDomain(row)
	return &p, nil
}

func (r *PointRuleRepository) List(ctx context.Context) ([]*loyalty.PointRule, error) {
	rows, err := r.q.ListPointRules(ctx)
	if err != nil {
		return nil, err
	}
	rules := make([]*loyalty.PointRule, 0, len(rows))
	for _, row := range rows {
		p := toPointRuleDomain(row)
		rules = append(rules, &p)
	}
	return rules, nil
}

func (r *PointRuleRepository) ListByLoyaltyProgramID(ctx context.Context, loyaltyProgramID uuid.UUID) ([]*loyalty.PointRule, error) {
	rows, err := r.q.ListPointRulesByLoyaltyProgramID(ctx, loyaltyProgramID)
	if err != nil {
		return nil, err
	}
	rules := make([]*loyalty.PointRule, 0, len(rows))
	for _, row := range rows {
		p := toPointRuleDomain(row)
		rules = append(rules, &p)
	}
	return rules, nil
}

func (r *PointRuleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePointRule(ctx, id)
}

func toPointRuleDomain(row sqlc.PointRule) loyalty.PointRule {
	return loyalty.PointRule{
		ID:               row.ID,
		LoyaltyProgramID: row.LoyaltyProgramID,
		Action:           row.Action,
		Points:           row.Points,
	}
}
