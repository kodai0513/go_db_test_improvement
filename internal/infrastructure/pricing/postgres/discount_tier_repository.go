package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// DiscountTierRepository は pricing.DiscountTierRepository の実装。
type DiscountTierRepository struct {
	q *sqlc.Queries
}

func NewDiscountTierRepository(db *sql.DB) *DiscountTierRepository {
	return &DiscountTierRepository{q: sqlc.New(db)}
}

var _ pricing.DiscountTierRepository = (*DiscountTierRepository)(nil)

func (r *DiscountTierRepository) Create(ctx context.Context, d *pricing.DiscountTier) error {
	row, err := r.q.CreateDiscountTier(ctx, sqlc.CreateDiscountTierParams{
		ID:              d.ID,
		MinQuantity:     d.MinQuantity,
		DiscountPercent: d.DiscountPercent,
	})
	if err != nil {
		return err
	}
	*d = toDiscountTierDomain(row)
	return nil
}

func (r *DiscountTierRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.DiscountTier, error) {
	row, err := r.q.GetDiscountTier(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toDiscountTierDomain(row)
	return &d, nil
}

func (r *DiscountTierRepository) List(ctx context.Context) ([]*pricing.DiscountTier, error) {
	rows, err := r.q.ListDiscountTiers(ctx)
	if err != nil {
		return nil, err
	}
	tiers := make([]*pricing.DiscountTier, 0, len(rows))
	for _, row := range rows {
		d := toDiscountTierDomain(row)
		tiers = append(tiers, &d)
	}
	return tiers, nil
}

func (r *DiscountTierRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteDiscountTier(ctx, id)
}

func toDiscountTierDomain(row sqlc.DiscountTier) pricing.DiscountTier {
	return pricing.DiscountTier{
		ID:              row.ID,
		MinQuantity:     row.MinQuantity,
		DiscountPercent: row.DiscountPercent,
	}
}
