package notification

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID
	MemberID  uuid.UUID
	Message   string
	IsRead    bool
	CreatedAt time.Time
}
