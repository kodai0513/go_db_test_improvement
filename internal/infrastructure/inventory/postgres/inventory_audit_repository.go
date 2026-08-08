package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// InventoryAuditRepository は inventory.InventoryAuditRepository の実装。
type InventoryAuditRepository struct {
	q *sqlc.Queries
}

func NewInventoryAuditRepository(db *sql.DB) *InventoryAuditRepository {
	return &InventoryAuditRepository{q: sqlc.New(db)}
}

var _ inventory.InventoryAuditRepository = (*InventoryAuditRepository)(nil)

func (r *InventoryAuditRepository) Create(ctx context.Context, a *inventory.InventoryAudit) error {
	row, err := r.q.CreateInventoryAudit(ctx, sqlc.CreateInventoryAuditParams{
		ID:          a.ID,
		WarehouseID: a.WarehouseID,
		PerformedAt: a.PerformedAt,
		Result:      a.Result,
	})
	if err != nil {
		return err
	}
	*a = toInventoryAuditDomain(row)
	return nil
}

func (r *InventoryAuditRepository) Get(ctx context.Context, id uuid.UUID) (*inventory.InventoryAudit, error) {
	row, err := r.q.GetInventoryAudit(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toInventoryAuditDomain(row)
	return &a, nil
}

func (r *InventoryAuditRepository) List(ctx context.Context) ([]*inventory.InventoryAudit, error) {
	rows, err := r.q.ListInventoryAudits(ctx)
	if err != nil {
		return nil, err
	}
	audits := make([]*inventory.InventoryAudit, 0, len(rows))
	for _, row := range rows {
		a := toInventoryAuditDomain(row)
		audits = append(audits, &a)
	}
	return audits, nil
}

func (r *InventoryAuditRepository) ListByWarehouseID(ctx context.Context, warehouseID uuid.UUID) ([]*inventory.InventoryAudit, error) {
	rows, err := r.q.ListInventoryAuditsByWarehouseID(ctx, warehouseID)
	if err != nil {
		return nil, err
	}
	audits := make([]*inventory.InventoryAudit, 0, len(rows))
	for _, row := range rows {
		a := toInventoryAuditDomain(row)
		audits = append(audits, &a)
	}
	return audits, nil
}

func (r *InventoryAuditRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteInventoryAudit(ctx, id)
}

func toInventoryAuditDomain(row sqlc.InventoryAudit) inventory.InventoryAudit {
	return inventory.InventoryAudit{
		ID:          row.ID,
		WarehouseID: row.WarehouseID,
		PerformedAt: row.PerformedAt,
		Result:      row.Result,
	}
}
