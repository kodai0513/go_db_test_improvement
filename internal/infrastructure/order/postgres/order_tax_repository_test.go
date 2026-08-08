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

func TestOrderTaxRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewOrderTaxRepository(testDB)
	ctx := context.Background()

	tx := &order.OrderTax{
		ID:           uuid.New(),
		OrderID:      uuid.New(),
		TaxAmountYen: 480,
		TaxRate:      10,
	}
	require.NoError(t, repo.Create(ctx, tx))

	got, err := repo.Get(ctx, tx.ID)
	require.NoError(t, err)
	assert.Equal(t, tx.OrderID, got.OrderID)
	assert.Equal(t, tx.TaxAmountYen, got.TaxAmountYen)
	assert.Equal(t, tx.TaxRate, got.TaxRate)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, tx.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, tx.ID))
	_, err = repo.Get(ctx, tx.ID)
	assert.Error(t, err)
}
