package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberSessionRepository は member.MemberSessionRepository の実装。
type MemberSessionRepository struct {
	q *sqlc.Queries
}

func NewMemberSessionRepository(db *sql.DB) *MemberSessionRepository {
	return &MemberSessionRepository{q: sqlc.New(db)}
}

var _ member.MemberSessionRepository = (*MemberSessionRepository)(nil)

func (r *MemberSessionRepository) Create(ctx context.Context, s *member.MemberSession) error {
	row, err := r.q.CreateMemberSession(ctx, sqlc.CreateMemberSessionParams{
		ID:           s.ID,
		MemberID:     s.MemberID,
		SessionToken: s.SessionToken,
		ExpiresAt:    s.ExpiresAt,
		CreatedAt:    s.CreatedAt,
	})
	if err != nil {
		return err
	}
	*s = toMemberSessionDomain(row)
	return nil
}

func (r *MemberSessionRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberSession, error) {
	row, err := r.q.GetMemberSession(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toMemberSessionDomain(row)
	return &s, nil
}

func (r *MemberSessionRepository) List(ctx context.Context) ([]*member.MemberSession, error) {
	rows, err := r.q.ListMemberSessions(ctx)
	if err != nil {
		return nil, err
	}
	sessions := make([]*member.MemberSession, 0, len(rows))
	for _, row := range rows {
		s := toMemberSessionDomain(row)
		sessions = append(sessions, &s)
	}
	return sessions, nil
}

func (r *MemberSessionRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberSession, error) {
	rows, err := r.q.ListMemberSessionsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	sessions := make([]*member.MemberSession, 0, len(rows))
	for _, row := range rows {
		s := toMemberSessionDomain(row)
		sessions = append(sessions, &s)
	}
	return sessions, nil
}

func (r *MemberSessionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberSession(ctx, id)
}

func toMemberSessionDomain(row sqlc.MemberSession) member.MemberSession {
	return member.MemberSession{
		ID:           row.ID,
		MemberID:     row.MemberID,
		SessionToken: row.SessionToken,
		ExpiresAt:    row.ExpiresAt,
		CreatedAt:    row.CreatedAt,
	}
}
