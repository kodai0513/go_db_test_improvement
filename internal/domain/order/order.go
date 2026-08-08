package order

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID              uuid.UUID
	MemberID        uuid.UUID
	TotalAmountYen  int32
	Status          string
	CreatedAt       time.Time
}
