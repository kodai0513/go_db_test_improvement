package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres"
)

func TestLoyaltyTierRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewLoyaltyTierRepository(testDB)
	ctx := context.Background()

	tier := &loyalty.LoyaltyTier{
		ID:               uuid.New(),
		LoyaltyProgramID: uuid.New(),
		Name:             "ゴールド",
		MinPoints:        1000,
	}
	require.NoError(t, repo.Create(ctx, tier))

	got, err := repo.Get(ctx, tier.ID)
	require.NoError(t, err)
	assert.Equal(t, tier.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProgram, err := repo.ListByLoyaltyProgramID(ctx, tier.LoyaltyProgramID)
	require.NoError(t, err)
	assert.Len(t, byProgram, 1)

	require.NoError(t, repo.Delete(ctx, tier.ID))
	_, err = repo.Get(ctx, tier.ID)
	assert.Error(t, err)
}

func TestPointRuleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPointRuleRepository(testDB)
	ctx := context.Background()

	pr := &loyalty.PointRule{
		ID:               uuid.New(),
		LoyaltyProgramID: uuid.New(),
		Action:           "purchase",
		Points:           10,
	}
	require.NoError(t, repo.Create(ctx, pr))

	got, err := repo.Get(ctx, pr.ID)
	require.NoError(t, err)
	assert.Equal(t, pr.Action, got.Action)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProgram, err := repo.ListByLoyaltyProgramID(ctx, pr.LoyaltyProgramID)
	require.NoError(t, err)
	assert.Len(t, byProgram, 1)

	require.NoError(t, repo.Delete(ctx, pr.ID))
	_, err = repo.Get(ctx, pr.ID)
	assert.Error(t, err)
}

func TestPointRedemptionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPointRedemptionRepository(testDB)
	ctx := context.Background()

	p := &loyalty.PointRedemption{
		ID:         uuid.New(),
		MemberID:   uuid.New(),
		Points:     100,
		RedeemedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Points, got.Points)

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

func TestLoyaltyTierHistoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewLoyaltyTierHistoryRepository(testDB)
	ctx := context.Background()

	h := &loyalty.LoyaltyTierHistory{
		ID:            uuid.New(),
		MemberID:      uuid.New(),
		LoyaltyTierID: uuid.New(),
		ChangedAt:     time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, h))

	got, err := repo.Get(ctx, h.ID)
	require.NoError(t, err)
	assert.Equal(t, h.LoyaltyTierID, got.LoyaltyTierID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, h.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, h.ID))
	_, err = repo.Get(ctx, h.ID)
	assert.Error(t, err)
}

func TestRewardItemRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewRewardItemRepository(testDB)
	ctx := context.Background()

	i := &loyalty.RewardItem{
		ID:         uuid.New(),
		Name:       "500円クーポン",
		PointsCost: 500,
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, i.ID))
	_, err = repo.Get(ctx, i.ID)
	assert.Error(t, err)
}

func TestRewardRedemptionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewRewardRedemptionRepository(testDB)
	ctx := context.Background()

	rr := &loyalty.RewardRedemption{
		ID:           uuid.New(),
		MemberID:     uuid.New(),
		RewardItemID: uuid.New(),
		RedeemedAt:   time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, rr))

	got, err := repo.Get(ctx, rr.ID)
	require.NoError(t, err)
	assert.Equal(t, rr.RewardItemID, got.RewardItemID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, rr.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, rr.ID))
	_, err = repo.Get(ctx, rr.ID)
	assert.Error(t, err)
}

func TestReferralBonusRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReferralBonusRepository(testDB)
	ctx := context.Background()

	b := &loyalty.ReferralBonus{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Points:    200,
		GrantedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, b))

	got, err := repo.Get(ctx, b.ID)
	require.NoError(t, err)
	assert.Equal(t, b.Points, got.Points)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, b.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, b.ID))
	_, err = repo.Get(ctx, b.ID)
	assert.Error(t, err)
}

func TestBirthdayRewardRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewBirthdayRewardRepository(testDB)
	ctx := context.Background()

	b := &loyalty.BirthdayReward{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Points:    300,
		GrantedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, b))

	got, err := repo.Get(ctx, b.ID)
	require.NoError(t, err)
	assert.Equal(t, b.Points, got.Points)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, b.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, b.ID))
	_, err = repo.Get(ctx, b.ID)
	assert.Error(t, err)
}

func TestLoyaltyPointExpirationRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewLoyaltyPointExpirationRepository(testDB)
	ctx := context.Background()

	e := &loyalty.LoyaltyPointExpiration{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Points:    50,
		ExpiresAt: time.Now().UTC().Add(30 * 24 * time.Hour).Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, e))

	got, err := repo.Get(ctx, e.ID)
	require.NoError(t, err)
	assert.Equal(t, e.Points, got.Points)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, e.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, e.ID))
	_, err = repo.Get(ctx, e.ID)
	assert.Error(t, err)
}
