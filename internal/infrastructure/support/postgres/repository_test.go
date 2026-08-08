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

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres"
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

	tk := &support.Ticket{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Subject:   "ログインできません",
		Status:    "open",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, tk))

	got, err := repo.Get(ctx, tk.ID)
	require.NoError(t, err)
	assert.Equal(t, tk.Subject, got.Subject)
	assert.Equal(t, tk.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, tk.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, tk.ID))
	_, err = repo.Get(ctx, tk.ID)
	assert.Error(t, err)
}
