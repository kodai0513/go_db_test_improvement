package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// StaffDepartmentRepository は staff.StaffDepartmentRepository の実装。
type StaffDepartmentRepository struct {
	q *sqlc.Queries
}

func NewStaffDepartmentRepository(db *sql.DB) *StaffDepartmentRepository {
	return &StaffDepartmentRepository{q: sqlc.New(db)}
}

var _ staff.StaffDepartmentRepository = (*StaffDepartmentRepository)(nil)

func (r *StaffDepartmentRepository) Create(ctx context.Context, sd *staff.StaffDepartment) error {
	row, err := r.q.CreateStaffDepartment(ctx, sqlc.CreateStaffDepartmentParams{
		ID:            sd.ID,
		StaffMemberID: sd.StaffMemberID,
		DepartmentID:  sd.DepartmentID,
	})
	if err != nil {
		return err
	}
	*sd = toStaffDepartmentDomain(row)
	return nil
}

func (r *StaffDepartmentRepository) Get(ctx context.Context, id uuid.UUID) (*staff.StaffDepartment, error) {
	row, err := r.q.GetStaffDepartment(ctx, id)
	if err != nil {
		return nil, err
	}
	sd := toStaffDepartmentDomain(row)
	return &sd, nil
}

func (r *StaffDepartmentRepository) List(ctx context.Context) ([]*staff.StaffDepartment, error) {
	rows, err := r.q.ListStaffDepartments(ctx)
	if err != nil {
		return nil, err
	}
	departments := make([]*staff.StaffDepartment, 0, len(rows))
	for _, row := range rows {
		sd := toStaffDepartmentDomain(row)
		departments = append(departments, &sd)
	}
	return departments, nil
}

func (r *StaffDepartmentRepository) ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*staff.StaffDepartment, error) {
	rows, err := r.q.ListStaffDepartmentsByStaffMemberID(ctx, staffMemberID)
	if err != nil {
		return nil, err
	}
	departments := make([]*staff.StaffDepartment, 0, len(rows))
	for _, row := range rows {
		sd := toStaffDepartmentDomain(row)
		departments = append(departments, &sd)
	}
	return departments, nil
}

func (r *StaffDepartmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStaffDepartment(ctx, id)
}

func toStaffDepartmentDomain(row sqlc.StaffDepartment) staff.StaffDepartment {
	return staff.StaffDepartment{
		ID:            row.ID,
		StaffMemberID: row.StaffMemberID,
		DepartmentID:  row.DepartmentID,
	}
}
