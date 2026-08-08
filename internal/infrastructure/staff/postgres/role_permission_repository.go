package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// RolePermissionRepository は staff.RolePermissionRepository の実装。
type RolePermissionRepository struct {
	q *sqlc.Queries
}

func NewRolePermissionRepository(db *sql.DB) *RolePermissionRepository {
	return &RolePermissionRepository{q: sqlc.New(db)}
}

var _ staff.RolePermissionRepository = (*RolePermissionRepository)(nil)

func (r *RolePermissionRepository) Create(ctx context.Context, rp *staff.RolePermission) error {
	row, err := r.q.CreateRolePermission(ctx, sqlc.CreateRolePermissionParams{
		ID:           rp.ID,
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
	})
	if err != nil {
		return err
	}
	*rp = toRolePermissionDomain(row)
	return nil
}

func (r *RolePermissionRepository) Get(ctx context.Context, id uuid.UUID) (*staff.RolePermission, error) {
	row, err := r.q.GetRolePermission(ctx, id)
	if err != nil {
		return nil, err
	}
	rp := toRolePermissionDomain(row)
	return &rp, nil
}

func (r *RolePermissionRepository) List(ctx context.Context) ([]*staff.RolePermission, error) {
	rows, err := r.q.ListRolePermissions(ctx)
	if err != nil {
		return nil, err
	}
	permissions := make([]*staff.RolePermission, 0, len(rows))
	for _, row := range rows {
		rp := toRolePermissionDomain(row)
		permissions = append(permissions, &rp)
	}
	return permissions, nil
}

func (r *RolePermissionRepository) ListByRoleID(ctx context.Context, roleID uuid.UUID) ([]*staff.RolePermission, error) {
	rows, err := r.q.ListRolePermissionsByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}
	permissions := make([]*staff.RolePermission, 0, len(rows))
	for _, row := range rows {
		rp := toRolePermissionDomain(row)
		permissions = append(permissions, &rp)
	}
	return permissions, nil
}

func (r *RolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteRolePermission(ctx, id)
}

func toRolePermissionDomain(row sqlc.RolePermission) staff.RolePermission {
	return staff.RolePermission{
		ID:           row.ID,
		RoleID:       row.RoleID,
		PermissionID: row.PermissionID,
	}
}
