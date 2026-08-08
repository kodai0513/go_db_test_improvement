package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// StaffScheduleRepository は staff.StaffScheduleRepository の実装。
type StaffScheduleRepository struct {
	q *sqlc.Queries
}

func NewStaffScheduleRepository(db *sql.DB) *StaffScheduleRepository {
	return &StaffScheduleRepository{q: sqlc.New(db)}
}

var _ staff.StaffScheduleRepository = (*StaffScheduleRepository)(nil)

func (r *StaffScheduleRepository) Create(ctx context.Context, s *staff.StaffSchedule) error {
	row, err := r.q.CreateStaffSchedule(ctx, sqlc.CreateStaffScheduleParams{
		ID:            s.ID,
		StaffMemberID: s.StaffMemberID,
		WorkDate:      s.WorkDate,
		Shift:         s.Shift,
	})
	if err != nil {
		return err
	}
	*s = toStaffScheduleDomain(row)
	return nil
}

func (r *StaffScheduleRepository) Get(ctx context.Context, id uuid.UUID) (*staff.StaffSchedule, error) {
	row, err := r.q.GetStaffSchedule(ctx, id)
	if err != nil {
		return nil, err
	}
	s := toStaffScheduleDomain(row)
	return &s, nil
}

func (r *StaffScheduleRepository) List(ctx context.Context) ([]*staff.StaffSchedule, error) {
	rows, err := r.q.ListStaffSchedules(ctx)
	if err != nil {
		return nil, err
	}
	schedules := make([]*staff.StaffSchedule, 0, len(rows))
	for _, row := range rows {
		s := toStaffScheduleDomain(row)
		schedules = append(schedules, &s)
	}
	return schedules, nil
}

func (r *StaffScheduleRepository) ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*staff.StaffSchedule, error) {
	rows, err := r.q.ListStaffSchedulesByStaffMemberID(ctx, staffMemberID)
	if err != nil {
		return nil, err
	}
	schedules := make([]*staff.StaffSchedule, 0, len(rows))
	for _, row := range rows {
		s := toStaffScheduleDomain(row)
		schedules = append(schedules, &s)
	}
	return schedules, nil
}

func (r *StaffScheduleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStaffSchedule(ctx, id)
}

func toStaffScheduleDomain(row sqlc.StaffSchedule) staff.StaffSchedule {
	return staff.StaffSchedule{
		ID:            row.ID,
		StaffMemberID: row.StaffMemberID,
		WorkDate:      row.WorkDate,
		Shift:         row.Shift,
	}
}
