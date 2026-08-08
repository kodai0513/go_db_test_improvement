package support

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TicketAssignment はチケットの担当者割り当てを表す。
type TicketAssignment struct {
	ID            uuid.UUID
	TicketID      uuid.UUID
	StaffMemberID uuid.UUID
	AssignedAt    time.Time
}

// TicketAssignmentRepository は ticket_assignments テーブルへのアクセスを抽象化する。
type TicketAssignmentRepository interface {
	Create(ctx context.Context, a *TicketAssignment) error
	Get(ctx context.Context, id uuid.UUID) (*TicketAssignment, error)
	List(ctx context.Context) ([]*TicketAssignment, error)
	ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*TicketAssignment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
