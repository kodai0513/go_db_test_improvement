package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres/sqlc"
)

// Repository implements inventory.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ inventory.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, i *inventory.Inventory) error {
	row, err := r.q.CreateInventory(ctx, sqlc.CreateInventoryParams{
		ID:        i.ID,
		ProductID: i.ProductID,
		Quantity:  i.Quantity,
		UpdatedAt: i.UpdatedAt,
	})
	if err != nil {
		return err
	}
	*i = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*inventory.Inventory, error) {
	row, err := r.q.GetInventory(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toDomain(row)
	return &i, nil
}

func (r *Repository) List(ctx context.Context) ([]*inventory.Inventory, error) {
	rows, err := r.q.ListInventories(ctx)
	if err != nil {
		return nil, err
	}
	inventories := make([]*inventory.Inventory, 0, len(rows))
	for _, row := range rows {
		i := toDomain(row)
		inventories = append(inventories, &i)
	}
	return inventories, nil
}

func toDomain(row sqlc.Inventory) inventory.Inventory {
	return inventory.Inventory{
		ID:        row.ID,
		ProductID: row.ProductID,
		Quantity:  row.Quantity,
		UpdatedAt: row.UpdatedAt,
	}
}
