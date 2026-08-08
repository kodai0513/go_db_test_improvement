package member

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
}
