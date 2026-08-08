package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// Repository implements member.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ member.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, m *member.Member) error {
	row, err := r.q.CreateMember(ctx, sqlc.CreateMemberParams{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	})
	if err != nil {
		return err
	}
	*m = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*member.Member, error) {
	row, err := r.q.GetMember(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toDomain(row)
	return &m, nil
}

func (r *Repository) List(ctx context.Context) ([]*member.Member, error) {
	rows, err := r.q.ListMembers(ctx)
	if err != nil {
		return nil, err
	}
	members := make([]*member.Member, 0, len(rows))
	for _, row := range rows {
		m := toDomain(row)
		members = append(members, &m)
	}
	return members, nil
}

func toDomain(row sqlc.Member) member.Member {
	return member.Member{
		ID:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt,
	}
}
