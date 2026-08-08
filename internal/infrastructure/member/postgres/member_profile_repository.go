package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberProfileRepository は member.MemberProfileRepository の実装。
type MemberProfileRepository struct {
	q *sqlc.Queries
}

func NewMemberProfileRepository(db *sql.DB) *MemberProfileRepository {
	return &MemberProfileRepository{q: sqlc.New(db)}
}

var _ member.MemberProfileRepository = (*MemberProfileRepository)(nil)

func (r *MemberProfileRepository) Create(ctx context.Context, p *member.MemberProfile) error {
	row, err := r.q.CreateMemberProfile(ctx, sqlc.CreateMemberProfileParams{
		ID:        p.ID,
		MemberID:  p.MemberID,
		BirthDate: p.BirthDate,
		Gender:    p.Gender,
		Bio:       p.Bio,
		AvatarURL: p.AvatarURL,
		UpdatedAt: p.UpdatedAt,
	})
	if err != nil {
		return err
	}
	*p = toMemberProfileDomain(row)
	return nil
}

func (r *MemberProfileRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberProfile, error) {
	row, err := r.q.GetMemberProfile(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toMemberProfileDomain(row)
	return &p, nil
}

func (r *MemberProfileRepository) List(ctx context.Context) ([]*member.MemberProfile, error) {
	rows, err := r.q.ListMemberProfiles(ctx)
	if err != nil {
		return nil, err
	}
	profiles := make([]*member.MemberProfile, 0, len(rows))
	for _, row := range rows {
		p := toMemberProfileDomain(row)
		profiles = append(profiles, &p)
	}
	return profiles, nil
}

func (r *MemberProfileRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberProfile, error) {
	rows, err := r.q.ListMemberProfilesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	profiles := make([]*member.MemberProfile, 0, len(rows))
	for _, row := range rows {
		p := toMemberProfileDomain(row)
		profiles = append(profiles, &p)
	}
	return profiles, nil
}

func (r *MemberProfileRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberProfile(ctx, id)
}

func toMemberProfileDomain(row sqlc.MemberProfile) member.MemberProfile {
	return member.MemberProfile{
		ID:        row.ID,
		MemberID:  row.MemberID,
		BirthDate: row.BirthDate,
		Gender:    row.Gender,
		Bio:       row.Bio,
		AvatarURL: row.AvatarURL,
		UpdatedAt: row.UpdatedAt,
	}
}
