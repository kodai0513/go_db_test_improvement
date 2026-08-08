package support

import (
	"context"

	"github.com/google/uuid"
)

// TicketAttachment はチケットメッセージに添付されたファイルを表す。
type TicketAttachment struct {
	ID              uuid.UUID
	TicketMessageID uuid.UUID
	URL             string
}

// TicketAttachmentRepository は ticket_attachments テーブルへのアクセスを抽象化する。
type TicketAttachmentRepository interface {
	Create(ctx context.Context, a *TicketAttachment) error
	Get(ctx context.Context, id uuid.UUID) (*TicketAttachment, error)
	List(ctx context.Context) ([]*TicketAttachment, error)
	ListByTicketMessageID(ctx context.Context, ticketMessageID uuid.UUID) ([]*TicketAttachment, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
