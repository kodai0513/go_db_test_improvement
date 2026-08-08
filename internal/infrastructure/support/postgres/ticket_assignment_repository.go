package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/support"
	"example.com/go-db-test-improvement/internal/infrastructure/support/postgres/sqlc"
)

// TicketAssignmentRepository は support.TicketAssignmentRepository の実装。
type TicketAssignmentRepository struct {
	q *sqlc.Queries
}

func NewTicketAssignmentRepository(db *sql.DB) *TicketAssignmentRepository {
	return &TicketAssignmentRepository{q: sqlc.New(db)}
}

var _ support.TicketAssignmentRepository = (*TicketAssignmentRepository)(nil)

func (r *TicketAssignmentRepository) Create(ctx context.Context, a *support.TicketAssignment) error {
	row, err := r.q.CreateTicketAssignment(ctx, sqlc.CreateTicketAssignmentParams{
		ID:            a.ID,
		TicketID:      a.TicketID,
		StaffMemberID: a.StaffMemberID,
		AssignedAt:    a.AssignedAt,
	})
	if err != nil {
		return err
	}
	*a = toTicketAssignmentDomain(row)
	return nil
}

func (r *TicketAssignmentRepository) Get(ctx context.Context, id uuid.UUID) (*support.TicketAssignment, error) {
	row, err := r.q.GetTicketAssignment(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toTicketAssignmentDomain(row)
	return &a, nil
}

func (r *TicketAssignmentRepository) List(ctx context.Context) ([]*support.TicketAssignment, error) {
	rows, err := r.q.ListTicketAssignments(ctx)
	if err != nil {
		return nil, err
	}
	assignments := make([]*support.TicketAssignment, 0, len(rows))
	for _, row := range rows {
		a := toTicketAssignmentDomain(row)
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *TicketAssignmentRepository) ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*support.TicketAssignment, error) {
	rows, err := r.q.ListTicketAssignmentsByTicketID(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	assignments := make([]*support.TicketAssignment, 0, len(rows))
	for _, row := range rows {
		a := toTicketAssignmentDomain(row)
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *TicketAssignmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteTicketAssignment(ctx, id)
}

func toTicketAssignmentDomain(row sqlc.TicketAssignment) support.TicketAssignment {
	return support.TicketAssignment{
		ID:            row.ID,
		TicketID:      row.TicketID,
		StaffMemberID: row.StaffMemberID,
		AssignedAt:    row.AssignedAt,
	}
}
