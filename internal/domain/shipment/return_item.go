package shipment

import (
	"context"

	"github.com/google/uuid"
)

// ReturnItem は返品依頼に含まれる商品明細を表す。
type ReturnItem struct {
	ID              uuid.UUID
	ReturnRequestID uuid.UUID
	OrderItemID     uuid.UUID
	Quantity        int32
}

// ReturnItemRepository は return_items テーブルへのアクセスを抽象化する。
type ReturnItemRepository interface {
	Create(ctx context.Context, ri *ReturnItem) error
	Get(ctx context.Context, id uuid.UUID) (*ReturnItem, error)
	List(ctx context.Context) ([]*ReturnItem, error)
	ListByReturnRequestID(ctx context.Context, returnRequestID uuid.UUID) ([]*ReturnItem, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
