package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// CarrierRepository は shipment.CarrierRepository の実装。
type CarrierRepository struct {
	q *sqlc.Queries
}

func NewCarrierRepository(db *sql.DB) *CarrierRepository {
	return &CarrierRepository{q: sqlc.New(db)}
}

var _ shipment.CarrierRepository = (*CarrierRepository)(nil)

func (r *CarrierRepository) Create(ctx context.Context, c *shipment.Carrier) error {
	row, err := r.q.CreateCarrier(ctx, sqlc.CreateCarrierParams{
		ID:   c.ID,
		Name: c.Name,
	})
	if err != nil {
		return err
	}
	*c = toCarrierDomain(row)
	return nil
}

func (r *CarrierRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.Carrier, error) {
	row, err := r.q.GetCarrier(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toCarrierDomain(row)
	return &c, nil
}

func (r *CarrierRepository) List(ctx context.Context) ([]*shipment.Carrier, error) {
	rows, err := r.q.ListCarriers(ctx)
	if err != nil {
		return nil, err
	}
	carriers := make([]*shipment.Carrier, 0, len(rows))
	for _, row := range rows {
		c := toCarrierDomain(row)
		carriers = append(carriers, &c)
	}
	return carriers, nil
}

func (r *CarrierRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteCarrier(ctx, id)
}

func toCarrierDomain(row sqlc.Carrier) shipment.Carrier {
	return shipment.Carrier{
		ID:   row.ID,
		Name: row.Name,
	}
}
