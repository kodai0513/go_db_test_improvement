package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres/sqlc"
)

// StaffActivityLogRepository は staff.StaffActivityLogRepository の実装。
type StaffActivityLogRepository struct {
	q *sqlc.Queries
}

func NewStaffActivityLogRepository(db *sql.DB) *StaffActivityLogRepository {
	return &StaffActivityLogRepository{q: sqlc.New(db)}
}

var _ staff.StaffActivityLogRepository = (*StaffActivityLogRepository)(nil)

func (r *StaffActivityLogRepository) Create(ctx context.Context, l *staff.StaffActivityLog) error {
	row, err := r.q.CreateStaffActivityLog(ctx, sqlc.CreateStaffActivityLogParams{
		ID:            l.ID,
		StaffMemberID: l.StaffMemberID,
		Action:        l.Action,
		OccurredAt:    l.OccurredAt,
	})
	if err != nil {
		return err
	}
	*l = toStaffActivityLogDomain(row)
	return nil
}

func (r *StaffActivityLogRepository) Get(ctx context.Context, id uuid.UUID) (*staff.StaffActivityLog, error) {
	row, err := r.q.GetStaffActivityLog(ctx, id)
	if err != nil {
		return nil, err
	}
	l := toStaffActivityLogDomain(row)
	return &l, nil
}

func (r *StaffActivityLogRepository) List(ctx context.Context) ([]*staff.StaffActivityLog, error) {
	rows, err := r.q.ListStaffActivityLogs(ctx)
	if err != nil {
		return nil, err
	}
	logs := make([]*staff.StaffActivityLog, 0, len(rows))
	for _, row := range rows {
		l := toStaffActivityLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *StaffActivityLogRepository) ListByStaffMemberID(ctx context.Context, staffMemberID uuid.UUID) ([]*staff.StaffActivityLog, error) {
	rows, err := r.q.ListStaffActivityLogsByStaffMemberID(ctx, staffMemberID)
	if err != nil {
		return nil, err
	}
	logs := make([]*staff.StaffActivityLog, 0, len(rows))
	for _, row := range rows {
		l := toStaffActivityLogDomain(row)
		logs = append(logs, &l)
	}
	return logs, nil
}

func (r *StaffActivityLogRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteStaffActivityLog(ctx, id)
}

func toStaffActivityLogDomain(row sqlc.StaffActivityLog) staff.StaffActivityLog {
	return staff.StaffActivityLog{
		ID:            row.ID,
		StaffMemberID: row.StaffMemberID,
		Action:        row.Action,
		OccurredAt:    row.OccurredAt,
	}
}
