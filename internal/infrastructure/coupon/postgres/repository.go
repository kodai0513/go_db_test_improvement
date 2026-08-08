package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// Repository implements coupon.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ coupon.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, c *coupon.Coupon) error {
	row, err := r.q.CreateCoupon(ctx, sqlc.CreateCouponParams{
		ID:           c.ID,
		Code:         c.Code,
		DiscountRate: c.DiscountRate,
		ExpiresAt:    c.ExpiresAt,
	})
	if err != nil {
		return err
	}
	*c = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*coupon.Coupon, error) {
	row, err := r.q.GetCoupon(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toDomain(row)
	return &c, nil
}

func (r *Repository) List(ctx context.Context) ([]*coupon.Coupon, error) {
	rows, err := r.q.ListCoupons(ctx)
	if err != nil {
		return nil, err
	}
	coupons := make([]*coupon.Coupon, 0, len(rows))
	for _, row := range rows {
		c := toDomain(row)
		coupons = append(coupons, &c)
	}
	return coupons, nil
}

func toDomain(row sqlc.Coupon) coupon.Coupon {
	return coupon.Coupon{
		ID:           row.ID,
		Code:         row.Code,
		DiscountRate: row.DiscountRate,
		ExpiresAt:    row.ExpiresAt,
	}
}
