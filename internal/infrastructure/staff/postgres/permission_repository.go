package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// PermissionRepository は staff.PermissionRepository の実装。
type PermissionRepository struct {
	q *sqlc.Queries
}

func NewPermissionRepository(db *sql.DB) *PermissionRepository {
	return &PermissionRepository{q: sqlc.New(db)}
}

var _ staff.PermissionRepository = (*PermissionRepository)(nil)

func (r *PermissionRepository) Create(ctx context.Context, p *staff.Permission) error {
	row, err := r.q.CreatePermission(ctx, sqlc.CreatePermissionParams{
		ID:   p.ID,
		Name: p.Name,
	})
	if err != nil {
		return err
	}
	*p = toPermissionDomain(row)
	return nil
}

func (r *PermissionRepository) Get(ctx context.Context, id uuid.UUID) (*staff.Permission, error) {
	row, err := r.q.GetPermission(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPermissionDomain(row)
	return &p, nil
}

func (r *PermissionRepository) List(ctx context.Context) ([]*staff.Permission, error) {
	rows, err := r.q.ListPermissions(ctx)
	if err != nil {
		return nil, err
	}
	permissions := make([]*staff.Permission, 0, len(rows))
	for _, row := range rows {
		p := toPermissionDomain(row)
		permissions = append(permissions, &p)
	}
	return permissions, nil
}

func (r *PermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePermission(ctx, id)
}

func toPermissionDomain(row sqlc.Permission) staff.Permission {
	return staff.Permission{
		ID:   row.ID,
		Name: row.Name,
	}
}
