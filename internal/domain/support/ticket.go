package support

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Ticket は問い合わせチケットを表す。
type Ticket struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Subject   string
	Status    string
	CreatedAt time.Time
}

// TicketRepository は tickets テーブルへのアクセスを抽象化する。
type TicketRepository interface {
	Create(ctx context.Context, t *Ticket) error
	Get(ctx context.Context, id uuid.UUID) (*Ticket, error)
	List(ctx context.Context) ([]*Ticket, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*Ticket, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
