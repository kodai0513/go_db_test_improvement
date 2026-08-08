package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberPreferenceRepository は member.MemberPreferenceRepository の実装。
type MemberPreferenceRepository struct {
	q *sqlc.Queries
}

func NewMemberPreferenceRepository(db *sql.DB) *MemberPreferenceRepository {
	return &MemberPreferenceRepository{q: sqlc.New(db)}
}

var _ member.MemberPreferenceRepository = (*MemberPreferenceRepository)(nil)

func (r *MemberPreferenceRepository) Create(ctx context.Context, p *member.MemberPreference) error {
	row, err := r.q.CreateMemberPreference(ctx, sqlc.CreateMemberPreferenceParams{
		ID:              p.ID,
		MemberID:        p.MemberID,
		Language:        p.Language,
		NewsletterOptIn: p.NewsletterOptIn,
		UpdatedAt:       p.UpdatedAt,
	})
	if err != nil {
		return err
	}
	*p = toMemberPreferenceDomain(row)
	return nil
}

func (r *MemberPreferenceRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberPreference, error) {
	row, err := r.q.GetMemberPreference(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toMemberPreferenceDomain(row)
	return &p, nil
}

func (r *MemberPreferenceRepository) List(ctx context.Context) ([]*member.MemberPreference, error) {
	rows, err := r.q.ListMemberPreferences(ctx)
	if err != nil {
		return nil, err
	}
	preferences := make([]*member.MemberPreference, 0, len(rows))
	for _, row := range rows {
		p := toMemberPreferenceDomain(row)
		preferences = append(preferences, &p)
	}
	return preferences, nil
}

func (r *MemberPreferenceRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberPreference, error) {
	rows, err := r.q.ListMemberPreferencesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	preferences := make([]*member.MemberPreference, 0, len(rows))
	for _, row := range rows {
		p := toMemberPreferenceDomain(row)
		preferences = append(preferences, &p)
	}
	return preferences, nil
}

func (r *MemberPreferenceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberPreference(ctx, id)
}

func toMemberPreferenceDomain(row sqlc.MemberPreference) member.MemberPreference {
	return member.MemberPreference{
		ID:              row.ID,
		MemberID:        row.MemberID,
		Language:        row.Language,
		NewsletterOptIn: row.NewsletterOptIn,
		UpdatedAt:       row.UpdatedAt,
	}
}
