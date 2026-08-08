package support

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TicketStatusHistory はチケットのステータス変更履歴を表す。
type TicketStatusHistory struct {
	ID        uuid.UUID
	TicketID  uuid.UUID
	Status    string
	ChangedAt time.Time
}

// TicketStatusHistoryRepository は ticket_status_histories テーブルへのアクセスを抽象化する。
type TicketStatusHistoryRepository interface {
	Create(ctx context.Context, h *TicketStatusHistory) error
	Get(ctx context.Context, id uuid.UUID) (*TicketStatusHistory, error)
	List(ctx context.Context) ([]*TicketStatusHistory, error)
	ListByTicketID(ctx context.Context, ticketID uuid.UUID) ([]*TicketStatusHistory, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
