package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// Repository は staff.StaffMemberRepository の実装。
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ staff.StaffMemberRepository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, m *staff.StaffMember) error {
	row, err := r.q.CreateStaffMember(ctx, sqlc.CreateStaffMemberParams{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	})
	if err != nil {
		return err
	}
	*m = toStaffMemberDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*staff.StaffMember, error) {
	row, err := r.q.GetStaffMember(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toStaffMemberDomain(row)
	return &m, nil
}

func (r *Repository) List(ctx context.Context) ([]*staff.StaffMember, error) {
	rows, err := r.q.ListStaffMembers(ctx)
	if err != nil {
		return nil, err
	}
	members := make([]*staff.StaffMember, 0, len(rows))
	for _, row := range rows {
		m := toStaffMemberDomain(row)
		members = append(members, &m)
	}
	return members, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStaffMember(ctx, id)
}

func toStaffMemberDomain(row sqlc.StaffMember) staff.StaffMember {
	return staff.StaffMember{
		ID:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt,
	}
}
