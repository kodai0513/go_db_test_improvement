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

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres"
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

	n := &notification.Notification{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Message:   "ご注文の商品を発送しました",
		IsRead:    false,
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, n))

	got, err := repo.Get(ctx, n.ID)
	require.NoError(t, err)
	assert.Equal(t, n.Message, got.Message)
	assert.Equal(t, n.IsRead, got.IsRead)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}
