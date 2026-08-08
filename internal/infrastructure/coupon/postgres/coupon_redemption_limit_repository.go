package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CouponRedemptionLimitRepository は coupon.CouponRedemptionLimitRepository の実装。
type CouponRedemptionLimitRepository struct {
	q *sqlc.Queries
}

func NewCouponRedemptionLimitRepository(db *sql.DB) *CouponRedemptionLimitRepository {
	return &CouponRedemptionLimitRepository{q: sqlc.New(db)}
}

var _ coupon.CouponRedemptionLimitRepository = (*CouponRedemptionLimitRepository)(nil)

func (r *CouponRedemptionLimitRepository) Create(ctx context.Context, l *coupon.CouponRedemptionLimit) error {
	row, err := r.q.CreateCouponRedemptionLimit(ctx, sqlc.CreateCouponRedemptionLimitParams{
		ID:             l.ID,
		CouponID:       l.CouponID,
		MaxRedemptions: l.MaxRedemptions,
	})
	if err != nil {
		return err
	}
	*l = toCouponRedemptionLimitDomain(row)
	return nil
}

func (r *CouponRedemptionLimitRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.CouponRedemptionLimit, error) {
	row, err := r.q.GetCouponRedemptionLimit(ctx, id)
	if err != nil {
		return nil, err
	}
	l := toCouponRedemptionLimitDomain(row)
	return &l, nil
}

func (r *CouponRedemptionLimitRepository) List(ctx context.Context) ([]*coupon.CouponRedemptionLimit, error) {
	rows, err := r.q.ListCouponRedemptionLimits(ctx)
	if err != nil {
		return nil, err
	}
	limits := make([]*coupon.CouponRedemptionLimit, 0, len(rows))
	for _, row := range rows {
		l := toCouponRedemptionLimitDomain(row)
		limits = append(limits, &l)
	}
	return limits, nil
}

func (r *CouponRedemptionLimitRepository) ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*coupon.CouponRedemptionLimit, error) {
	rows, err := r.q.ListCouponRedemptionLimitsByCouponID(ctx, couponID)
	if err != nil {
		return nil, err
	}
	limits := make([]*coupon.CouponRedemptionLimit, 0, len(rows))
	for _, row := range rows {
		l := toCouponRedemptionLimitDomain(row)
		limits = append(limits, &l)
	}
	return limits, nil
}

func (r *CouponRedemptionLimitRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCouponRedemptionLimit(ctx, id)
}

func toCouponRedemptionLimitDomain(row sqlc.CouponRedemptionLimit) coupon.CouponRedemptionLimit {
	return coupon.CouponRedemptionLimit{
		ID:             row.ID,
		CouponID:       row.CouponID,
		MaxRedemptions: row.MaxRedemptions,
	}
}
