package product

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID
	Name      string
	PriceYen  int32
	CreatedAt time.Time
}
