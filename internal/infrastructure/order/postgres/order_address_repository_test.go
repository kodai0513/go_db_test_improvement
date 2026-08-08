package postgres_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres"
)

func TestOrderAddressRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewOrderAddressRepository(testDB)
	ctx := context.Background()

	a := &order.OrderAddress{
		ID:         uuid.New(),
		OrderID:    uuid.New(),
		PostalCode: "150-0001",
		Address:    "東京都渋谷区神宮前1-1-1",
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.OrderID, got.OrderID)
	assert.Equal(t, a.PostalCode, got.PostalCode)
	assert.Equal(t, a.Address, got.Address)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, a.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}
