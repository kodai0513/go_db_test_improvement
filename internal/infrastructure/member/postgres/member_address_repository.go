package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberAddressRepository は member.MemberAddressRepository の実装。
type MemberAddressRepository struct {
	q *sqlc.Queries
}

func NewMemberAddressRepository(db *sql.DB) *MemberAddressRepository {
	return &MemberAddressRepository{q: sqlc.New(db)}
}

var _ member.MemberAddressRepository = (*MemberAddressRepository)(nil)

func (r *MemberAddressRepository) Create(ctx context.Context, a *member.MemberAddress) error {
	row, err := r.q.CreateMemberAddress(ctx, sqlc.CreateMemberAddressParams{
		ID:         a.ID,
		MemberID:   a.MemberID,
		PostalCode: a.PostalCode,
		Prefecture: a.Prefecture,
		City:       a.City,
		Line1:      a.Line1,
		Line2:      a.Line2,
		IsDefault:  a.IsDefault,
		CreatedAt:  a.CreatedAt,
	})
	if err != nil {
		return err
	}
	*a = toMemberAddressDomain(row)
	return nil
}

func (r *MemberAddressRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberAddress, error) {
	row, err := r.q.GetMemberAddress(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toMemberAddressDomain(row)
	return &a, nil
}

func (r *MemberAddressRepository) List(ctx context.Context) ([]*member.MemberAddress, error) {
	rows, err := r.q.ListMemberAddresses(ctx)
	if err != nil {
		return nil, err
	}
	addresses := make([]*member.MemberAddress, 0, len(rows))
	for _, row := range rows {
		a := toMemberAddressDomain(row)
		addresses = append(addresses, &a)
	}
	return addresses, nil
}

func (r *MemberAddressRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberAddress, error) {
	rows, err := r.q.ListMemberAddressesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	addresses := make([]*member.MemberAddress, 0, len(rows))
	for _, row := range rows {
		a := toMemberAddressDomain(row)
		addresses = append(addresses, &a)
	}
	return addresses, nil
}

func (r *MemberAddressRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberAddress(ctx, id)
}

func toMemberAddressDomain(row sqlc.MemberAddress) member.MemberAddress {
	return member.MemberAddress{
		ID:         row.ID,
		MemberID:   row.MemberID,
		PostalCode: row.PostalCode,
		Prefecture: row.Prefecture,
		City:       row.City,
		Line1:      row.Line1,
		Line2:      row.Line2,
		IsDefault:  row.IsDefault,
		CreatedAt:  row.CreatedAt,
	}
}
