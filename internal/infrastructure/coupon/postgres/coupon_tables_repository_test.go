package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/coupon"
	"example.com/go-db-test-improvement/internal/infrastructure/coupon/postgres"
)

func TestCouponUsageRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCouponUsageRepository(testDB)
	ctx := context.Background()

	u := &coupon.CouponUsage{
		ID:       uuid.New(),
		CouponID: uuid.New(),
		OrderID:  uuid.New(),
		UsedAt:   time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, u))

	got, err := repo.Get(ctx, u.ID)
	require.NoError(t, err)
	assert.Equal(t, u.OrderID, got.OrderID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCoupon, err := repo.ListByCouponID(ctx, u.CouponID)
	require.NoError(t, err)
	assert.Len(t, byCoupon, 1)

	require.NoError(t, repo.Delete(ctx, u.ID))
	_, err = repo.Get(ctx, u.ID)
	assert.Error(t, err)
}

func TestCampaignRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCampaignRepository(testDB)
	ctx := context.Background()

	c := &coupon.Campaign{
		ID:       uuid.New(),
		Name:     "夏の大感謝祭",
		StartsAt: time.Now().UTC().Truncate(time.Microsecond),
		EndsAt:   time.Now().UTC().Add(14 * 24 * time.Hour).Truncate(time.Microsecond),
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

func TestCampaignTargetRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCampaignTargetRepository(testDB)
	ctx := context.Background()

	tg := &coupon.CampaignTarget{
		ID:         uuid.New(),
		CampaignID: uuid.New(),
		MemberID:   uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, tg))

	got, err := repo.Get(ctx, tg.ID)
	require.NoError(t, err)
	assert.Equal(t, tg.MemberID, got.MemberID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCampaign, err := repo.ListByCampaignID(ctx, tg.CampaignID)
	require.NoError(t, err)
	assert.Len(t, byCampaign, 1)

	require.NoError(t, repo.Delete(ctx, tg.ID))
	_, err = repo.Get(ctx, tg.ID)
	assert.Error(t, err)
}

func TestDiscountRuleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewDiscountRuleRepository(testDB)
	ctx := context.Background()

	r := &coupon.DiscountRule{
		ID:           uuid.New(),
		CouponID:     uuid.New(),
		MinAmountYen: 3000,
	}
	require.NoError(t, repo.Create(ctx, r))

	got, err := repo.Get(ctx, r.ID)
	require.NoError(t, err)
	assert.Equal(t, r.MinAmountYen, got.MinAmountYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCoupon, err := repo.ListByCouponID(ctx, r.CouponID)
	require.NoError(t, err)
	assert.Len(t, byCoupon, 1)

	require.NoError(t, repo.Delete(ctx, r.ID))
	_, err = repo.Get(ctx, r.ID)
	assert.Error(t, err)
}

func TestCouponCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCouponCategoryRepository(testDB)
	ctx := context.Background()

	c := &coupon.CouponCategory{
		ID:         uuid.New(),
		CouponID:   uuid.New(),
		CategoryID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.CategoryID, got.CategoryID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCoupon, err := repo.ListByCouponID(ctx, c.CouponID)
	require.NoError(t, err)
	assert.Len(t, byCoupon, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestCouponProductRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCouponProductRepository(testDB)
	ctx := context.Background()

	p := &coupon.CouponProduct{
		ID:        uuid.New(),
		CouponID:  uuid.New(),
		ProductID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.ProductID, got.ProductID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCoupon, err := repo.ListByCouponID(ctx, p.CouponID)
	require.NoError(t, err)
	assert.Len(t, byCoupon, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestPromoCodeRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPromoCodeRepository(testDB)
	ctx := context.Background()

	p := &coupon.PromoCode{
		ID:         uuid.New(),
		Code:       "SUMMER2026",
		CampaignID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Code, got.Code)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCampaign, err := repo.ListByCampaignID(ctx, p.CampaignID)
	require.NoError(t, err)
	assert.Len(t, byCampaign, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestCouponRedemptionLimitRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCouponRedemptionLimitRepository(testDB)
	ctx := context.Background()

	l := &coupon.CouponRedemptionLimit{
		ID:             uuid.New(),
		CouponID:       uuid.New(),
		MaxRedemptions: 1,
	}
	require.NoError(t, repo.Create(ctx, l))

	got, err := repo.Get(ctx, l.ID)
	require.NoError(t, err)
	assert.Equal(t, l.MaxRedemptions, got.MaxRedemptions)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCoupon, err := repo.ListByCouponID(ctx, l.CouponID)
	require.NoError(t, err)
	assert.Len(t, byCoupon, 1)

	require.NoError(t, repo.Delete(ctx, l.ID))
	_, err = repo.Get(ctx, l.ID)
	assert.Error(t, err)
}

func TestCouponStackingRuleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewCouponStackingRuleRepository(testDB)
	ctx := context.Background()

	r := &coupon.CouponStackingRule{
		ID:                    uuid.New(),
		CouponID:              uuid.New(),
		StackableWithCouponID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, r))

	got, err := repo.Get(ctx, r.ID)
	require.NoError(t, err)
	assert.Equal(t, r.StackableWithCouponID, got.StackableWithCouponID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byCoupon, err := repo.ListByCouponID(ctx, r.CouponID)
	require.NoError(t, err)
	assert.Len(t, byCoupon, 1)

	require.NoError(t, repo.Delete(ctx, r.ID))
	_, err = repo.Get(ctx, r.ID)
	assert.Error(t, err)
}
