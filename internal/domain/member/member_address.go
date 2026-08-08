package member

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// MemberAddress は会員の住所を表す。
type MemberAddress struct {
	ID         uuid.UUID
	MemberID   uuid.UUID
	PostalCode string
	Prefecture string
	City       string
	Line1      string
	Line2      string
	IsDefault  bool
	CreatedAt  time.Time
}

// MemberAddressRepository は member_addresses テーブルへのアクセスを抽象化する。
type MemberAddressRepository interface {
	Create(ctx context.Context, a *MemberAddress) error
	Get(ctx context.Context, id uuid.UUID) (*MemberAddress, error)
	List(ctx context.Context) ([]*MemberAddress, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*MemberAddress, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
