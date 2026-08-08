package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CouponUsageRepository は coupon.CouponUsageRepository の実装。
type CouponUsageRepository struct {
	q *sqlc.Queries
}

func NewCouponUsageRepository(db *sql.DB) *CouponUsageRepository {
	return &CouponUsageRepository{q: sqlc.New(db)}
}

var _ coupon.CouponUsageRepository = (*CouponUsageRepository)(nil)

func (r *CouponUsageRepository) Create(ctx context.Context, u *coupon.CouponUsage) error {
	row, err := r.q.CreateCouponUsage(ctx, sqlc.CreateCouponUsageParams{
		ID:       u.ID,
		CouponID: u.CouponID,
		OrderID:  u.OrderID,
		UsedAt:   u.UsedAt,
	})
	if err != nil {
		return err
	}
	*u = toCouponUsageDomain(row)
	return nil
}

func (r *CouponUsageRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.CouponUsage, error) {
	row, err := r.q.GetCouponUsage(ctx, id)
	if err != nil {
		return nil, err
	}
	u := toCouponUsageDomain(row)
	return &u, nil
}

func (r *CouponUsageRepository) List(ctx context.Context) ([]*coupon.CouponUsage, error) {
	rows, err := r.q.ListCouponUsages(ctx)
	if err != nil {
		return nil, err
	}
	usages := make([]*coupon.CouponUsage, 0, len(rows))
	for _, row := range rows {
		u := toCouponUsageDomain(row)
		usages = append(usages, &u)
	}
	return usages, nil
}

func (r *CouponUsageRepository) ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*coupon.CouponUsage, error) {
	rows, err := r.q.ListCouponUsagesByCouponID(ctx, couponID)
	if err != nil {
		return nil, err
	}
	usages := make([]*coupon.CouponUsage, 0, len(rows))
	for _, row := range rows {
		u := toCouponUsageDomain(row)
		usages = append(usages, &u)
	}
	return usages, nil
}

func (r *CouponUsageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCouponUsage(ctx, id)
}

func toCouponUsageDomain(row sqlc.CouponUsage) coupon.CouponUsage {
	return coupon.CouponUsage{
		ID:       row.ID,
		CouponID: row.CouponID,
		OrderID:  row.OrderID,
		UsedAt:   row.UsedAt,
	}
}
