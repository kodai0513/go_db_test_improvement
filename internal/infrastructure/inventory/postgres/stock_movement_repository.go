package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// StockMovementRepository は inventory.StockMovementRepository の実装。
type StockMovementRepository struct {
	q *sqlc.Queries
}

func NewStockMovementRepository(db *sql.DB) *StockMovementRepository {
	return &StockMovementRepository{q: sqlc.New(db)}
}

var _ inventory.StockMovementRepository = (*StockMovementRepository)(nil)

func (r *StockMovementRepository) Create(ctx context.Context, m *inventory.StockMovement) error {
	row, err := r.q.CreateStockMovement(ctx, sqlc.CreateStockMovementParams{
		ID:            m.ID,
		InventoryID:   m.InventoryID,
		QuantityDelta: m.QuantityDelta,
		Reason:        m.Reason,
		CreatedAt:     m.CreatedAt,
	})
	if err != nil {
		return err
	}
	*m = toStockMovementDomain(row)
	return nil
}

func (r *StockMovementRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.StockMovement, error) {
	row, err := r.q.GetStockMovement(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toStockMovementDomain(row)
	return &m, nil
}

func (r *StockMovementRepository) List(ctx context.Context) ([]*inventory.StockMovement, error) {
	rows, err := r.q.ListStockMovements(ctx)
	if err != nil {
		return nil, err
	}
	movements := make([]*inventory.StockMovement, 0, len(rows))
	for _, row := range rows {
		m := toStockMovementDomain(row)
		movements = append(movements, &m)
	}
	return movements, nil
}

func (r *StockMovementRepository) ListByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]*inventory.StockMovement, error) {
	rows, err := r.q.ListStockMovementsByInventoryID(ctx, inventoryID)
	if err != nil {
		return nil, err
	}
	movements := make([]*inventory.StockMovement, 0, len(rows))
	for _, row := range rows {
		m := toStockMovementDomain(row)
		movements = append(movements, &m)
	}
	return movements, nil
}

func (r *StockMovementRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStockMovement(ctx, id)
}

func toStockMovementDomain(row sqlc.StockMovement) inventory.StockMovement {
	return inventory.StockMovement{
		ID:            row.ID,
		InventoryID:   row.InventoryID,
		QuantityDelta: row.QuantityDelta,
		Reason:        row.Reason,
		CreatedAt:     row.CreatedAt,
	}
}
