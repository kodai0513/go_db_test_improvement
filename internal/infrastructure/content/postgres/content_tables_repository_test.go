package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres"
)

func TestPageRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPageRepository(testDB)
	ctx := context.Background()

	p := &content.Page{
		ID:    uuid.New(),
		Slug:  "about-us",
		Title: "会社概要",
		Body:  "当社について",
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Slug, got.Slug)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestBlogPostRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBlogPostRepository(testDB)
	ctx := context.Background()

	p := &content.BlogPost{
		ID:          uuid.New(),
		Title:       "新商品のお知らせ",
		Body:        "新商品を発売しました",
		PublishedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Title, got.Title)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestBlogCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBlogCategoryRepository(testDB)
	ctx := context.Background()

	c := &content.BlogCategory{
		ID:   uuid.New(),
		Name: "お知らせ",
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestBlogPostTagRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBlogPostTagRepository(testDB)
	ctx := context.Background()

	tag := &content.BlogPostTag{
		ID:         uuid.New(),
		BlogPostID: uuid.New(),
		Tag:        "新商品",
	}
	require.NoError(t, repo.Create(ctx, tag))

	got, err := repo.Get(ctx, tag.ID)
	require.NoError(t, err)
	assert.Equal(t, tag.Tag, got.Tag)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPost, err := repo.ListByBlogPostID(ctx, tag.BlogPostID)
	require.NoError(t, err)
	assert.Len(t, byPost, 1)

	require.NoError(t, repo.Delete(ctx, tag.ID))
	_, err = repo.Get(ctx, tag.ID)
	assert.Error(t, err)
}

func TestMediaAssetRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMediaAssetRepository(testDB)
	ctx := context.Background()

	m := &content.MediaAsset{
		ID:        uuid.New(),
		URL:       "https://example.com/asset.jpg",
		MediaType: "image/jpeg",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.MediaType, got.MediaType)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, m.ID))
	_, err = repo.Get(ctx, m.ID)
	assert.Error(t, err)
}

func TestSiteSettingRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewSiteSettingRepository(testDB)
	ctx := context.Background()

	s := &content.SiteSetting{
		ID:           uuid.New(),
		SettingKey:   "site_name",
		SettingValue: "サンプルストア",
	}
	require.NoError(t, repo.Create(ctx, s))

	got, err := repo.Get(ctx, s.ID)
	require.NoError(t, err)
	assert.Equal(t, s.SettingValue, got.SettingValue)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, s.ID))
	_, err = repo.Get(ctx, s.ID)
	assert.Error(t, err)
}

func TestMenuRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMenuRepository(testDB)
	ctx := context.Background()

	m := &content.Menu{
		ID:   uuid.New(),
		Name: "グローバルナビ",
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, m.ID))
	_, err = repo.Get(ctx, m.ID)
	assert.Error(t, err)
}

func TestMenuItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMenuItemRepository(testDB)
	ctx := context.Background()

	m := &content.MenuItem{
		ID:        uuid.New(),
		MenuID:    uuid.New(),
		Label:     "商品一覧",
		URL:       "/products",
		SortOrder: 1,
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.Label, got.Label)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMenu, err := repo.ListByMenuID(ctx, m.MenuID)
	require.NoError(t, err)
	assert.Len(t, byMenu, 1)

	require.NoError(t, repo.Delete(ctx, m.ID))
	_, err = repo.Get(ctx, m.ID)
	assert.Error(t, err)
}

func TestContentTranslationRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewContentTranslationRepository(testDB)
	ctx := context.Background()

	tr := &content.ContentTranslation{
		ID:             uuid.New(),
		PageID:         uuid.New(),
		Locale:         "en",
		TranslatedBody: "About us",
	}
	require.NoError(t, repo.Create(ctx, tr))

	got, err := repo.Get(ctx, tr.ID)
	require.NoError(t, err)
	assert.Equal(t, tr.TranslatedBody, got.TranslatedBody)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPage, err := repo.ListByPageID(ctx, tr.PageID)
	require.NoError(t, err)
	assert.Len(t, byPage, 1)

	require.NoError(t, repo.Delete(ctx, tr.ID))
	_, err = repo.Get(ctx, tr.ID)
	assert.Error(t, err)
}
