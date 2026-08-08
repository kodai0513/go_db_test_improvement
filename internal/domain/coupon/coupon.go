package coupon

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID           uuid.UUID
	Code         string
	DiscountRate int32
	ExpiresAt    time.Time
}
