package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/shipment"
	"example.com/go-db-test-improvement/internal/infrastructure/shipment/postgres"
)

func TestShipmentItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewShipmentItemRepository(testDB)
	ctx := context.Background()

	si := &shipment.ShipmentItem{
		ID:          uuid.New(),
		ShipmentID:  uuid.New(),
		OrderItemID: uuid.New(),
		Quantity:    2,
	}
	require.NoError(t, repo.Create(ctx, si))

	got, err := repo.Get(ctx, si.ID)
	require.NoError(t, err)
	assert.Equal(t, si.Quantity, got.Quantity)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byShipment, err := repo.ListByShipmentID(ctx, si.ShipmentID)
	require.NoError(t, err)
	assert.Len(t, byShipment, 1)

	require.NoError(t, repo.Delete(ctx, si.ID))
	_, err = repo.Get(ctx, si.ID)
	assert.Error(t, err)
}

func TestCarrierRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCarrierRepository(testDB)
	ctx := context.Background()

	c := &shipment.Carrier{
		ID:   uuid.New(),
		Name: "サンプル運送",
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestTrackingEventRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTrackingEventRepository(testDB)
	ctx := context.Background()

	te := &shipment.TrackingEvent{
		ID:         uuid.New(),
		ShipmentID: uuid.New(),
		Status:     "in_transit",
		OccurredAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, te))

	got, err := repo.Get(ctx, te.ID)
	require.NoError(t, err)
	assert.Equal(t, te.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byShipment, err := repo.ListByShipmentID(ctx, te.ShipmentID)
	require.NoError(t, err)
	assert.Len(t, byShipment, 1)

	require.NoError(t, repo.Delete(ctx, te.ID))
	_, err = repo.Get(ctx, te.ID)
	assert.Error(t, err)
}

func TestShippingAddressRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewShippingAddressRepository(testDB)
	ctx := context.Background()

	sa := &shipment.ShippingAddress{
		ID:         uuid.New(),
		ShipmentID: uuid.New(),
		PostalCode: "530-0001",
		Address:    "大阪府大阪市北区梅田1-1-1",
	}
	require.NoError(t, repo.Create(ctx, sa))

	got, err := repo.Get(ctx, sa.ID)
	require.NoError(t, err)
	assert.Equal(t, sa.Address, got.Address)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byShipment, err := repo.ListByShipmentID(ctx, sa.ShipmentID)
	require.NoError(t, err)
	assert.Len(t, byShipment, 1)

	require.NoError(t, repo.Delete(ctx, sa.ID))
	_, err = repo.Get(ctx, sa.ID)
	assert.Error(t, err)
}

func TestShippingRateRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewShippingRateRepository(testDB)
	ctx := context.Background()

	sr := &shipment.ShippingRate{
		ID:        uuid.New(),
		CarrierID: uuid.New(),
		ZoneName:  "関東",
		PriceYen:  800,
	}
	require.NoError(t, repo.Create(ctx, sr))

	got, err := repo.Get(ctx, sr.ID)
	require.NoError(t, err)
	assert.Equal(t, sr.PriceYen, got.PriceYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCarrier, err := repo.ListByCarrierID(ctx, sr.CarrierID)
	require.NoError(t, err)
	assert.Len(t, byCarrier, 1)

	require.NoError(t, repo.Delete(ctx, sr.ID))
	_, err = repo.Get(ctx, sr.ID)
	assert.Error(t, err)
}

func TestShippingZoneRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewShippingZoneRepository(testDB)
	ctx := context.Background()

	sz := &shipment.ShippingZone{
		ID:   uuid.New(),
		Name: "関西エリア",
	}
	require.NoError(t, repo.Create(ctx, sz))

	got, err := repo.Get(ctx, sz.ID)
	require.NoError(t, err)
	assert.Equal(t, sz.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, sz.ID))
	_, err = repo.Get(ctx, sz.ID)
	assert.Error(t, err)
}

func TestDeliveryAttemptRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewDeliveryAttemptRepository(testDB)
	ctx := context.Background()

	da := &shipment.DeliveryAttempt{
		ID:          uuid.New(),
		ShipmentID:  uuid.New(),
		AttemptedAt: time.Now().UTC().Truncate(time.Microsecond),
		Succeeded:   false,
	}
	require.NoError(t, repo.Create(ctx, da))

	got, err := repo.Get(ctx, da.ID)
	require.NoError(t, err)
	assert.Equal(t, da.Succeeded, got.Succeeded)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byShipment, err := repo.ListByShipmentID(ctx, da.ShipmentID)
	require.NoError(t, err)
	assert.Len(t, byShipment, 1)

	require.NoError(t, repo.Delete(ctx, da.ID))
	_, err = repo.Get(ctx, da.ID)
	assert.Error(t, err)
}

func TestReturnRequestRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReturnRequestRepository(testDB)
	ctx := context.Background()

	rr := &shipment.ReturnRequest{
		ID:        uuid.New(),
		OrderID:   uuid.New(),
		Reason:    "サイズが合わなかった",
		Status:    "requested",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, rr))

	got, err := repo.Get(ctx, rr.ID)
	require.NoError(t, err)
	assert.Equal(t, rr.Reason, got.Reason)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, rr.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, rr.ID))
	_, err = repo.Get(ctx, rr.ID)
	assert.Error(t, err)
}

func TestReturnItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReturnItemRepository(testDB)
	ctx := context.Background()

	ri := &shipment.ReturnItem{
		ID:              uuid.New(),
		ReturnRequestID: uuid.New(),
		OrderItemID:     uuid.New(),
		Quantity:        1,
	}
	require.NoError(t, repo.Create(ctx, ri))

	got, err := repo.Get(ctx, ri.ID)
	require.NoError(t, err)
	assert.Equal(t, ri.Quantity, got.Quantity)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byRequest, err := repo.ListByReturnRequestID(ctx, ri.ReturnRequestID)
	require.NoError(t, err)
	assert.Len(t, byRequest, 1)

	require.NoError(t, repo.Delete(ctx, ri.ID))
	_, err = repo.Get(ctx, ri.ID)
	assert.Error(t, err)
}
