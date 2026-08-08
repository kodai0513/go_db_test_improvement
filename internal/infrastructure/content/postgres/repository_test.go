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

	"example.com/go-db-test-improvement/internal/domain/content"
	"example.com/go-db-test-improvement/internal/infrastructure/content/postgres"
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

func TestRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.New(testDB)
	ctx := context.Background()

	b := &content.Banner{
		ID:       uuid.New(),
		Title:    "夏のセール",
		ImageURL: "https://example.com/banner.png",
		StartsAt: time.Now().UTC().Truncate(time.Microsecond),
		EndsAt:   time.Now().UTC().Add(7 * 24 * time.Hour).Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, b))

	got, err := repo.Get(ctx, b.ID)
	require.NoError(t, err)
	assert.Equal(t, b.Title, got.Title)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, b.ID))
	_, err = repo.Get(ctx, b.ID)
	assert.Error(t, err)
}
