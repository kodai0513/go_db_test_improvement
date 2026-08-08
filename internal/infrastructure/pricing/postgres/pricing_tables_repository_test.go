package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/pricing"
	"example.com/go-db-test-improvement/internal/infrastructure/pricing/postgres"
)

func TestPriceRuleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPriceRuleRepository(testDB)
	ctx := context.Background()

	p := &pricing.PriceRule{
		ID:          uuid.New(),
		PriceListID: uuid.New(),
		ProductID:   uuid.New(),
		PriceYen:    1980,
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.PriceYen, got.PriceYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byList, err := repo.ListByPriceListID(ctx, p.PriceListID)
	require.NoError(t, err)
	assert.Len(t, byList, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestTaxRateRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTaxRateRepository(testDB)
	ctx := context.Background()

	tr := &pricing.TaxRate{
		ID:   uuid.New(),
		Name: "標準税率",
		Rate: 10,
	}
	require.NoError(t, repo.Create(ctx, tr))

	got, err := repo.Get(ctx, tr.ID)
	require.NoError(t, err)
	assert.Equal(t, tr.Rate, got.Rate)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, tr.ID))
	_, err = repo.Get(ctx, tr.ID)
	assert.Error(t, err)
}

func TestTaxRuleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTaxRuleRepository(testDB)
	ctx := context.Background()

	tr := &pricing.TaxRule{
		ID:        uuid.New(),
		TaxRateID: uuid.New(),
		Region:    "JP",
	}
	require.NoError(t, repo.Create(ctx, tr))

	got, err := repo.Get(ctx, tr.ID)
	require.NoError(t, err)
	assert.Equal(t, tr.Region, got.Region)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byRate, err := repo.ListByTaxRateID(ctx, tr.TaxRateID)
	require.NoError(t, err)
	assert.Len(t, byRate, 1)

	require.NoError(t, repo.Delete(ctx, tr.ID))
	_, err = repo.Get(ctx, tr.ID)
	assert.Error(t, err)
}

func TestCurrencyRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCurrencyRepository(testDB)
	ctx := context.Background()

	c := &pricing.Currency{
		ID:   uuid.New(),
		Code: "JPY",
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.Code, got.Code)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestCurrencyRateRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCurrencyRateRepository(testDB)
	ctx := context.Background()

	c := &pricing.CurrencyRate{
		ID:          uuid.New(),
		CurrencyID:  uuid.New(),
		RateToJPY:   150,
		EffectiveAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.RateToJPY, got.RateToJPY)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCurrency, err := repo.ListByCurrencyID(ctx, c.CurrencyID)
	require.NoError(t, err)
	assert.Len(t, byCurrency, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestDiscountTierRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewDiscountTierRepository(testDB)
	ctx := context.Background()

	d := &pricing.DiscountTier{
		ID:              uuid.New(),
		MinQuantity:     10,
		DiscountPercent: 5,
	}
	require.NoError(t, repo.Create(ctx, d))

	got, err := repo.Get(ctx, d.ID)
	require.NoError(t, err)
	assert.Equal(t, d.DiscountPercent, got.DiscountPercent)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, d.ID))
	_, err = repo.Get(ctx, d.ID)
	assert.Error(t, err)
}

func TestPriceHistoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPriceHistoryRepository(testDB)
	ctx := context.Background()

	p := &pricing.PriceHistory{
		ID:        uuid.New(),
		ProductID: uuid.New(),
		PriceYen:  2480,
		ChangedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.PriceYen, got.PriceYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, p.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestBundleOfferRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBundleOfferRepository(testDB)
	ctx := context.Background()

	b := &pricing.BundleOffer{
		ID:       uuid.New(),
		Name:     "お得な3点セット",
		PriceYen: 4980,
	}
	require.NoError(t, repo.Create(ctx, b))

	got, err := repo.Get(ctx, b.ID)
	require.NoError(t, err)
	assert.Equal(t, b.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, b.ID))
	_, err = repo.Get(ctx, b.ID)
	assert.Error(t, err)
}

func TestBundleOfferItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBundleOfferItemRepository(testDB)
	ctx := context.Background()

	b := &pricing.BundleOfferItem{
		ID:            uuid.New(),
		BundleOfferID: uuid.New(),
		ProductID:     uuid.New(),
		Quantity:      2,
	}
	require.NoError(t, repo.Create(ctx, b))

	got, err := repo.Get(ctx, b.ID)
	require.NoError(t, err)
	assert.Equal(t, b.Quantity, got.Quantity)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOffer, err := repo.ListByBundleOfferID(ctx, b.BundleOfferID)
	require.NoError(t, err)
	assert.Len(t, byOffer, 1)

	require.NoError(t, repo.Delete(ctx, b.ID))
	_, err = repo.Get(ctx, b.ID)
	assert.Error(t, err)
}
