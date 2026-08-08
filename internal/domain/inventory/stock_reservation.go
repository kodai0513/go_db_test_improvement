package inventory

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// StockReservation は在庫の引当を表す。
type StockReservation struct {
	ID          uuid.UUID
	InventoryID uuid.UUID
	OrderID     uuid.UUID
	Quantity    int32
	CreatedAt   time.Time
}

// StockReservationRepository は stock_reservations テーブルへのアクセスを抽象化する。
type StockReservationRepository interface {
	Create(ctx context.Context, r *StockReservation) error
	Get(ctx context.Context, id uuid.UUID) (*StockReservation, error)
	List(ctx context.Context) ([]*StockReservation, error)
	ListByInventoryID(ctx context.Context, inventoryID uuid.UUID) ([]*StockReservation, error)
	ListByOrderID(ctx context.Context, orderID uuid.UUID) ([]*StockReservation, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
