package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// BundleOfferRepository は pricing.BundleOfferRepository の実装。
type BundleOfferRepository struct {
	q *sqlc.Queries
}

func NewBundleOfferRepository(db *sql.DB) *BundleOfferRepository {
	return &BundleOfferRepository{q: sqlc.New(db)}
}

var _ pricing.BundleOfferRepository = (*BundleOfferRepository)(nil)

func (r *BundleOfferRepository) Create(ctx context.Context, b *pricing.BundleOffer) error {
	row, err := r.q.CreateBundleOffer(ctx, sqlc.CreateBundleOfferParams{
		ID:       b.ID,
		Name:     b.Name,
		PriceYen: b.PriceYen,
	})
	if err != nil {
		return err
	}
	*b = toBundleOfferDomain(row)
	return nil
}

func (r *BundleOfferRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.BundleOffer, error) {
	row, err := r.q.GetBundleOffer(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toBundleOfferDomain(row)
	return &b, nil
}

func (r *BundleOfferRepository) List(ctx context.Context) ([]*pricing.BundleOffer, error) {
	rows, err := r.q.ListBundleOffers(ctx)
	if err != nil {
		return nil, err
	}
	offers := make([]*pricing.BundleOffer, 0, len(rows))
	for _, row := range rows {
		b := toBundleOfferDomain(row)
		offers = append(offers, &b)
	}
	return offers, nil
}

func (r *BundleOfferRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBundleOffer(ctx, id)
}

func toBundleOfferDomain(row sqlc.BundleOffer) pricing.BundleOffer {
	return pricing.BundleOffer{
		ID:       row.ID,
		Name:     row.Name,
		PriceYen: row.PriceYen,
	}
}
