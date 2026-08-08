package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres"
)

func TestMemberProfileRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberProfileRepository(testDB)
	ctx := context.Background()

	p := &member.MemberProfile{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Gender:    "female",
		Bio:       "よろしくお願いします",
		AvatarURL: "https://example.com/avatar.png",
		UpdatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Bio, got.Bio)

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

func TestMemberAddressRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberAddressRepository(testDB)
	ctx := context.Background()

	a := &member.MemberAddress{
		ID:         uuid.New(),
		MemberID:   uuid.New(),
		PostalCode: "150-0001",
		Prefecture: "東京都",
		City:       "渋谷区",
		Line1:      "神宮前1-1-1",
		Line2:      "",
		IsDefault:  true,
		CreatedAt:  time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.City, got.City)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, a.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestMemberCredentialRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberCredentialRepository(testDB)
	ctx := context.Background()

	c := &member.MemberCredential{
		ID:            uuid.New(),
		MemberID:      uuid.New(),
		PasswordHash:  "hashed-password",
		LastChangedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, c))

	got, err := repo.Get(ctx, c.ID)
	require.NoError(t, err)
	assert.Equal(t, c.PasswordHash, got.PasswordHash)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, c.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, c.ID))
	_, err = repo.Get(ctx, c.ID)
	assert.Error(t, err)
}

func TestMemberSessionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberSessionRepository(testDB)
	ctx := context.Background()

	s := &member.MemberSession{
		ID:           uuid.New(),
		MemberID:     uuid.New(),
		SessionToken: "session-token-abc",
		ExpiresAt:    time.Now().UTC().Add(24 * time.Hour).Truncate(time.Microsecond),
		CreatedAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, s))

	got, err := repo.Get(ctx, s.ID)
	require.NoError(t, err)
	assert.Equal(t, s.SessionToken, got.SessionToken)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, s.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, s.ID))
	_, err = repo.Get(ctx, s.ID)
	assert.Error(t, err)
}

func TestMemberPreferenceRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberPreferenceRepository(testDB)
	ctx := context.Background()

	p := &member.MemberPreference{
		ID:              uuid.New(),
		MemberID:        uuid.New(),
		Language:        "ja",
		NewsletterOptIn: true,
		UpdatedAt:       time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Language, got.Language)

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

func TestMemberTagRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberTagRepository(testDB)
	ctx := context.Background()

	tag := &member.MemberTag{
		ID:        uuid.New(),
		Name:      "VIP",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, tag))

	got, err := repo.Get(ctx, tag.ID)
	require.NoError(t, err)
	assert.Equal(t, tag.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, tag.ID))
	_, err = repo.Get(ctx, tag.ID)
	assert.Error(t, err)
}

func TestMemberTagAssignmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberTagAssignmentRepository(testDB)
	ctx := context.Background()

	a := &member.MemberTagAssignment{
		ID:          uuid.New(),
		MemberID:    uuid.New(),
		MemberTagID: uuid.New(),
		AssignedAt:  time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.MemberTagID, got.MemberTagID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, a.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestMemberReferralRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberReferralRepository(testDB)
	ctx := context.Background()

	r := &member.MemberReferral{
		ID:               uuid.New(),
		ReferrerMemberID: uuid.New(),
		ReferredMemberID: uuid.New(),
		RewardYen:        500,
		CreatedAt:        time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, r))

	got, err := repo.Get(ctx, r.ID)
	require.NoError(t, err)
	assert.Equal(t, r.RewardYen, got.RewardYen)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReferrer, err := repo.ListByReferrerMemberID(ctx, r.ReferrerMemberID)
	require.NoError(t, err)
	assert.Len(t, byReferrer, 1)

	require.NoError(t, repo.Delete(ctx, r.ID))
	_, err = repo.Get(ctx, r.ID)
	assert.Error(t, err)
}

func TestMemberPointAccountRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewMemberPointAccountRepository(testDB)
	ctx := context.Background()

	a := &member.MemberPointAccount{
		ID:        uuid.New(),
		MemberID:  uuid.New(),
		Balance:   1000,
		UpdatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.Balance, got.Balance)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByMemberID(ctx, a.MemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}
