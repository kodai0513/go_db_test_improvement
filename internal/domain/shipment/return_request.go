package shipment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReturnRequest は返品依頼を表す。
type ReturnRequest struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Reason    string
	Status    string
	CreatedAt time.Time
}

// ReturnRequestRepository は return_requests テーブルへのアクセスを抽象化する。
type ReturnRequestRepository interface {
	Create(ctx context.Context, rr *ReturnRequest) error
	Get(ctx context.Context, id uuid.UUID) (*ReturnRequest, error)
	List(ctx context.Context) ([]*ReturnRequest, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*ReturnRequest, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
