package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// PromoCodeRepository は coupon.PromoCodeRepository の実装。
type PromoCodeRepository struct {
	q *sqlc.Queries
}

func NewPromoCodeRepository(db *sql.DB) *PromoCodeRepository {
	return &PromoCodeRepository{q: sqlc.New(db)}
}

var _ coupon.PromoCodeRepository = (*PromoCodeRepository)(nil)

func (r *PromoCodeRepository) Create(ctx context.Context, p *coupon.PromoCode) error {
	row, err := r.q.CreatePromoCode(ctx, sqlc.CreatePromoCodeParams{
		ID:         p.ID,
		Code:       p.Code,
		CampaignID: p.CampaignID,
	})
	if err != nil {
		return err
	}
	*p = toPromoCodeDomain(row)
	return nil
}

func (r *PromoCodeRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.PromoCode, error) {
	row, err := r.q.GetPromoCode(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPromoCodeDomain(row)
	return &p, nil
}

func (r *PromoCodeRepository) List(ctx context.Context) ([]*coupon.PromoCode, error) {
	rows, err := r.q.ListPromoCodes(ctx)
	if err != nil {
		return nil, err
	}
	codes := make([]*coupon.PromoCode, 0, len(rows))
	for _, row := range rows {
		p := toPromoCodeDomain(row)
		codes = append(codes, &p)
	}
	return codes, nil
}

func (r *PromoCodeRepository) ListByCampaignID(ctx context.Context, campaignID uuid.UUID) ([]*coupon.PromoCode, error) {
	rows, err := r.q.ListPromoCodesByCampaignID(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	codes := make([]*coupon.PromoCode, 0, len(rows))
	for _, row := range rows {
		p := toPromoCodeDomain(row)
		codes = append(codes, &p)
	}
	return codes, nil
}

func (r *PromoCodeRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePromoCode(ctx, id)
}

func toPromoCodeDomain(row sqlc.PromoCode) coupon.PromoCode {
	return coupon.PromoCode{
		ID:         row.ID,
		Code:       row.Code,
		CampaignID: row.CampaignID,
	}
}
