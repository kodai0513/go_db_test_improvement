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

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres"
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

	p := &product.Product{
		ID:        uuid.New(),
		Name:      "Wireless Mouse",
		PriceYen:  2980,
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Name, got.Name)
	assert.Equal(t, p.PriceYen, got.PriceYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestRepository_ListProductCountAndAvgPriceByCategory(t *testing.T) {
	repo := postgres.New(testDB)
	categoryRepo := postgres.NewCategoryRepository(testDB)
	productCategoryRepo := postgres.NewProductCategoryRepository(testDB)
	ctx := context.Background()
	now := time.Now().UTC().Truncate(time.Microsecond)

	catA := &product.Category{
		ID:               uuid.New(),
		Name:             "家電-agg",
		ParentCategoryID: uuid.New(),
		CreatedAt:        now,
	}
	catB := &product.Category{
		ID:               uuid.New(),
		Name:             "書籍-agg",
		ParentCategoryID: uuid.New(),
		CreatedAt:        now,
	}
	require.NoError(t, categoryRepo.Create(ctx, catA))
	require.NoError(t, categoryRepo.Create(ctx, catB))

	pricesA := []int32{1000, 2000, 3000}
	for _, price := range pricesA {
		p := &product.Product{ID: uuid.New(), Name: "Category Agg Product A", PriceYen: price, CreatedAt: now}
		require.NoError(t, repo.Create(ctx, p))
		pc := &product.ProductCategory{ID: uuid.New(), ProductID: p.ID, CategoryID: catA.ID, CreatedAt: now}
		require.NoError(t, productCategoryRepo.Create(ctx, pc))
	}

	pB := &product.Product{ID: uuid.New(), Name: "Category Agg Product B", PriceYen: 500, CreatedAt: now}
	require.NoError(t, repo.Create(ctx, pB))
	pcB := &product.ProductCategory{ID: uuid.New(), ProductID: pB.ID, CategoryID: catB.ID, CreatedAt: now}
	require.NoError(t, productCategoryRepo.Create(ctx, pcB))

	summaries, err := repo.ListProductCountAndAvgPriceByCategory(ctx)
	require.NoError(t, err)

	var gotA, gotB *product.ProductCategorySummary
	for _, s := range summaries {
		switch s.CategoryID {
		case catA.ID:
			gotA = s
		case catB.ID:
			gotB = s
		}
	}
	require.NotNil(t, gotA)
	require.NotNil(t, gotB)

	assert.Equal(t, "家電-agg", gotA.CategoryName)
	assert.Equal(t, int64(3), gotA.ProductCount)
	assert.InDelta(t, 2000.0, gotA.AvgPriceYen, 0.001)

	assert.Equal(t, "書籍-agg", gotB.CategoryName)
	assert.Equal(t, int64(1), gotB.ProductCount)
	assert.InDelta(t, 500.0, gotB.AvgPriceYen, 0.001)
}

func TestRepository_ListVariantPriceRangeByProduct(t *testing.T) {
	repo := postgres.New(testDB)
	variantRepo := postgres.NewProductVariantRepository(testDB)
	ctx := context.Background()
	now := time.Now().UTC().Truncate(time.Microsecond)

	p1 := &product.Product{ID: uuid.New(), Name: "Variant Range Product 1-agg", PriceYen: 1000, CreatedAt: now}
	require.NoError(t, repo.Create(ctx, p1))
	for _, price := range []int32{100, 500, 250} {
		v := &product.ProductVariant{
			ID:        uuid.New(),
			ProductID: p1.ID,
			SKU:       "VARIANT-RANGE-1-AGG-" + uuid.New().String()[:8],
			PriceYen:  price,
			CreatedAt: now,
		}
		require.NoError(t, variantRepo.Create(ctx, v))
	}

	p2 := &product.Product{ID: uuid.New(), Name: "Variant Range Product 2-agg", PriceYen: 999, CreatedAt: now}
	require.NoError(t, repo.Create(ctx, p2))
	v2 := &product.ProductVariant{
		ID:        uuid.New(),
		ProductID: p2.ID,
		SKU:       "VARIANT-RANGE-2-AGG-" + uuid.New().String()[:8],
		PriceYen:  999,
		CreatedAt: now,
	}
	require.NoError(t, variantRepo.Create(ctx, v2))

	ranges, err := repo.ListVariantPriceRangeByProduct(ctx)
	require.NoError(t, err)

	var got1, got2 *product.ProductVariantPriceRange
	for _, r := range ranges {
		switch r.ProductID {
		case p1.ID:
			got1 = r
		case p2.ID:
			got2 = r
		}
	}
	require.NotNil(t, got1)
	require.NotNil(t, got2)

	assert.Equal(t, "Variant Range Product 1-agg", got1.ProductName)
	assert.Equal(t, int64(3), got1.VariantCount)
	assert.Equal(t, int32(100), got1.MinPriceYen)
	assert.Equal(t, int32(500), got1.MaxPriceYen)

	assert.Equal(t, "Variant Range Product 2-agg", got2.ProductName)
	assert.Equal(t, int64(1), got2.VariantCount)
	assert.Equal(t, int32(999), got2.MinPriceYen)
	assert.Equal(t, int32(999), got2.MaxPriceYen)
}

func TestRepository_ListProductCountByTag(t *testing.T) {
	repo := postgres.New(testDB)
	tagRepo := postgres.NewProductTagRepository(testDB)
	assignmentRepo := postgres.NewProductTagAssignmentRepository(testDB)
	ctx := context.Background()
	now := time.Now().UTC().Truncate(time.Microsecond)

	tag1 := &product.ProductTag{ID: uuid.New(), Name: "新着-agg"}
	tag2 := &product.ProductTag{ID: uuid.New(), Name: "セール-agg"}
	require.NoError(t, tagRepo.Create(ctx, tag1))
	require.NoError(t, tagRepo.Create(ctx, tag2))

	var productsForTag1 []uuid.UUID
	for i := 0; i < 3; i++ {
		p := &product.Product{ID: uuid.New(), Name: "Tag Agg Product", PriceYen: 100, CreatedAt: now}
		require.NoError(t, repo.Create(ctx, p))
		productsForTag1 = append(productsForTag1, p.ID)
		a := &product.ProductTagAssignment{ID: uuid.New(), ProductID: p.ID, ProductTagID: tag1.ID}
		require.NoError(t, assignmentRepo.Create(ctx, a))
	}

	a2 := &product.ProductTagAssignment{ID: uuid.New(), ProductID: productsForTag1[0], ProductTagID: tag2.ID}
	require.NoError(t, assignmentRepo.Create(ctx, a2))

	summaries, err := repo.ListProductCountByTag(ctx)
	require.NoError(t, err)

	var got1, got2 *product.ProductTagSummary
	for _, s := range summaries {
		switch s.ProductTagID {
		case tag1.ID:
			got1 = s
		case tag2.ID:
			got2 = s
		}
	}
	require.NotNil(t, got1)
	require.NotNil(t, got2)

	assert.Equal(t, "新着-agg", got1.TagName)
	assert.Equal(t, int64(3), got1.ProductCount)

	assert.Equal(t, "セール-agg", got2.TagName)
	assert.Equal(t, int64(1), got2.ProductCount)
}
