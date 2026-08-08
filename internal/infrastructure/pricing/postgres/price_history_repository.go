package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres/sqlc"
)

// PriceHistoryRepository は pricing.PriceHistoryRepository の実装。
type PriceHistoryRepository struct {
	q *sqlc.Queries
}

func NewPriceHistoryRepository(db *sql.DB) *PriceHistoryRepository {
	return &PriceHistoryRepository{q: sqlc.New(db)}
}

var _ pricing.PriceHistoryRepository = (*PriceHistoryRepository)(nil)

func (r *PriceHistoryRepository) Create(ctx context.Context, p *pricing.PriceHistory) error {
	row, err := r.q.CreatePriceHistory(ctx, sqlc.CreatePriceHistoryParams{
		ID:        p.ID,
		ProductID: p.ProductID,
		PriceYen:  p.PriceYen,
		ChangedAt: p.ChangedAt,
	})
	if err != nil {
		return err
	}
	*p = toPriceHistoryDomain(row)
	return nil
}

func (r *PriceHistoryRepository) Get(ctx context.Context, id uuid.UUID) (*pricing.PriceHistory, error) {
	row, err := r.q.GetPriceHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPriceHistoryDomain(row)
	return &p, nil
}

func (r *PriceHistoryRepository) List(ctx context.Context) ([]*pricing.PriceHistory, error) {
	rows, err := r.q.ListPriceHistories(ctx)
	if err != nil {
		return nil, err
	}
	histories := make([]*pricing.PriceHistory, 0, len(rows))
	for _, row := range rows {
		p := toPriceHistoryDomain(row)
		histories = append(histories, &p)
	}
	return histories, nil
}

func (r *PriceHistoryRepository) ListByProductID(ctx context.Context, productID uuid.UUID) ([]*pricing.PriceHistory, error) {
	rows, err := r.q.ListPriceHistoriesByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	histories := make([]*pricing.PriceHistory, 0, len(rows))
	for _, row := range rows {
		p := toPriceHistoryDomain(row)
		histories = append(histories, &p)
	}
	return histories, nil
}

func (r *PriceHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePriceHistory(ctx, id)
}

func toPriceHistoryDomain(row sqlc.PriceHistory) pricing.PriceHistory {
	return pricing.PriceHistory{
		ID:        row.ID,
		ProductID: row.ProductID,
		PriceYen:  row.PriceYen,
		ChangedAt: row.ChangedAt,
	}
}
