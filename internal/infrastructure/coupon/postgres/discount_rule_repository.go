package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// DiscountRuleRepository は coupon.DiscountRuleRepository の実装。
type DiscountRuleRepository struct {
	q *sqlc.Queries
}

func NewDiscountRuleRepository(db *sql.DB) *DiscountRuleRepository {
	return &DiscountRuleRepository{q: sqlc.New(db)}
}

var _ coupon.DiscountRuleRepository = (*DiscountRuleRepository)(nil)

func (r *DiscountRuleRepository) Create(ctx context.Context, d *coupon.DiscountRule) error {
	row, err := r.q.CreateDiscountRule(ctx, sqlc.CreateDiscountRuleParams{
		ID:           d.ID,
		CouponID:     d.CouponID,
		MinAmountYen: d.MinAmountYen,
	})
	if err != nil {
		return err
	}
	*d = toDiscountRuleDomain(row)
	return nil
}

func (r *DiscountRuleRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.DiscountRule, error) {
	row, err := r.q.GetDiscountRule(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toDiscountRuleDomain(row)
	return &d, nil
}

func (r *DiscountRuleRepository) List(ctx context.Context) ([]*coupon.DiscountRule, error) {
	rows, err := r.q.ListDiscountRules(ctx)
	if err != nil {
		return nil, err
	}
	rules := make([]*coupon.DiscountRule, 0, len(rows))
	for _, row := range rows {
		d := toDiscountRuleDomain(row)
		rules = append(rules, &d)
	}
	return rules, nil
}

func (r *DiscountRuleRepository) ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*coupon.DiscountRule, error) {
	rows, err := r.q.ListDiscountRulesByCouponID(ctx, couponID)
	if err != nil {
		return nil, err
	}
	rules := make([]*coupon.DiscountRule, 0, len(rows))
	for _, row := range rows {
		d := toDiscountRuleDomain(row)
		rules = append(rules, &d)
	}
	return rules, nil
}

func (r *DiscountRuleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteDiscountRule(ctx, id)
}

func toDiscountRuleDomain(row sqlc.DiscountRule) coupon.DiscountRule {
	return coupon.DiscountRule{
		ID:           row.ID,
		CouponID:     row.CouponID,
		MinAmountYen: row.MinAmountYen,
	}
}
