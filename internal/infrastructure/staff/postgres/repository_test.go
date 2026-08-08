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

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres"
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

	m := &staff.StaffMember{
		ID:        uuid.New(),
		Name:      "山田花子",
		Email:     "hanako@example.com",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
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
