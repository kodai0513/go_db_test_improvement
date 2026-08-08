package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// ShippingZoneRepository は shipment.ShippingZoneRepository の実装。
type ShippingZoneRepository struct {
	q *sqlc.Queries
}

func NewShippingZoneRepository(db *sql.DB) *ShippingZoneRepository {
	return &ShippingZoneRepository{q: sqlc.New(db)}
}

var _ shipment.ShippingZoneRepository = (*ShippingZoneRepository)(nil)

func (r *ShippingZoneRepository) Create(ctx context.Context, sz *shipment.ShippingZone) error {
	row, err := r.q.CreateShippingZone(ctx, sqlc.CreateShippingZoneParams{
		ID:   sz.ID,
		Name: sz.Name,
	})
	if err != nil {
		return err
	}
	*sz = toShippingZoneDomain(row)
	return nil
}

func (r *ShippingZoneRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.ShippingZone, error) {
	row, err := r.q.GetShippingZone(ctx, id)
	if err != nil {
		return nil, err
	}
	sz := toShippingZoneDomain(row)
	return &sz, nil
}

func (r *ShippingZoneRepository) List(ctx context.Context) ([]*shipment.ShippingZone, error) {
	rows, err := r.q.ListShippingZones(ctx)
	if err != nil {
		return nil, err
	}
	zones := make([]*shipment.ShippingZone, 0, len(rows))
	for _, row := range rows {
		sz := toShippingZoneDomain(row)
		zones = append(zones, &sz)
	}
	return zones, nil
}

func (r *ShippingZoneRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteShippingZone(ctx, id)
}

func toShippingZoneDomain(row sqlc.ShippingZone) shipment.ShippingZone {
	return shipment.ShippingZone{
		ID:   row.ID,
		Name: row.Name,
	}
}
