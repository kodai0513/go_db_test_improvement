package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres/sqlc"
)

// CampaignTargetRepository は coupon.CampaignTargetRepository の実装。
type CampaignTargetRepository struct {
	q *sqlc.Queries
}

func NewCampaignTargetRepository(db *sql.DB) *CampaignTargetRepository {
	return &CampaignTargetRepository{q: sqlc.New(db)}
}

var _ coupon.CampaignTargetRepository = (*CampaignTargetRepository)(nil)

func (r *CampaignTargetRepository) Create(ctx context.Context, t *coupon.CampaignTarget) error {
	row, err := r.q.CreateCampaignTarget(ctx, sqlc.CreateCampaignTargetParams{
		ID:         t.ID,
		CampaignID: t.CampaignID,
		MemberID:   t.MemberID,
	})
	if err != nil {
		return err
	}
	*t = toCampaignTargetDomain(row)
	return nil
}

func (r *CampaignTargetRepository) Get(ctx context.Context, id uuid.UUID) (*coupon.CampaignTarget, error) {
	row, err := r.q.GetCampaignTarget(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toCampaignTargetDomain(row)
	return &t, nil
}

func (r *CampaignTargetRepository) List(ctx context.Context) ([]*coupon.CampaignTarget, error) {
	rows, err := r.q.ListCampaignTargets(ctx)
	if err != nil {
		return nil, err
	}
	targets := make([]*coupon.CampaignTarget, 0, len(rows))
	for _, row := range rows {
		t := toCampaignTargetDomain(row)
		targets = append(targets, &t)
	}
	return targets, nil
}

func (r *CampaignTargetRepository) ListByCampaignID(ctx context.Context, campaignID uuid.UUID) ([]*coupon.CampaignTarget, error) {
	rows, err := r.q.ListCampaignTargetsByCampaignID(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	targets := make([]*coupon.CampaignTarget, 0, len(rows))
	for _, row := range rows {
		t := toCampaignTargetDomain(row)
		targets = append(targets, &t)
	}
	return targets, nil
}

func (r *CampaignTargetRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCampaignTarget(ctx, id)
}

func toCampaignTargetDomain(row sqlc.CampaignTarget) coupon.CampaignTarget {
	return coupon.CampaignTarget{
		ID:         row.ID,
		CampaignID: row.CampaignID,
		MemberID:   row.MemberID,
	}
}
