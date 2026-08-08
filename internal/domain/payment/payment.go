package payment

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	AmountYen int32
	Method    string
	PaidAt    time.Time
}
