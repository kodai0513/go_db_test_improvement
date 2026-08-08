package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// LoyaltyTierRepository は loyalty.LoyaltyTierRepository の実装。
type LoyaltyTierRepository struct {
	q *sqlc.Queries
}

func NewLoyaltyTierRepository(db *sql.DB) *LoyaltyTierRepository {
	return &LoyaltyTierRepository{q: sqlc.New(db)}
}

var _ loyalty.LoyaltyTierRepository = (*LoyaltyTierRepository)(nil)

func (r *LoyaltyTierRepository) Create(ctx context.Context, t *loyalty.LoyaltyTier) error {
	row, err := r.q.CreateLoyaltyTier(ctx, sqlc.CreateLoyaltyTierParams{
		ID:               t.ID,
		LoyaltyProgramID: t.LoyaltyProgramID,
		Name:             t.Name,
		MinPoints:        t.MinPoints,
	})
	if err != nil {
		return err
	}
	*t = toLoyaltyTierDomain(row)
	return nil
}

func (r *LoyaltyTierRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.LoyaltyTier, error) {
	row, err := r.q.GetLoyaltyTier(ctx, id)
	if err != nil {
		return nil, err
	}
	t := toLoyaltyTierDomain(row)
	return &t, nil
}

func (r *LoyaltyTierRepository) List(ctx context.Context) ([]*loyalty.LoyaltyTier, error) {
	rows, err := r.q.ListLoyaltyTiers(ctx)
	if err != nil {
		return nil, err
	}
	tiers := make([]*loyalty.LoyaltyTier, 0, len(rows))
	for _, row := range rows {
		t := toLoyaltyTierDomain(row)
		tiers = append(tiers, &t)
	}
	return tiers, nil
}

func (r *LoyaltyTierRepository) ListByLoyaltyProgramID(ctx context.Context, loyaltyProgramID uuid.UUID) ([]*loyalty.LoyaltyTier, error) {
	rows, err := r.q.ListLoyaltyTiersByLoyaltyProgramID(ctx, loyaltyProgramID)
	if err != nil {
		return nil, err
	}
	tiers := make([]*loyalty.LoyaltyTier, 0, len(rows))
	for _, row := range rows {
		t := toLoyaltyTierDomain(row)
		tiers = append(tiers, &t)
	}
	return tiers, nil
}

func (r *LoyaltyTierRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteLoyaltyTier(ctx, id)
}

func toLoyaltyTierDomain(row sqlc.LoyaltyTier) loyalty.LoyaltyTier {
	return loyalty.LoyaltyTier{
		ID:               row.ID,
		LoyaltyProgramID: row.LoyaltyProgramID,
		Name:             row.Name,
		MinPoints:        row.MinPoints,
	}
}
