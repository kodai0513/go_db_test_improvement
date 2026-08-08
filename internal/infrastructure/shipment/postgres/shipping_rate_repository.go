package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// ShippingRateRepository は shipment.ShippingRateRepository の実装。
type ShippingRateRepository struct {
	q *sqlc.Queries
}

func NewShippingRateRepository(db *sql.DB) *ShippingRateRepository {
	return &ShippingRateRepository{q: sqlc.New(db)}
}

var _ shipment.ShippingRateRepository = (*ShippingRateRepository)(nil)

func (r *ShippingRateRepository) Create(ctx context.Context, sr *shipment.ShippingRate) error {
	row, err := r.q.CreateShippingRate(ctx, sqlc.CreateShippingRateParams{
		ID:        sr.ID,
		CarrierID: sr.CarrierID,
		ZoneName:  sr.ZoneName,
		PriceYen:  sr.PriceYen,
	})
	if err != nil {
		return err
	}
	*sr = toShippingRateDomain(row)
	return nil
}

func (r *ShippingRateRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.ShippingRate, error) {
	row, err := r.q.GetShippingRate(ctx, id)
	if err != nil {
		return nil, err
	}
	sr := toShippingRateDomain(row)
	return &sr, nil
}

func (r *ShippingRateRepository) List(ctx context.Context) ([]*shipment.ShippingRate, error) {
	rows, err := r.q.ListShippingRates(ctx)
	if err != nil {
		return nil, err
	}
	rates := make([]*shipment.ShippingRate, 0, len(rows))
	for _, row := range rows {
		sr := toShippingRateDomain(row)
		rates = append(rates, &sr)
	}
	return rates, nil
}

func (r *ShippingRateRepository) ListByCarrierID(ctx context.Context, carrierID uuid.UUID) ([]*shipment.ShippingRate, error) {
	rows, err := r.q.ListShippingRatesByCarrierID(ctx, carrierID)
	if err != nil {
		return nil, err
	}
	rates := make([]*shipment.ShippingRate, 0, len(rows))
	for _, row := range rows {
		sr := toShippingRateDomain(row)
		rates = append(rates, &sr)
	}
	return rates, nil
}

func (r *ShippingRateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteShippingRate(ctx, id)
}

func toShippingRateDomain(row sqlc.ShippingRate) shipment.ShippingRate {
	return shipment.ShippingRate{
		ID:        row.ID,
		CarrierID: row.CarrierID,
		ZoneName:  row.ZoneName,
		PriceYen:  row.PriceYen,
	}
}
