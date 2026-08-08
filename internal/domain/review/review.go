package review

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	MemberID  uuid.UUID
	Rating    int32
	Comment   string
	CreatedAt time.Time
}
