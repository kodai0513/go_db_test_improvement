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

func TestCartItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCartItemRepository(testDB)
	ctx := context.Background()

	i := &order.CartItem{
		ID:        uuid.New(),
		CartID:    uuid.New(),
		ProductID: uuid.New(),
		Quantity:  3,
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.CartID, got.CartID)
	assert.Equal(t, i.ProductID, got.ProductID)
	assert.Equal(t, i.Quantity, got.Quantity)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCart, err := repo.ListByCartID(ctx, i.CartID)
	require.NoError(t, err)
	assert.Len(t, byCart, 1)

	require.NoError(t, repo.Delete(ctx, i.ID))
	_, err = repo.Get(ctx, i.ID)
	assert.Error(t, err)
}
