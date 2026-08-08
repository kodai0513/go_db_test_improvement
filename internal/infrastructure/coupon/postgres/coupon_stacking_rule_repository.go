package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CouponStackingRuleRepository は coupon.CouponStackingRuleRepository の実装。
type CouponStackingRuleRepository struct {
	q *sqlc.Queries
}

func NewCouponStackingRuleRepository(db *sql.DB) *CouponStackingRuleRepository {
	return &CouponStackingRuleRepository{q: sqlc.New(db)}
}

var _ coupon.CouponStackingRuleRepository = (*CouponStackingRuleRepository)(nil)

func (r *CouponStackingRuleRepository) Create(ctx context.Context, s *coupon.CouponStackingRule) error {
	row, err := r.q.CreateCouponStackingRule(ctx, sqlc.CreateCouponStackingRuleParams{
		ID:                    s.ID,
		CouponID:              s.CouponID,
		StackableWithCouponID: s.StackableWithCouponID,
	})
	if err != nil {
		return err
	}
	*s = toCouponStackingRuleDomain(row)
	return nil
}

func (r *CouponStackingRuleRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.CouponStackingRule, error) {
	row, err := r.q.GetCouponStackingRule(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toCouponStackingRuleDomain(row)
	return &s, nil
}

func (r *CouponStackingRuleRepository) List(ctx context.Context) ([]*coupon.CouponStackingRule, error) {
	rows, err := r.q.ListCouponStackingRules(ctx)
	if err != nil {
		return nil, err
	}
	rules := make([]*coupon.CouponStackingRule, 0, len(rows))
	for _, row := range rows {
		s := toCouponStackingRuleDomain(row)
		rules = append(rules, &s)
	}
	return rules, nil
}

func (r *CouponStackingRuleRepository) ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*coupon.CouponStackingRule, error) {
	rows, err := r.q.ListCouponStackingRulesByCouponID(ctx, couponID)
	if err != nil {
		return nil, err
	}
	rules := make([]*coupon.CouponStackingRule, 0, len(rows))
	for _, row := range rows {
		s := toCouponStackingRuleDomain(row)
		rules = append(rules, &s)
	}
	return rules, nil
}

func (r *CouponStackingRuleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCouponStackingRule(ctx, id)
}

func toCouponStackingRuleDomain(row sqlc.CouponStackingRule) coupon.CouponStackingRule {
	return coupon.CouponStackingRule{
		ID:                    row.ID,
		CouponID:              row.CouponID,
		StackableWithCouponID: row.StackableWithCouponID,
	}
}
