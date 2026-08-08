package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres"
)

func TestTicketMessageRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTicketMessageRepository(testDB)
	ctx := context.Background()

	m := &support.TicketMessage{
		ID:             uuid.New(),
		TicketID:       uuid.New(),
		SenderMemberID: uuid.New(),
		Body:           "ご連絡ありがとうございます",
		CreatedAt:      time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, m))

	got, err := repo.Get(ctx, m.ID)
	require.NoError(t, err)
	assert.Equal(t, m.Body, got.Body)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byTicket, err := repo.ListByTicketID(ctx, m.TicketID)
	require.NoError(t, err)
	assert.Len(t, byTicket, 1)

	require.NoError(t, repo.Delete(ctx, m.ID))
	_, err = repo.Get(ctx, m.ID)
	assert.Error(t, err)
}

func TestTicketAttachmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTicketAttachmentRepository(testDB)
	ctx := context.Background()

	a := &support.TicketAttachment{
		ID:              uuid.New(),
		TicketMessageID: uuid.New(),
		URL:             "https://example.com/attachment.png",
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.URL, got.URL)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMessage, err := repo.ListByTicketMessageID(ctx, a.TicketMessageID)
	require.NoError(t, err)
	assert.Len(t, byMessage, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestTicketCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTicketCategoryRepository(testDB)
	ctx := context.Background()

	c := &support.TicketCategory{
		ID:   uuid.New(),
		Name: "アカウント",
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

func TestTicketAssignmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTicketAssignmentRepository(testDB)
	ctx := context.Background()

	a := &support.TicketAssignment{
		ID:            uuid.New(),
		TicketID:      uuid.New(),
		StaffMemberID: uuid.New(),
		AssignedAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.StaffMemberID, got.StaffMemberID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byTicket, err := repo.ListByTicketID(ctx, a.TicketID)
	require.NoError(t, err)
	assert.Len(t, byTicket, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestTicketStatusHistoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewTicketStatusHistoryRepository(testDB)
	ctx := context.Background()

	h := &support.TicketStatusHistory{
		ID:        uuid.New(),
		TicketID:  uuid.New(),
		Status:    "resolved",
		ChangedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, h))

	got, err := repo.Get(ctx, h.ID)
	require.NoError(t, err)
	assert.Equal(t, h.Status, got.Status)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byTicket, err := repo.ListByTicketID(ctx, h.TicketID)
	require.NoError(t, err)
	assert.Len(t, byTicket, 1)

	require.NoError(t, repo.Delete(ctx, h.ID))
	_, err = repo.Get(ctx, h.ID)
	assert.Error(t, err)
}

func TestFaqRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewFaqRepository(testDB)
	ctx := context.Background()

	f := &support.Faq{
		ID:        uuid.New(),
		Question:  "配送にはどれくらいかかりますか?",
		Answer:    "通常3〜5営業日です",
		CreatedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, f))

	got, err := repo.Get(ctx, f.ID)
	require.NoError(t, err)
	assert.Equal(t, f.Answer, got.Answer)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, f.ID))
	_, err = repo.Get(ctx, f.ID)
	assert.Error(t, err)
}

func TestFaqCategoryRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewFaqCategoryRepository(testDB)
	ctx := context.Background()

	c := &support.FaqCategory{
		ID:   uuid.New(),
		Name: "配送",
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

func TestSupportAgentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewSupportAgentRepository(testDB)
	ctx := context.Background()

	a := &support.SupportAgent{
		ID:    uuid.New(),
		Name:  "サポート太郎",
		Email: "support@example.com",
	}
	require.NoError(t, repo.Create(ctx, a))

	got, err := repo.Get(ctx, a.ID)
	require.NoError(t, err)
	assert.Equal(t, a.Email, got.Email)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, a.ID))
	_, err = repo.Get(ctx, a.ID)
	assert.Error(t, err)
}

func TestSatisfactionSurveyRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewSatisfactionSurveyRepository(testDB)
	ctx := context.Background()

	s := &support.SatisfactionSurvey{
		ID:          uuid.New(),
		TicketID:    uuid.New(),
		Score:       5,
		Comment:     "迅速な対応ありがとうございました",
		SubmittedAt: time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, s))

	got, err := repo.Get(ctx, s.ID)
	require.NoError(t, err)
	assert.Equal(t, s.Score, got.Score)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byTicket, err := repo.ListByTicketID(ctx, s.TicketID)
	require.NoError(t, err)
	assert.Len(t, byTicket, 1)

	require.NoError(t, repo.Delete(ctx, s.ID))
	_, err = repo.Get(ctx, s.ID)
	assert.Error(t, err)
}
