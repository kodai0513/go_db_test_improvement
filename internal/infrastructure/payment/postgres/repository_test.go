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

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres"
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

	p := &payment.Payment{
		ID:        uuid.New(),
		OrderID:   uuid.New(),
		AmountYen: 5980,
		Method:    "credit_card",
		PaidAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Method, got.Method)
	assert.Equal(t, p.AmountYen, got.AmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)
}
