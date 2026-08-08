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

func TestOrderNoteRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewOrderNoteRepository(testDB)
	ctx := context.Background()

	n := &order.OrderNote{
		ID:        uuid.New(),
		OrderID:   uuid.New(),
		Note:      "在庫確認済み",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, n))

	got, err := repo.Get(ctx, n.ID)
	require.NoError(t, err)
	assert.Equal(t, n.OrderID, got.OrderID)
	assert.Equal(t, n.Note, got.Note)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, n.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, n.ID))
	_, err = repo.Get(ctx, n.ID)
	assert.Error(t, err)
}
