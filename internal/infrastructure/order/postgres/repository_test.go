package postgres_test

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres"
	"example.com/go-db-test-improvement/internal/testhelper"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	db, cleanup, err := testhelper.StartPostgres()
	if err != nil {
		panic(err)
	}
	testDB = db
	code := m.Run()
	cleanup()
	os.Exit(code)
}

func TestRepository_CreateGetList(t *testing.T) {
	repo := postgres.New(testDB)
	ctx := context.Background()

	o := &order.Order{
		ID:             uuid.New(),
		MemberID:       uuid.New(),
		TotalAmountYen: 5980,
		Status:         "pending",
		CreatedAt:      time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, o))

	got, err := repo.Get(ctx, o.ID)
	require.NoError(t, err)
	assert.Equal(t, o.Status, got.Status)
	assert.Equal(t, o.TotalAmountYen, got.TotalAmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestRepository_ListOrderSummaryByMember(t *testing.T) {
	repo := postgres.New(testDB)
	ctx := context.Background()

	memberID := uuid.New()
	orders := []*order.Order{
		{
			ID:             uuid.New(),
			MemberID:       memberID,
			TotalAmountYen: 1000,
			Status:         "pending-agg",
			CreatedAt:      time.Now().UTC().Truncate(time.Microsecond),
		},
		{
			ID:             uuid.New(),
			MemberID:       memberID,
			TotalAmountYen: 2000,
			Status:         "pending-agg",
			CreatedAt:      time.Now().UTC().Truncate(time.Microsecond),
		},
	}
	for _, o := range orders {
		require.NoError(t, repo.Create(ctx, o))
	}

	summaries, err := repo.ListOrderSummaryByMember(ctx)
	require.NoError(t, err)

	var found *order.OrderSummaryByMember
	for _, s := range summaries {
		if s.MemberID == memberID {
			found = s
			break
		}
	}
	require.NotNil(t, found)
	assert.Equal(t, int64(2), found.OrderCount)
	assert.Equal(t, int64(3000), found.TotalAmountYen)
	assert.InDelta(t, 1500.0, found.AvgAmountYen, 0.001)
}

func TestRepository_ListOrderCountByStatus(t *testing.T) {
	repo := postgres.New(testDB)
	ctx := context.Background()

	status := "shipped-agg"
	for i := 0; i < 3; i++ {
		o := &order.Order{
			ID:             uuid.New(),
			MemberID:       uuid.New(),
			TotalAmountYen: 1000,
			Status:         status,
			CreatedAt:      time.Now().UTC().Truncate(time.Microsecond),
		}
		require.NoError(t, repo.Create(ctx, o))
	}

	counts, err := repo.ListOrderCountByStatus(ctx)
	require.NoError(t, err)

	var found *order.OrderCountByStatus
	for _, c := range counts {
		if c.Status == status {
			found = c
			break
		}
	}
	require.NotNil(t, found)
	assert.Equal(t, int64(3), found.OrderCount)
}

func TestRepository_ListSalesQuantityAndAmountByProduct(t *testing.T) {
	itemRepo := postgres.NewOrderItemRepository(testDB)
	ctx := context.Background()

	productID := uuid.New()
	items := []*order.OrderItem{
		{
			ID:           uuid.New(),
			OrderID:      uuid.New(),
			ProductID:    productID,
			Quantity:     2,
			UnitPriceYen: 1000,
		},
		{
			ID:           uuid.New(),
			OrderID:      uuid.New(),
			ProductID:    productID,
			Quantity:     3,
			UnitPriceYen: 1000,
		},
	}
	for _, i := range items {
		require.NoError(t, itemRepo.Create(ctx, i))
	}

	repo := postgres.New(testDB)
	sales, err := repo.ListSalesQuantityAndAmountByProduct(ctx)
	require.NoError(t, err)

	var found *order.SalesQuantityAndAmountByProduct
	for _, s := range sales {
		if s.ProductID == productID {
			found = s
			break
		}
	}
	require.NotNil(t, found)
	assert.Equal(t, int64(5), found.TotalQuantity)
	assert.Equal(t, int64(5000), found.TotalAmountYen)
}
