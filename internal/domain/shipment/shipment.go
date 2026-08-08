package shipment

import (
	"database/sql"

	"github.com/google/uuid"
)

type Shipment struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Address   string
	Status    string
	ShippedAt sql.NullTime
}
