package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// StaffRoleRepository は staff.StaffRoleRepository の実装。
type StaffRoleRepository struct {
	q *sqlc.Queries
}

func NewStaffRoleRepository(db *sql.DB) *StaffRoleRepository {
	return &StaffRoleRepository{q: sqlc.New(db)}
}

var _ staff.StaffRoleRepository = (*StaffRoleRepository)(nil)

func (r *StaffRoleRepository) Create(ctx context.Context, sr *staff.StaffRole) error {
	row, err := r.q.CreateStaffRole(ctx, sqlc.CreateStaffRoleParams{
		ID:            sr.ID,
		StaffMemberID: sr.StaffMemberID,
		RoleID:        sr.RoleID,
	})
	if err != nil {
		return err
	}
	*sr = toStaffRoleDomain(row)
	return nil
}

func (r *StaffRoleRepository) Get(ctx context.Context, id uuid.UUID) (*staff.StaffRole, error) {
	row, err := r.q.GetStaffRole(ctx, id)
	if err != nil {
		return nil, err
	}
	sr := toStaffRoleDomain(row)
	return &sr, nil
}

func (r *StaffRoleRepository) List(ctx context.Context) ([]*staff.StaffRole, error) {
	rows, err := r.q.ListStaffRoles(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*staff.StaffRole, 0, len(rows))
	for _, row := range rows {
		sr := toStaffRoleDomain(row)
		roles = append(roles, &sr)
	}
	return roles, nil
}

func (r *StaffRoleRepository) ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*staff.StaffRole, error) {
	rows, err := r.q.ListStaffRolesByStaffMemberID(ctx, staffMemberID)
	if err != nil {
		return nil, err
	}
	roles := make([]*staff.StaffRole, 0, len(rows))
	for _, row := range rows {
		sr := toStaffRoleDomain(row)
		roles = append(roles, &sr)
	}
	return roles, nil
}

func (r *StaffRoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStaffRole(ctx, id)
}

func toStaffRoleDomain(row sqlc.StaffRole) staff.StaffRole {
	return staff.StaffRole{
		ID:            row.ID,
		StaffMemberID: row.StaffMemberID,
		RoleID:        row.RoleID,
	}
}
