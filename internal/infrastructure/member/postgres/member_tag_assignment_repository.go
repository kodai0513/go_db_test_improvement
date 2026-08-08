package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberTagAssignmentRepository は member.MemberTagAssignmentRepository の実装。
type MemberTagAssignmentRepository struct {
	q *sqlc.Queries
}

func NewMemberTagAssignmentRepository(db *sql.DB) *MemberTagAssignmentRepository {
	return &MemberTagAssignmentRepository{q: sqlc.New(db)}
}

var _ member.MemberTagAssignmentRepository = (*MemberTagAssignmentRepository)(nil)

func (r *MemberTagAssignmentRepository) Create(ctx context.Context, a *member.MemberTagAssignment) error {
	row, err := r.q.CreateMemberTagAssignment(ctx, sqlc.CreateMemberTagAssignmentParams{
		ID:          a.ID,
		MemberID:    a.MemberID,
		MemberTagID: a.MemberTagID,
		AssignedAt:  a.AssignedAt,
	})
	if err != nil {
		return err
	}
	*a = toMemberTagAssignmentDomain(row)
	return nil
}

func (r *MemberTagAssignmentRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberTagAssignment, error) {
	row, err := r.q.GetMemberTagAssignment(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toMemberTagAssignmentDomain(row)
	return &a, nil
}

func (r *MemberTagAssignmentRepository) List(ctx context.Context) ([]*member.MemberTagAssignment, error) {
	rows, err := r.q.ListMemberTagAssignments(ctx)
	if err != nil {
		return nil, err
	}
	assignments := make([]*member.MemberTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toMemberTagAssignmentDomain(row)
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *MemberTagAssignmentRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberTagAssignment, error) {
	rows, err := r.q.ListMemberTagAssignmentsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	assignments := make([]*member.MemberTagAssignment, 0, len(rows))
	for _, row := range rows {
		a := toMemberTagAssignmentDomain(row)
		assignments = append(assignments, &a)
	}
	return assignments, nil
}

func (r *MemberTagAssignmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberTagAssignment(ctx, id)
}

func toMemberTagAssignmentDomain(row sqlc.MemberTagAssignment) member.MemberTagAssignment {
	return member.MemberTagAssignment{
		ID:          row.ID,
		MemberID:    row.MemberID,
		MemberTagID: row.MemberTagID,
		AssignedAt:  row.AssignedAt,
	}
}
