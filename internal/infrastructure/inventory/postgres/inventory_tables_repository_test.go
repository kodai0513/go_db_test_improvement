package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/inventory"
	"example.com/go-db-test-improvement/internal/infrastructure/inventory/postgres"
)

func TestWarehouseRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewWarehouseRepository(testDB)
	ctx := context.Background()

	w := &inventory.Warehouse{
		ID:        uuid.New(),
		Name:      "東京第一倉庫",
		Address:   "東京都江東区豊洲1-1-1",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, w))

	got, err := repo.Get(ctx, w.ID)
	require.NoError(t, err)
	assert.Equal(t, w.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, w.ID))
	_, err = repo.Get(ctx, w.ID)
	assert.Error(t, err)
}

func TestStockMovementRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStockMovementRepository(testDB)
	ctx := context.Background()

	m := &inventory.StockMovement{
		ID:            uuid.New(),
		InventoryID:   uuid.New(),
		QuantityDelta: -5,
		Reason:        "出荷",
		CreatedAt:     time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.Reason, got.Reason)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byInventory, err := repo.ListByInventoryID(ctx, m.InventoryID)
	require.NoError(t, err)
	assert.Len(t, byInventory, 1)

	require.NoError(t, repo.Delete(ctx, m.ID))
	_, err = repo.Get(ctx, m.ID)
	assert.Error(t, err)
}

func TestStockReservationRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStockReservationRepository(testDB)
	ctx := context.Background()

	r := &inventory.StockReservation{
		ID:          uuid.New(),
		InventoryID: uuid.New(),
		OrderID:     uuid.New(),
		Quantity:    2,
		CreatedAt:   time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, r))

	got, err := repo.Get(ctx, r.ID)
	require.NoError(t, err)
	assert.Equal(t, r.Quantity, got.Quantity)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byInventory, err := repo.ListByInventoryID(ctx, r.InventoryID)
	require.NoError(t, err)
	assert.Len(t, byInventory, 1)

	byOrder, err := repo.ListByOrderID(ctx, r.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, r.ID))
	_, err = repo.Get(ctx, r.ID)
	assert.Error(t, err)
}

func TestSupplierRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewSupplierRepository(testDB)
	ctx := context.Background()

	s := &inventory.Supplier{
		ID:           uuid.New(),
		Name:         "株式会社サプライヤー",
		ContactEmail: "supplier@example.com",
		CreatedAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, s))

	got, err := repo.Get(ctx, s.ID)
	require.NoError(t, err)
	assert.Equal(t, s.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, s.ID))
	_, err = repo.Get(ctx, s.ID)
	assert.Error(t, err)
}

func TestPurchaseOrderRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPurchaseOrderRepository(testDB)
	ctx := context.Background()

	p := &inventory.PurchaseOrder{
		ID:         uuid.New(),
		SupplierID: uuid.New(),
		Status:     "pending",
		CreatedAt:  time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	bySupplier, err := repo.ListBySupplierID(ctx, p.SupplierID)
	require.NoError(t, err)
	assert.Len(t, bySupplier, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestPurchaseOrderItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPurchaseOrderItemRepository(testDB)
	ctx := context.Background()

	i := &inventory.PurchaseOrderItem{
		ID:              uuid.New(),
		PurchaseOrderID: uuid.New(),
		ProductID:       uuid.New(),
		Quantity:        10,
		UnitPriceYen:    500,
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.Quantity, got.Quantity)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByPurchaseOrderID(ctx, i.PurchaseOrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	byProduct, err := repo.ListByProductID(ctx, i.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	require.NoError(t, repo.Delete(ctx, i.ID))
	_, err = repo.Get(ctx, i.ID)
	assert.Error(t, err)
}

func TestSupplierProductRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewSupplierProductRepository(testDB)
	ctx := context.Background()

	sp := &inventory.SupplierProduct{
		ID:           uuid.New(),
		SupplierID:   uuid.New(),
		ProductID:    uuid.New(),
		LeadTimeDays: 7,
	}
	require.NoError(t, repo.Create(ctx, sp))

	got, err := repo.Get(ctx, sp.ID)
	require.NoError(t, err)
	assert.Equal(t, sp.LeadTimeDays, got.LeadTimeDays)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	bySupplier, err := repo.ListBySupplierID(ctx, sp.SupplierID)
	require.NoError(t, err)
	assert.Len(t, bySupplier, 1)

	byProduct, err := repo.ListByProductID(ctx, sp.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	require.NoError(t, repo.Delete(ctx, sp.ID))
	_, err = repo.Get(ctx, sp.ID)
	assert.Error(t, err)
}

func TestStockAdjustmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStockAdjustmentRepository(testDB)
	ctx := context.Background()

	a := &inventory.StockAdjustment{
		ID:               uuid.New(),
		InventoryID:      uuid.New(),
		AdjustedQuantity: -3,
		Note:             "棚卸差異",
		CreatedAt:        time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.Note, got.Note)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byInventory, err := repo.ListByInventoryID(ctx, a.InventoryID)
	require.NoError(t, err)
	assert.Len(t, byInventory, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestInventoryAuditRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewInventoryAuditRepository(testDB)
	ctx := context.Background()

	a := &inventory.InventoryAudit{
		ID:          uuid.New(),
		WarehouseID: uuid.New(),
		PerformedAt: time.Now().UTC().Truncate(time.Microsecond),
		Result:      "一致",
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.Result, got.Result)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byWarehouse, err := repo.ListByWarehouseID(ctx, a.WarehouseID)
	require.NoError(t, err)
	assert.Len(t, byWarehouse, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}
