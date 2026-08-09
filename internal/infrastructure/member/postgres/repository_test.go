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

func TestRepository_ListMemberCountAndAvgPointBalanceByTag(t *testing.T) {
	repo := postgres.New(testDB)
	tagRepo := postgres.NewMemberTagRepository(testDB)
	assignmentRepo := postgres.NewMemberTagAssignmentRepository(testDB)
	pointRepo := postgres.NewMemberPointAccountRepository(testDB)
	ctx := context.Background()

	vip := &member.MemberTag{ID: uuid.New(), Name: "VIP-agg", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, tagRepo.Create(ctx, vip))
	newcomer := &member.MemberTag{ID: uuid.New(), Name: "Newcomer-agg", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, tagRepo.Create(ctx, newcomer))

	memberA := &member.Member{ID: uuid.New(), Name: "Agg A", Email: "agg-a@example.com", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, repo.Create(ctx, memberA))
	memberB := &member.Member{ID: uuid.New(), Name: "Agg B", Email: "agg-b@example.com", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, repo.Create(ctx, memberB))

	require.NoError(t, assignmentRepo.Create(ctx, &member.MemberTagAssignment{ID: uuid.New(), MemberID: memberA.ID, MemberTagID: vip.ID, AssignedAt: time.Now().UTC().Truncate(time.Microsecond)}))
	require.NoError(t, assignmentRepo.Create(ctx, &member.MemberTagAssignment{ID: uuid.New(), MemberID: memberB.ID, MemberTagID: vip.ID, AssignedAt: time.Now().UTC().Truncate(time.Microsecond)}))

	require.NoError(t, pointRepo.Create(ctx, &member.MemberPointAccount{ID: uuid.New(), MemberID: memberA.ID, Balance: 100, UpdatedAt: time.Now().UTC().Truncate(time.Microsecond)}))
	require.NoError(t, pointRepo.Create(ctx, &member.MemberPointAccount{ID: uuid.New(), MemberID: memberB.ID, Balance: 300, UpdatedAt: time.Now().UTC().Truncate(time.Microsecond)}))

	summaries, err := repo.ListMemberCountAndAvgPointBalanceByTag(ctx)
	require.NoError(t, err)

	var vipSummary *member.MemberTagPointSummary
	for _, s := range summaries {
		if s.MemberTagID == vip.ID {
			vipSummary = s
		}
	}
	require.NotNil(t, vipSummary)
	assert.EqualValues(t, 2, vipSummary.MemberCount)
	assert.InDelta(t, 200.0, vipSummary.AvgPointBalance, 0.001)
}

func TestRepository_ListTopMembersByAddressCountAndPointBalance(t *testing.T) {
	repo := postgres.New(testDB)
	addressRepo := postgres.NewMemberAddressRepository(testDB)
	pointRepo := postgres.NewMemberPointAccountRepository(testDB)
	ctx := context.Background()

	rich := &member.Member{ID: uuid.New(), Name: "Rich Ranker", Email: "rich-ranker@example.com", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, repo.Create(ctx, rich))
	poor := &member.Member{ID: uuid.New(), Name: "Poor Ranker", Email: "poor-ranker@example.com", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, repo.Create(ctx, poor))

	require.NoError(t, pointRepo.Create(ctx, &member.MemberPointAccount{ID: uuid.New(), MemberID: rich.ID, Balance: 5000, UpdatedAt: time.Now().UTC().Truncate(time.Microsecond)}))
	require.NoError(t, pointRepo.Create(ctx, &member.MemberPointAccount{ID: uuid.New(), MemberID: poor.ID, Balance: 10, UpdatedAt: time.Now().UTC().Truncate(time.Microsecond)}))

	require.NoError(t, addressRepo.Create(ctx, &member.MemberAddress{
		ID: uuid.New(), MemberID: rich.ID, PostalCode: "100-0001", Prefecture: "東京都", City: "千代田区", Line1: "1-1-1",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}))

	ranking, err := repo.ListTopMembersByAddressCountAndPointBalance(ctx, 1)
	require.NoError(t, err)
	require.Len(t, ranking, 1)
	assert.Equal(t, rich.ID, ranking[0].MemberID)
	assert.EqualValues(t, 1, ranking[0].AddressCount)
	assert.EqualValues(t, 5000, ranking[0].PointBalance)
}

func TestRepository_ListReferralRewardSummaryByReferrer(t *testing.T) {
	repo := postgres.New(testDB)
	referralRepo := postgres.NewMemberReferralRepository(testDB)
	ctx := context.Background()

	referrer := &member.Member{ID: uuid.New(), Name: "Referrer", Email: "referrer-agg@example.com", CreatedAt: time.Now().UTC().Truncate(time.Microsecond)}
	require.NoError(t, repo.Create(ctx, referrer))

	require.NoError(t, referralRepo.Create(ctx, &member.MemberReferral{
		ID: uuid.New(), ReferrerMemberID: referrer.ID, ReferredMemberID: uuid.New(), RewardYen: 500, CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}))
	require.NoError(t, referralRepo.Create(ctx, &member.MemberReferral{
		ID: uuid.New(), ReferrerMemberID: referrer.ID, ReferredMemberID: uuid.New(), RewardYen: 300, CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}))

	summaries, err := repo.ListReferralRewardSummaryByReferrer(ctx)
	require.NoError(t, err)

	var found *member.ReferralRewardSummary
	for _, s := range summaries {
		if s.ReferrerMemberID == referrer.ID {
			found = s
		}
	}
	require.NotNil(t, found)
	assert.EqualValues(t, 2, found.ReferralCount)
	assert.EqualValues(t, 800, found.TotalRewardYen)
}
