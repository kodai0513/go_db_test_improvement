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

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres"
	"example.com/go-db-test-improvement/internal/testhelper"
)

var testDB *sql.DB

// TestMain boots a dedicated PostgreSQL container for this package alone.
// `go test ./...` runs one of these per infrastructure package, all at
// roughly the same time - that is the bottleneck this repository
// reproduces.
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

	m := &member.Member{
		ID:        uuid.New(),
		Name:      "Taro Yamada",
		Email:     "taro@example.com",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.Name, got.Name)
	assert.Equal(t, m.Email, got.Email)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}
