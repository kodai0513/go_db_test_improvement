package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres/sqlc"
)

// ShippingAddressRepository は shipment.ShippingAddressRepository の実装。
type ShippingAddressRepository struct {
	q *sqlc.Queries
}

func NewShippingAddressRepository(db *sql.DB) *ShippingAddressRepository {
	return &ShippingAddressRepository{q: sqlc.New(db)}
}

var _ shipment.ShippingAddressRepository = (*ShippingAddressRepository)(nil)

func (r *ShippingAddressRepository) Create(ctx context.Context, sa *shipment.ShippingAddress) error {
	row, err := r.q.CreateShippingAddress(ctx, sqlc.CreateShippingAddressParams{
		ID:         sa.ID,
		ShipmentID: sa.ShipmentID,
		PostalCode: sa.PostalCode,
		Address:    sa.Address,
	})
	if err != nil {
		return err
	}
	*sa = toShippingAddressDomain(row)
	return nil
}

func (r *ShippingAddressRepository) Get(ctx context.Context, id uuid.UUID) (*shipment.ShippingAddress, error) {
	row, err := r.q.GetShippingAddress(ctx, id)
	if err != nil {
		return nil, err
	}
	sa := toShippingAddressDomain(row)
	return &sa, nil
}

func (r *ShippingAddressRepository) List(ctx context.Context) ([]*shipment.ShippingAddress, error) {
	rows, err := r.q.ListShippingAddresses(ctx)
	if err != nil {
		return nil, err
	}
	addresses := make([]*shipment.ShippingAddress, 0, len(rows))
	for _, row := range rows {
		sa := toShippingAddressDomain(row)
		addresses = append(addresses, &sa)
	}
	return addresses, nil
}

func (r *ShippingAddressRepository) ListByShipmentID(ctx context.Context, shipmentID uuid.UUID) ([]*shipment.ShippingAddress, error) {
	rows, err := r.q.ListShippingAddressesByShipmentID(ctx, shipmentID)
	if err != nil {
		return nil, err
	}
	addresses := make([]*shipment.ShippingAddress, 0, len(rows))
	for _, row := range rows {
		sa := toShippingAddressDomain(row)
		addresses = append(addresses, &sa)
	}
	return addresses, nil
}

func (r *ShippingAddressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteShippingAddress(ctx, id)
}

func toShippingAddressDomain(row sqlc.ShippingAddress) shipment.ShippingAddress {
	return shipment.ShippingAddress{
		ID:         row.ID,
		ShipmentID: row.ShipmentID,
		PostalCode: row.PostalCode,
		Address:    row.Address,
	}
}
