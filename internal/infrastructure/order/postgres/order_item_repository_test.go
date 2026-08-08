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

func TestOrderItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewOrderItemRepository(testDB)
	ctx := context.Background()

	i := &order.OrderItem{
		ID:           uuid.New(),
		OrderID:      uuid.New(),
		ProductID:    uuid.New(),
		Quantity:     2,
		UnitPriceYen: 1980,
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.OrderID, got.OrderID)
	assert.Equal(t, i.ProductID, got.ProductID)
	assert.Equal(t, i.Quantity, got.Quantity)
	assert.Equal(t, i.UnitPriceYen, got.UnitPriceYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, i.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, i.ID))
	_, err = repo.Get(ctx, i.ID)
	assert.Error(t, err)
}
