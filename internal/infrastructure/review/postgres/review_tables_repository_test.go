package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/review"
	"example.com/go-db-test-improvement/internal/infrastructure/review/postgres"
)

func TestReviewImageRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewImageRepository(testDB)
	ctx := context.Background()

	i := &review.ReviewImage{
		ID:       uuid.New(),
		ReviewID: uuid.New(),
		URL:      "https://example.com/image.png",
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.URL, got.URL)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReview, err := repo.ListByReviewID(ctx, i.ReviewID)
	require.NoError(t, err)
	assert.Len(t, byReview, 1)

	require.NoError(t, repo.Delete(ctx, i.ID))
	_, err = repo.Get(ctx, i.ID)
	assert.Error(t, err)
}

func TestReviewVoteRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewVoteRepository(testDB)
	ctx := context.Background()

	v := &review.ReviewVote{
		ID:        uuid.New(),
		ReviewID:  uuid.New(),
		MemberID:  uuid.New(),
		IsHelpful: true,
	}
	require.NoError(t, repo.Create(ctx, v))

	got, err := repo.Get(ctx, v.ID)
	require.NoError(t, err)
	assert.Equal(t, v.IsHelpful, got.IsHelpful)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReview, err := repo.ListByReviewID(ctx, v.ReviewID)
	require.NoError(t, err)
	assert.Len(t, byReview, 1)

	require.NoError(t, repo.Delete(ctx, v.ID))
	_, err = repo.Get(ctx, v.ID)
	assert.Error(t, err)
}

func TestReviewReplyRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewReplyRepository(testDB)
	ctx := context.Background()

	rep := &review.ReviewReply{
		ID:        uuid.New(),
		ReviewID:  uuid.New(),
		MemberID:  uuid.New(),
		Comment:   "同感です",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, rep))

	got, err := repo.Get(ctx, rep.ID)
	require.NoError(t, err)
	assert.Equal(t, rep.Comment, got.Comment)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReview, err := repo.ListByReviewID(ctx, rep.ReviewID)
	require.NoError(t, err)
	assert.Len(t, byReview, 1)

	require.NoError(t, repo.Delete(ctx, rep.ID))
	_, err = repo.Get(ctx, rep.ID)
	assert.Error(t, err)
}

func TestReviewReportRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewReportRepository(testDB)
	ctx := context.Background()

	rp := &review.ReviewReport{
		ID:               uuid.New(),
		ReviewID:         uuid.New(),
		ReporterMemberID: uuid.New(),
		Reason:           "スパムの疑いがあります",
		CreatedAt:        time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, rp))

	got, err := repo.Get(ctx, rp.ID)
	require.NoError(t, err)
	assert.Equal(t, rp.Reason, got.Reason)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReview, err := repo.ListByReviewID(ctx, rp.ReviewID)
	require.NoError(t, err)
	assert.Len(t, byReview, 1)

	require.NoError(t, repo.Delete(ctx, rp.ID))
	_, err = repo.Get(ctx, rp.ID)
	assert.Error(t, err)
}

func TestReviewTagRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewTagRepository(testDB)
	ctx := context.Background()

	tag := &review.ReviewTag{
		ID:   uuid.New(),
		Name: "コスパ良し",
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

func TestReviewTagAssignmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewTagAssignmentRepository(testDB)
	ctx := context.Background()

	a := &review.ReviewTagAssignment{
		ID:          uuid.New(),
		ReviewID:    uuid.New(),
		ReviewTagID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.ReviewTagID, got.ReviewTagID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReview, err := repo.ListByReviewID(ctx, a.ReviewID)
	require.NoError(t, err)
	assert.Len(t, byReview, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestReviewSummaryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewSummaryRepository(testDB)
	ctx := context.Background()

	s := &review.ReviewSummary{
		ID:            uuid.New(),
		ProductID:     uuid.New(),
		AverageRating: 4,
		ReviewCount:   10,
	}
	require.NoError(t, repo.Create(ctx, s))

	got, err := repo.Get(ctx, s.ID)
	require.NoError(t, err)
	assert.Equal(t, s.AverageRating, got.AverageRating)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byProduct, err := repo.ListByProductID(ctx, s.ProductID)
	require.NoError(t, err)
	assert.Len(t, byProduct, 1)

	require.NoError(t, repo.Delete(ctx, s.ID))
	_, err = repo.Get(ctx, s.ID)
	assert.Error(t, err)
}

func TestReviewInvitationRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewInvitationRepository(testDB)
	ctx := context.Background()

	i := &review.ReviewInvitation{
		ID:       uuid.New(),
		OrderID:  uuid.New(),
		MemberID: uuid.New(),
		SentAt:   time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, i))

	got, err := repo.Get(ctx, i.ID)
	require.NoError(t, err)
	assert.Equal(t, i.MemberID, got.MemberID)

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

func TestReviewModerationLogRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewReviewModerationLogRepository(testDB)
	ctx := context.Background()

	l := &review.ReviewModerationLog{
		ID:        uuid.New(),
		ReviewID:  uuid.New(),
		Action:    "非表示化",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, l))

	got, err := repo.Get(ctx, l.ID)
	require.NoError(t, err)
	assert.Equal(t, l.Action, got.Action)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byReview, err := repo.ListByReviewID(ctx, l.ReviewID)
	require.NoError(t, err)
	assert.Len(t, byReview, 1)

	require.NoError(t, repo.Delete(ctx, l.ID))
	_, err = repo.Get(ctx, l.ID)
	assert.Error(t, err)
}
