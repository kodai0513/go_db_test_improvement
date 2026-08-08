package inventory

import (
	"time"

	"github.com/google/uuid"
)

type Inventory struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	Quantity  int32
	UpdatedAt time.Time
}
