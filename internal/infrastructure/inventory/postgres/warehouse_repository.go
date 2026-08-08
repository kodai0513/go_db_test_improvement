package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// WarehouseRepository は inventory.WarehouseRepository の実装。
type WarehouseRepository struct {
	q *sqlc.Queries
}

func NewWarehouseRepository(db *sql.DB) *WarehouseRepository {
	return &WarehouseRepository{q: sqlc.New(db)}
}

var _ inventory.WarehouseRepository = (*WarehouseRepository)(nil)

func (r *WarehouseRepository) Create(ctx context.Context, w *inventory.Warehouse) error {
	row, err := r.q.CreateWarehouse(ctx, sqlc.CreateWarehouseParams{
		ID:        w.ID,
		Name:      w.Name,
		Address:   w.Address,
		CreatedAt: w.CreatedAt,
	})
	if err != nil {
		return err
	}
	*w = toWarehouseDomain(row)
	return nil
}

func (r *WarehouseRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.Warehouse, error) {
	row, err := r.q.GetWarehouse(ctx, id)
	if err != nil {
		return nil, err
	}
	w := toWarehouseDomain(row)
	return &w, nil
}

func (r *WarehouseRepository) List(ctx context.Context) ([]*inventory.Warehouse, error) {
	rows, err := r.q.ListWarehouses(ctx)
	if err != nil {
		return nil, err
	}
	warehouses := make([]*inventory.Warehouse, 0, len(rows))
	for _, row := range rows {
		w := toWarehouseDomain(row)
		warehouses = append(warehouses, &w)
	}
	return warehouses, nil
}

func (r *WarehouseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteWarehouse(ctx, id)
}

func toWarehouseDomain(row sqlc.Warehouse) inventory.Warehouse {
	return inventory.Warehouse{
		ID:        row.ID,
		Name:      row.Name,
		Address:   row.Address,
		CreatedAt: row.CreatedAt,
	}
}
