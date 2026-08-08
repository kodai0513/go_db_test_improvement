package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// BundleOfferItemRepository は pricing.BundleOfferItemRepository の実装。
type BundleOfferItemRepository struct {
	q *sqlc.Queries
}

func NewBundleOfferItemRepository(db *sql.DB) *BundleOfferItemRepository {
	return &BundleOfferItemRepository{q: sqlc.New(db)}
}

var _ pricing.BundleOfferItemRepository = (*BundleOfferItemRepository)(nil)

func (r *BundleOfferItemRepository) Create(ctx context.Context, b *pricing.BundleOfferItem) error {
	row, err := r.q.CreateBundleOfferItem(ctx, sqlc.CreateBundleOfferItemParams{
		ID:            b.ID,
		BundleOfferID: b.BundleOfferID,
		ProductID:     b.ProductID,
		Quantity:      b.Quantity,
	})
	if err != nil {
		return err
	}
	*b = toBundleOfferItemDomain(row)
	return nil
}

func (r *BundleOfferItemRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.BundleOfferItem, error) {
	row, err := r.q.GetBundleOfferItem(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toBundleOfferItemDomain(row)
	return &b, nil
}

func (r *BundleOfferItemRepository) List(ctx context.Context) ([]*pricing.BundleOfferItem, error) {
	rows, err := r.q.ListBundleOfferItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*pricing.BundleOfferItem, 0, len(rows))
	for _, row := range rows {
		b := toBundleOfferItemDomain(row)
		items = append(items, &b)
	}
	return items, nil
}

func (r *BundleOfferItemRepository) ListByBundleOfferID(ctx context.Context, bundleOfferID uuid.UUID) ([]*pricing.BundleOfferItem, error) {
	rows, err := r.q.ListBundleOfferItemsByBundleOfferID(ctx, bundleOfferID)
	if err != nil {
		return nil, err
	}
	items := make([]*pricing.BundleOfferItem, 0, len(rows))
	for _, row := range rows {
		b := toBundleOfferItemDomain(row)
		items = append(items, &b)
	}
	return items, nil
}

func (r *BundleOfferItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBundleOfferItem(ctx, id)
}

func toBundleOfferItemDomain(row sqlc.BundleOfferItem) pricing.BundleOfferItem {
	return pricing.BundleOfferItem{
		ID:            row.ID,
		BundleOfferID: row.BundleOfferID,
		ProductID:     row.ProductID,
		Quantity:      row.Quantity,
	}
}
