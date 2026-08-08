package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/payment"
	"example.com/go-db-test-improvement/internal/infrastructure/payment/postgres"
)

func TestPaymentMethodRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPaymentMethodRepository(testDB)
	ctx := context.Background()

	p := &payment.PaymentMethod{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Type:      "credit_card",
		IsDefault: true,
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Type, got.Type)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, p.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestPaymentTransactionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPaymentTransactionRepository(testDB)
	ctx := context.Background()

	tr := &payment.PaymentTransaction{
		ID:          uuid.New(),
		PaymentID:   uuid.New(),
		Status:      "succeeded",
		ProcessedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, tr))

	got, err := repo.Get(ctx, tr.ID)
	require.NoError(t, err)
	assert.Equal(t, tr.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPayment, err := repo.ListByPaymentID(ctx, tr.PaymentID)
	require.NoError(t, err)
	assert.Len(t, byPayment, 1)

	require.NoError(t, repo.Delete(ctx, tr.ID))
	_, err = repo.Get(ctx, tr.ID)
	assert.Error(t, err)
}

func TestRefundRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewRefundRepository(testDB)
	ctx := context.Background()

	r := &payment.Refund{
		ID:         uuid.New(),
		PaymentID:  uuid.New(),
		AmountYen:  1000,
		Reason:     "商品不良",
		RefundedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, r))

	got, err := repo.Get(ctx, r.ID)
	require.NoError(t, err)
	assert.Equal(t, r.Reason, got.Reason)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPayment, err := repo.ListByPaymentID(ctx, r.PaymentID)
	require.NoError(t, err)
	assert.Len(t, byPayment, 1)

	require.NoError(t, repo.Delete(ctx, r.ID))
	_, err = repo.Get(ctx, r.ID)
	assert.Error(t, err)
}

func TestInvoiceRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewInvoiceRepository(testDB)
	ctx := context.Background()

	i := &payment.Invoice{
		ID:        uuid.New(),
		OrderID:   uuid.New(),
		AmountYen: 5000,
		IssuedAt:  time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.AmountYen, got.AmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byOrder, err := repo.ListByOrderID(ctx, i.OrderID)
	require.NoError(t, err)
	assert.Len(t, byOrder, 1)

	require.NoError(t, repo.Delete(ctx, i.ID))
	_, err = repo.Get(ctx, i.ID)
	assert.Error(t, err)
}

func TestPaymentProviderRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPaymentProviderRepository(testDB)
	ctx := context.Background()

	p := &payment.PaymentProvider{
		ID:   uuid.New(),
		Name: "Stripe",
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestPaymentProviderAccountRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPaymentProviderAccountRepository(testDB)
	ctx := context.Background()

	a := &payment.PaymentProviderAccount{
		ID:                uuid.New(),
		PaymentProviderID: uuid.New(),
		MemberID:          uuid.New(),
		ExternalAccountID: "cus_abc123",
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.ExternalAccountID, got.ExternalAccountID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProvider, err := repo.ListByPaymentProviderID(ctx, a.PaymentProviderID)
	require.NoError(t, err)
	assert.Len(t, byProvider, 1)

	byMember, err := repo.ListByMemberID(ctx, a.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestInstallmentPlanRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewInstallmentPlanRepository(testDB)
	ctx := context.Background()

	p := &payment.InstallmentPlan{
		ID:                   uuid.New(),
		PaymentID:            uuid.New(),
		NumberOfInstallments: 3,
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.NumberOfInstallments, got.NumberOfInstallments)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPayment, err := repo.ListByPaymentID(ctx, p.PaymentID)
	require.NoError(t, err)
	assert.Len(t, byPayment, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestInstallmentPaymentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewInstallmentPaymentRepository(testDB)
	ctx := context.Background()

	p := &payment.InstallmentPayment{
		ID:                uuid.New(),
		InstallmentPlanID: uuid.New(),
		AmountYen:         3300,
		DueAt:             time.Now().UTC().Add(30 * 24 * time.Hour).Truncate(time.Microsecond),
		PaidAt:            time.Time{},
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.AmountYen, got.AmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPlan, err := repo.ListByInstallmentPlanID(ctx, p.InstallmentPlanID)
	require.NoError(t, err)
	assert.Len(t, byPlan, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestPaymentDisputeRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPaymentDisputeRepository(testDB)
	ctx := context.Background()

	d := &payment.PaymentDispute{
		ID:        uuid.New(),
		PaymentID: uuid.New(),
		Reason:    "不正利用の疑い",
		Status:    "under_review",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, d))

	got, err := repo.Get(ctx, d.ID)
	require.NoError(t, err)
	assert.Equal(t, d.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byPayment, err := repo.ListByPaymentID(ctx, d.PaymentID)
	require.NoError(t, err)
	assert.Len(t, byPayment, 1)

	require.NoError(t, repo.Delete(ctx, d.ID))
	_, err = repo.Get(ctx, d.ID)
	assert.Error(t, err)
}
