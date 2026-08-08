package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/product"
	"example.com/go-db-test-improvement/internal/infrastructure/product/postgres"
)

func TestCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCategoryRepository(testDB)
	ctx := context.Background()

	c := &product.Category{
		ID:               uuid.New(),
		Name:             "家電",
		ParentCategoryID: uuid.New(),
		CreatedAt:        time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byParent, err := repo.ListByParentCategoryID(ctx, c.ParentCategoryID)
	require.NoError(t, err)
	assert.Len(t, byParent, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestProductCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductCategoryRepository(testDB)
	ctx := context.Background()

	pc := &product.ProductCategory{
		ID:         uuid.New(),
		ProductID:  uuid.New(),
		CategoryID: uuid.New(),
		CreatedAt:  time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, pc))

	got, err := repo.Get(ctx, pc.ID)
	require.NoError(t, err)
	assert.Equal(t, pc.CategoryID, got.CategoryID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, pc.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	byCategory, err := repo.ListByCategoryID(ctx, pc.CategoryID)
	require.NoError(t, err)
	assert.Len(t, byCategory, 1)

	require.NoError(t, repo.Delete(ctx, pc.ID))
	_, err = repo.Get(ctx, pc.ID)
	assert.Error(t, err)
}

func TestProductAttributeRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductAttributeRepository(testDB)
	ctx := context.Background()

	a := &product.ProductAttribute{
		ID:        uuid.New(),
		Name:      "カラー",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestProductAttributeValueRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductAttributeValueRepository(testDB)
	ctx := context.Background()

	v := &product.ProductAttributeValue{
		ID:                 uuid.New(),
		ProductID:          uuid.New(),
		ProductAttributeID: uuid.New(),
		Value:              "レッド",
	}
	require.NoError(t, repo.Create(ctx, v))

	got, err := repo.Get(ctx, v.ID)
	require.NoError(t, err)
	assert.Equal(t, v.Value, got.Value)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, v.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	byAttribute, err := repo.ListByProductAttributeID(ctx, v.ProductAttributeID)
	require.NoError(t, err)
	assert.Len(t, byAttribute, 1)

	require.NoError(t, repo.Delete(ctx, v.ID))
	_, err = repo.Get(ctx, v.ID)
	assert.Error(t, err)
}

func TestProductVariantRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductVariantRepository(testDB)
	ctx := context.Background()

	v := &product.ProductVariant{
		ID:        uuid.New(),
		ProductID: uuid.New(),
		SKU:       "SKU-0001",
		PriceYen:  3980,
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, v))

	got, err := repo.Get(ctx, v.ID)
	require.NoError(t, err)
	assert.Equal(t, v.SKU, got.SKU)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, v.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	require.NoError(t, repo.Delete(ctx, v.ID))
	_, err = repo.Get(ctx, v.ID)
	assert.Error(t, err)
}

func TestProductImageRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductImageRepository(testDB)
	ctx := context.Background()

	img := &product.ProductImage{
		ID:        uuid.New(),
		ProductID: uuid.New(),
		URL:       "https://example.com/product.png",
		SortOrder: 1,
	}
	require.NoError(t, repo.Create(ctx, img))

	got, err := repo.Get(ctx, img.ID)
	require.NoError(t, err)
	assert.Equal(t, img.URL, got.URL)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, img.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	require.NoError(t, repo.Delete(ctx, img.ID))
	_, err = repo.Get(ctx, img.ID)
	assert.Error(t, err)
}

func TestBrandRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBrandRepository(testDB)
	ctx := context.Background()

	b := &product.Brand{
		ID:        uuid.New(),
		Name:      "サンプルブランド",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
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

func TestProductTagRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductTagRepository(testDB)
	ctx := context.Background()

	tag := &product.ProductTag{
		ID:   uuid.New(),
		Name: "新着",
	}
	require.NoError(t, repo.Create(ctx, tag))

	got, err := repo.Get(ctx, tag.ID)
	require.NoError(t, err)
	assert.Equal(t, tag.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, tag.ID))
	_, err = repo.Get(ctx, tag.ID)
	assert.Error(t, err)
}

func TestProductTagAssignmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewProductTagAssignmentRepository(testDB)
	ctx := context.Background()

	a := &product.ProductTagAssignment{
		ID:           uuid.New(),
		ProductID:    uuid.New(),
		ProductTagID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.ProductTagID, got.ProductTagID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, a.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	byTag, err := repo.ListByProductTagID(ctx, a.ProductTagID)
	require.NoError(t, err)
	assert.Len(t, byTag, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}
