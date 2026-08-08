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

func TestOrderDiscountRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewOrderDiscountRepository(testDB)
	ctx := context.Background()

	d := &order.OrderDiscount{
		ID:                uuid.New(),
		OrderID:            uuid.New(),
		CouponID:           uuid.New(),
		DiscountAmountYen: 300,
	}
	require.NoError(t, repo.Create(ctx, d))

	got, err := repo.Get(ctx, d.ID)
	require.NoError(t, err)
	assert.Equal(t, d.OrderID, got.OrderID)
	assert.Equal(t, d.DiscountAmountYen, got.DiscountAmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, d.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, d.ID))
	_, err = repo.Get(ctx, d.ID)
	assert.Error(t, err)
}
