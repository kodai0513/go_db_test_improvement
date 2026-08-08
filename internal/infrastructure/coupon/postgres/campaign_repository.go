package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CampaignRepository は coupon.CampaignRepository の実装。
type CampaignRepository struct {
	q *sqlc.Queries
}

func NewCampaignRepository(db *sql.DB) *CampaignRepository {
	return &CampaignRepository{q: sqlc.New(db)}
}

var _ coupon.CampaignRepository = (*CampaignRepository)(nil)

func (r *CampaignRepository) Create(ctx context.Context, c *coupon.Campaign) error {
	row, err := r.q.CreateCampaign(ctx, sqlc.CreateCampaignParams{
		ID:       c.ID,
		Name:     c.Name,
		StartsAt: c.StartsAt,
		EndsAt:   c.EndsAt,
	})
	if err != nil {
		return err
	}
	*c = toCampaignDomain(row)
	return nil
}

func (r *CampaignRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.Campaign, error) {
	row, err := r.q.GetCampaign(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCampaignDomain(row)
	return &c, nil
}

func (r *CampaignRepository) List(ctx context.Context) ([]*coupon.Campaign, error) {
	rows, err := r.q.ListCampaigns(ctx)
	if err != nil {
		return nil, err
	}
	campaigns := make([]*coupon.Campaign, 0, len(rows))
	for _, row := range rows {
		c := toCampaignDomain(row)
		campaigns = append(campaigns, &c)
	}
	return campaigns, nil
}

func (r *CampaignRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCampaign(ctx, id)
}

func toCampaignDomain(row sqlc.Campaign) coupon.Campaign {
	return coupon.Campaign{
		ID:       row.ID,
		Name:     row.Name,
		StartsAt: row.StartsAt,
		EndsAt:   row.EndsAt,
	}
}
