package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// DeliveryAttemptRepository は shipment.DeliveryAttemptRepository の実装。
type DeliveryAttemptRepository struct {
	q *sqlc.Queries
}

func NewDeliveryAttemptRepository(db *sql.DB) *DeliveryAttemptRepository {
	return &DeliveryAttemptRepository{q: sqlc.New(db)}
}

var _ shipment.DeliveryAttemptRepository = (*DeliveryAttemptRepository)(nil)

func (r *DeliveryAttemptRepository) Create(ctx context.Context, da *shipment.DeliveryAttempt) error {
	row, err := r.q.CreateDeliveryAttempt(ctx, sqlc.CreateDeliveryAttemptParams{
		ID:          da.ID,
		ShipmentID:  da.ShipmentID,
		AttemptedAt: da.AttemptedAt,
		Succeeded:   da.Succeeded,
	})
	if err != nil {
		return err
	}
	*da = toDeliveryAttemptDomain(row)
	return nil
}

func (r *DeliveryAttemptRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.DeliveryAttempt, error) {
	row, err := r.q.GetDeliveryAttempt(ctx, id)
	if err != nil {
		return nil, err
	}
	da := toDeliveryAttemptDomain(row)
	return &da, nil
}

func (r *DeliveryAttemptRepository) List(ctx context.Context) ([]*shipment.DeliveryAttempt, error) {
	rows, err := r.q.ListDeliveryAttempts(ctx)
	if err != nil {
		return nil, err
	}
	attempts := make([]*shipment.DeliveryAttempt, 0, len(rows))
	for _, row := range rows {
		da := toDeliveryAttemptDomain(row)
		attempts = append(attempts, &da)
	}
	return attempts, nil
}

func (r *DeliveryAttemptRepository) ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*shipment.DeliveryAttempt, error) {
	rows, err := r.q.ListDeliveryAttemptsByShipmentID(ctx, shipmentID)
	if err != nil {
		return nil, err
	}
	attempts := make([]*shipment.DeliveryAttempt, 0, len(rows))
	for _, row := range rows {
		da := toDeliveryAttemptDomain(row)
		attempts = append(attempts, &da)
	}
	return attempts, nil
}

func (r *DeliveryAttemptRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteDeliveryAttempt(ctx, id)
}

func toDeliveryAttemptDomain(row sqlc.DeliveryAttempt) shipment.DeliveryAttempt {
	return shipment.DeliveryAttempt{
		ID:          row.ID,
		ShipmentID:  row.ShipmentID,
		AttemptedAt: row.AttemptedAt,
		Succeeded:   row.Succeeded,
	}
}
