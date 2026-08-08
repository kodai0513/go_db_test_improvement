package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// TrackingEventRepository は shipment.TrackingEventRepository の実装。
type TrackingEventRepository struct {
	q *sqlc.Queries
}

func NewTrackingEventRepository(db *sql.DB) *TrackingEventRepository {
	return &TrackingEventRepository{q: sqlc.New(db)}
}

var _ shipment.TrackingEventRepository = (*TrackingEventRepository)(nil)

func (r *TrackingEventRepository) Create(ctx context.Context, te *shipment.TrackingEvent) error {
	row, err := r.q.CreateTrackingEvent(ctx, sqlc.CreateTrackingEventParams{
		ID:         te.ID,
		ShipmentID: te.ShipmentID,
		Status:     te.Status,
		OccurredAt: te.OccurredAt,
	})
	if err != nil {
		return err
	}
	*te = toTrackingEventDomain(row)
	return nil
}

func (r *TrackingEventRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.TrackingEvent, error) {
	row, err := r.q.GetTrackingEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	te := toTrackingEventDomain(row)
	return &te, nil
}

func (r *TrackingEventRepository) List(ctx context.Context) ([]*shipment.TrackingEvent, error) {
	rows, err := r.q.ListTrackingEvents(ctx)
	if err != nil {
		return nil, err
	}
	events := make([]*shipment.TrackingEvent, 0, len(rows))
	for _, row := range rows {
		te := toTrackingEventDomain(row)
		events = append(events, &te)
	}
	return events, nil
}

func (r *TrackingEventRepository) ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*shipment.TrackingEvent, error) {
	rows, err := r.q.ListTrackingEventsByShipmentID(ctx, shipmentID)
	if err != nil {
		return nil, err
	}
	events := make([]*shipment.TrackingEvent, 0, len(rows))
	for _, row := range rows {
		te := toTrackingEventDomain(row)
		events = append(events, &te)
	}
	return events, nil
}

func (r *TrackingEventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTrackingEvent(ctx, id)
}

func toTrackingEventDomain(row sqlc.TrackingEvent) shipment.TrackingEvent {
	return shipment.TrackingEvent{
		ID:         row.ID,
		ShipmentID: row.ShipmentID,
		Status:     row.Status,
		OccurredAt: row.OccurredAt,
	}
}
