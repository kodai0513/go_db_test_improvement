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

	"example.com/go-db-test-improvement/internal/domain/order"
	"example.com/go-db-test-improvement/internal/infrastructure/order/postgres"
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

	o := &order.Order{
		ID:             uuid.New(),
		MemberID:       uuid.New(),
		TotalAmountYen: 5980,
		Status:         "pending",
		CreatedAt:      time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, o))

	got, err := repo.Get(ctx, o.ID)
	require.NoError(t, err)
	assert.Equal(t, o.Status, got.Status)
	assert.Equal(t, o.TotalAmountYen, got.TotalAmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}
