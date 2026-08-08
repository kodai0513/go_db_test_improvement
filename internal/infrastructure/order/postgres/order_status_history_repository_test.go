package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres"
)

func TestOrderStatusHistoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewOrderStatusHistoryRepository(testDB)
	ctx := context.Background()

	h := &order.OrderStatusHistory{
		ID:        uuid.New(),
		OrderID:   uuid.New(),
		Status:    "shipped",
		ChangedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, h))

	got, err := repo.Get(ctx, h.ID)
	require.NoError(t, err)
	assert.Equal(t, h.OrderID, got.OrderID)
	assert.Equal(t, h.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, h.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, h.ID))
	_, err = repo.Get(ctx, h.ID)
	assert.Error(t, err)
}
