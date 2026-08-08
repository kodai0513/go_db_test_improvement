package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// StockAdjustmentRepository は inventory.StockAdjustmentRepository の実装。
type StockAdjustmentRepository struct {
	q *sqlc.Queries
}

func NewStockAdjustmentRepository(db *sql.DB) *StockAdjustmentRepository {
	return &StockAdjustmentRepository{q: sqlc.New(db)}
}

var _ inventory.StockAdjustmentRepository = (*StockAdjustmentRepository)(nil)

func (r *StockAdjustmentRepository) Create(ctx context.Context, a *inventory.StockAdjustment) error {
	row, err := r.q.CreateStockAdjustment(ctx, sqlc.CreateStockAdjustmentParams{
		ID:               a.ID,
		InventoryID:      a.InventoryID,
		AdjustedQuantity: a.AdjustedQuantity,
		Note:             a.Note,
		CreatedAt:        a.CreatedAt,
	})
	if err != nil {
		return err
	}
	*a = toStockAdjustmentDomain(row)
	return nil
}

func (r *StockAdjustmentRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.StockAdjustment, error) {
	row, err := r.q.GetStockAdjustment(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toStockAdjustmentDomain(row)
	return &a, nil
}

func (r *StockAdjustmentRepository) List(ctx context.Context) ([]*inventory.StockAdjustment, error) {
	rows, err := r.q.ListStockAdjustments(ctx)
	if err != nil {
		return nil, err
	}
	adjustments := make([]*inventory.StockAdjustment, 0, len(rows))
	for _, row := range rows {
		a := toStockAdjustmentDomain(row)
		adjustments = append(adjustments, &a)
	}
	return adjustments, nil
}

func (r *StockAdjustmentRepository) ListByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]*inventory.StockAdjustment, error) {
	rows, err := r.q.ListStockAdjustmentsByInventoryID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	adjustments := make([]*inventory.StockAdjustment, 0, len(rows))
	for _, row := range rows {
		a := toStockAdjustmentDomain(row)
		adjustments = append(adjustments, &a)
	}
	return adjustments, nil
}

func (r *StockAdjustmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStockAdjustment(ctx, id)
}

func toStockAdjustmentDomain(row sqlc.StockAdjustment) inventory.StockAdjustment {
	return inventory.StockAdjustment{
		ID:               row.ID,
		InventoryID:      row.InventoryID,
		AdjustedQuantity: row.AdjustedQuantity,
		Note:             row.Note,
		CreatedAt:        row.CreatedAt,
	}
}
