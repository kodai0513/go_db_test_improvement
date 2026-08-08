package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// Repository implements shipment.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ shipment.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, s *shipment.Shipment) error {
	row, err := r.q.CreateShipment(ctx, sqlc.CreateShipmentParams{
		ID:        s.ID,
		OrderID:   s.OrderID,
		Address:   s.Address,
		Status:    s.Status,
		ShippedAt: s.ShippedAt,
	})
	if err != nil {
		return err
	}
	*s = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*shipment.Shipment, error) {
	row, err := r.q.GetShipment(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toDomain(row)
	return &s, nil
}

func (r *Repository) List(ctx context.Context) ([]*shipment.Shipment, error) {
	rows, err := r.q.ListShipments(ctx)
	if err != nil {
		return nil, err
	}
	shipments := make([]*shipment.Shipment, 0, len(rows))
	for _, row := range rows {
		s := toDomain(row)
		shipments = append(shipments, &s)
	}
	return shipments, nil
}

func toDomain(row sqlc.Shipment) shipment.Shipment {
	return shipment.Shipment{
		ID:        row.ID,
		OrderID:   row.OrderID,
		Address:   row.Address,
		Status:    row.Status,
		ShippedAt: row.ShippedAt,
	}
}
