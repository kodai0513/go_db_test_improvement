package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// RoleRepository は staff.RoleRepository の実装。
type RoleRepository struct {
	q *sqlc.Queries
}

func NewRoleRepository(db *sql.DB) *RoleRepository {
	return &RoleRepository{q: sqlc.New(db)}
}

var _ staff.RoleRepository = (*RoleRepository)(nil)

func (r *RoleRepository) Create(ctx context.Context, role *staff.Role) error {
	row, err := r.q.CreateRole(ctx, sqlc.CreateRoleParams{
		ID:   role.ID,
		Name: role.Name,
	})
	if err != nil {
		return err
	}
	*role = toRoleDomain(row)
	return nil
}

func (r *RoleRepository) Get(ctx context.Context, id uuid.UUID) (*staff.Role, error) {
	row, err := r.q.GetRole(ctx, id)
	if err != nil {
		return nil, err
	}
	role := toRoleDomain(row)
	return &role, nil
}

func (r *RoleRepository) List(ctx context.Context) ([]*staff.Role, error) {
	rows, err := r.q.ListRoles(ctx)
	if err != nil {
		return nil, err
	}
	roles := make([]*staff.Role, 0, len(rows))
	for _, row := range rows {
		role := toRoleDomain(row)
		roles = append(roles, &role)
	}
	return roles, nil
}

func (r *RoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteRole(ctx, id)
}

func toRoleDomain(row sqlc.Role) staff.Role {
	return staff.Role{
		ID:   row.ID,
		Name: row.Name,
	}
}
