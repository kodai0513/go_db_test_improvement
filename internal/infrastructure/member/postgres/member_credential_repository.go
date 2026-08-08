package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberCredentialRepository は member.MemberCredentialRepository の実装。
type MemberCredentialRepository struct {
	q *sqlc.Queries
}

func NewMemberCredentialRepository(db *sql.DB) *MemberCredentialRepository {
	return &MemberCredentialRepository{q: sqlc.New(db)}
}

var _ member.MemberCredentialRepository = (*MemberCredentialRepository)(nil)

func (r *MemberCredentialRepository) Create(ctx context.Context, c *member.MemberCredential) error {
	row, err := r.q.CreateMemberCredential(ctx, sqlc.CreateMemberCredentialParams{
		ID:            c.ID,
		MemberID:      c.MemberID,
		PasswordHash:  c.PasswordHash,
		LastChangedAt: c.LastChangedAt,
	})
	if err != nil {
		return err
	}
	*c = toMemberCredentialDomain(row)
	return nil
}

func (r *MemberCredentialRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberCredential, error) {
	row, err := r.q.GetMemberCredential(ctx, id)
	if err != nil {
		return nil, err
	}
	c := toMemberCredentialDomain(row)
	return &c, nil
}

func (r *MemberCredentialRepository) List(ctx context.Context) ([]*member.MemberCredential, error) {
	rows, err := r.q.ListMemberCredentials(ctx)
	if err != nil {
		return nil, err
	}
	credentials := make([]*member.MemberCredential, 0, len(rows))
	for _, row := range rows {
		c := toMemberCredentialDomain(row)
		credentials = append(credentials, &c)
	}
	return credentials, nil
}

func (r *MemberCredentialRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberCredential, error) {
	rows, err := r.q.ListMemberCredentialsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	credentials := make([]*member.MemberCredential, 0, len(rows))
	for _, row := range rows {
		c := toMemberCredentialDomain(row)
		credentials = append(credentials, &c)
	}
	return credentials, nil
}

func (r *MemberCredentialRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberCredential(ctx, id)
}

func toMemberCredentialDomain(row sqlc.MemberCredential) member.MemberCredential {
	return member.MemberCredential{
		ID:            row.ID,
		MemberID:      row.MemberID,
		PasswordHash:  row.PasswordHash,
		LastChangedAt: row.LastChangedAt,
	}
}
