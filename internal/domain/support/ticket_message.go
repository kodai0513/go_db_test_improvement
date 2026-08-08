package support

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TicketMessage はチケットに紐づくメッセージを表す。
type TicketMessage struct {
	ID             uuid.UUID
	TicketID       uuid.UUID
	SenderMemberID uuid.UUID
	Body           string
	CreatedAt      time.Time
}

// TicketMessageRepository は ticket_messages テーブルへのアクセスを抽象化する。
type TicketMessageRepository interface {
	Create(ctx context.Context, m *TicketMessage) error
	Get(ctx context.Context, id uuid.UUID) (*TicketMessage, error)
	List(ctx context.Context) ([]*TicketMessage, error)
	ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*TicketMessage, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
