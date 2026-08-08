package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CouponProductRepository は coupon.CouponProductRepository の実装。
type CouponProductRepository struct {
	q *sqlc.Queries
}

func NewCouponProductRepository(db *sql.DB) *CouponProductRepository {
	return &CouponProductRepository{q: sqlc.New(db)}
}

var _ coupon.CouponProductRepository = (*CouponProductRepository)(nil)

func (r *CouponProductRepository) Create(ctx context.Context, p *coupon.CouponProduct) error {
	row, err := r.q.CreateCouponProduct(ctx, sqlc.CreateCouponProductParams{
		ID:        p.ID,
		CouponID:  p.CouponID,
		ProductID: p.ProductID,
	})
	if err != nil {
		return err
	}
	*p = toCouponProductDomain(row)
	return nil
}

func (r *CouponProductRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.CouponProduct, error) {
	row, err := r.q.GetCouponProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toCouponProductDomain(row)
	return &p, nil
}

func (r *CouponProductRepository) List(ctx context.Context) ([]*coupon.CouponProduct, error) {
	rows, err := r.q.ListCouponProducts(ctx)
	if err != nil {
		return nil, err
	}
	products := make([]*coupon.CouponProduct, 0, len(rows))
	for _, row := range rows {
		p := toCouponProductDomain(row)
		products = append(products, &p)
	}
	return products, nil
}

func (r *CouponProductRepository) ListByCouponID(ctx context.Context, couponID uuid.UUID) ([]*coupon.CouponProduct, error) {
	rows, err := r.q.ListCouponProductsByCouponID(ctx, couponID)
	if err != nil {
		return nil, err
	}
	products := make([]*coupon.CouponProduct, 0, len(rows))
	for _, row := range rows {
		p := toCouponProductDomain(row)
		products = append(products, &p)
	}
	return products, nil
}

func (r *CouponProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCouponProduct(ctx, id)
}

func toCouponProductDomain(row sqlc.CouponProduct) coupon.CouponProduct {
	return coupon.CouponProduct{
		ID:        row.ID,
		CouponID:  row.CouponID,
		ProductID: row.ProductID,
	}
}
