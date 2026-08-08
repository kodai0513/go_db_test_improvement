package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// DepartmentRepository は staff.DepartmentRepository の実装。
type DepartmentRepository struct {
	q *sqlc.Queries
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{q: sqlc.New(db)}
}

var _ staff.DepartmentRepository = (*DepartmentRepository)(nil)

func (r *DepartmentRepository) Create(ctx context.Context, d *staff.Department) error {
	row, err := r.q.CreateDepartment(ctx, sqlc.CreateDepartmentParams{
		ID:   d.ID,
		Name: d.Name,
	})
	if err != nil {
		return err
	}
	*d = toDepartmentDomain(row)
	return nil
}

func (r *DepartmentRepository) Get(ctx context.Context, id uuid.UUID) (*staff.Department, error) {
	row, err := r.q.GetDepartment(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toDepartmentDomain(row)
	return &d, nil
}

func (r *DepartmentRepository) List(ctx context.Context) ([]*staff.Department, error) {
	rows, err := r.q.ListDepartments(ctx)
	if err != nil {
		return nil, err
	}
	departments := make([]*staff.Department, 0, len(rows))
	for _, row := range rows {
		d := toDepartmentDomain(row)
		departments = append(departments, &d)
	}
	return departments, nil
}

func (r *DepartmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteDepartment(ctx, id)
}

func toDepartmentDomain(row sqlc.Department) staff.Department {
	return staff.Department{
		ID:   row.ID,
		Name: row.Name,
	}
}
