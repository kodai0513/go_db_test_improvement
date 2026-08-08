package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/merchant"
	"example.com/go-db-test-improvement/internal/infrastructure/merchant/postgres"
)

func TestVendorProductRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorProductRepository(testDB)
	ctx := context.Background()

	p := &merchant.VendorProduct{
		ID:        uuid.New(),
		VendorID:  uuid.New(),
		ProductID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.ProductID, got.ProductID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, p.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestVendorPayoutRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorPayoutRepository(testDB)
	ctx := context.Background()

	p := &merchant.VendorPayout{
		ID:        uuid.New(),
		VendorID:  uuid.New(),
		AmountYen: 50000,
		PaidAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.AmountYen, got.AmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, p.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestVendorReviewRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorReviewRepository(testDB)
	ctx := context.Background()

	rv := &merchant.VendorReview{
		ID:       uuid.New(),
		VendorID: uuid.New(),
		MemberID: uuid.New(),
		Rating:   4,
		Comment:  "対応が丁寧でした",
	}
	require.NoError(t, repo.Create(ctx, rv))

	got, err := repo.Get(ctx, rv.ID)
	require.NoError(t, err)
	assert.Equal(t, rv.Comment, got.Comment)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, rv.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, rv.ID))
	_, err = repo.Get(ctx, rv.ID)
	assert.Error(t, err)
}

func TestVendorDocumentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorDocumentRepository(testDB)
	ctx := context.Background()

	d := &merchant.VendorDocument{
		ID:       uuid.New(),
		VendorID: uuid.New(),
		URL:      "https://example.com/doc.pdf",
		DocType:  "許認可証",
	}
	require.NoError(t, repo.Create(ctx, d))

	got, err := repo.Get(ctx, d.ID)
	require.NoError(t, err)
	assert.Equal(t, d.DocType, got.DocType)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, d.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, d.ID))
	_, err = repo.Get(ctx, d.ID)
	assert.Error(t, err)
}

func TestVendorCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorCategoryRepository(testDB)
	ctx := context.Background()

	c := &merchant.VendorCategory{
		ID:         uuid.New(),
		VendorID:   uuid.New(),
		CategoryID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.CategoryID, got.CategoryID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, c.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestVendorCommissionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorCommissionRepository(testDB)
	ctx := context.Background()

	c := &merchant.VendorCommission{
		ID:          uuid.New(),
		VendorID:    uuid.New(),
		RatePercent: 8,
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.RatePercent, got.RatePercent)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, c.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestVendorBankAccountRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorBankAccountRepository(testDB)
	ctx := context.Background()

	a := &merchant.VendorBankAccount{
		ID:            uuid.New(),
		VendorID:      uuid.New(),
		AccountNumber: "1234567",
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.AccountNumber, got.AccountNumber)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, a.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestVendorContractRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorContractRepository(testDB)
	ctx := context.Background()

	c := &merchant.VendorContract{
		ID:       uuid.New(),
		VendorID: uuid.New(),
		StartsAt: time.Now().UTC().Truncate(time.Microsecond),
		EndsAt:   time.Now().UTC().Add(365 * 24 * time.Hour).Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.VendorID, got.VendorID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, c.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestVendorPerformanceMetricRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewVendorPerformanceMetricRepository(testDB)
	ctx := context.Background()

	m := &merchant.VendorPerformanceMetric{
		ID:         uuid.New(),
		VendorID:   uuid.New(),
		MetricName: "on_time_shipping_rate",
		Value:      95,
		RecordedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.Value, got.Value)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byVendor, err := repo.ListByVendorID(ctx, m.VendorID)
	require.NoError(t, err)
	assert.Len(t, byVendor, 1)

	require.NoError(t, repo.Delete(ctx, m.ID))
	_, err = repo.Get(ctx, m.ID)
	assert.Error(t, err)
}
