package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberPointAccountRepository は member.MemberPointAccountRepository の実装。
type MemberPointAccountRepository struct {
	q *sqlc.Queries
}

func NewMemberPointAccountRepository(db *sql.DB) *MemberPointAccountRepository {
	return &MemberPointAccountRepository{q: sqlc.New(db)}
}

var _ member.MemberPointAccountRepository = (*MemberPointAccountRepository)(nil)

func (r *MemberPointAccountRepository) Create(ctx context.Context, a *member.MemberPointAccount) error {
	row, err := r.q.CreateMemberPointAccount(ctx, sqlc.CreateMemberPointAccountParams{
		ID:        a.ID,
		MemberID:  a.MemberID,
		Balance:   a.Balance,
		UpdatedAt: a.UpdatedAt,
	})
	if err != nil {
		return err
	}
	*a = toMemberPointAccountDomain(row)
	return nil
}

func (r *MemberPointAccountRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberPointAccount, error) {
	row, err := r.q.GetMemberPointAccount(ctx, id)
	if err != nil {
		return nil, err
	}
	a := toMemberPointAccountDomain(row)
	return &a, nil
}

func (r *MemberPointAccountRepository) List(ctx context.Context) ([]*member.MemberPointAccount, error) {
	rows, err := r.q.ListMemberPointAccounts(ctx)
	if err != nil {
		return nil, err
	}
	accounts := make([]*member.MemberPointAccount, 0, len(rows))
	for _, row := range rows {
		a := toMemberPointAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *MemberPointAccountRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*member.MemberPointAccount, error) {
	rows, err := r.q.ListMemberPointAccountsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	accounts := make([]*member.MemberPointAccount, 0, len(rows))
	for _, row := range rows {
		a := toMemberPointAccountDomain(row)
		accounts = append(accounts, &a)
	}
	return accounts, nil
}

func (r *MemberPointAccountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberPointAccount(ctx, id)
}

func toMemberPointAccountDomain(row sqlc.MemberPointAccount) member.MemberPointAccount {
	return member.MemberPointAccount{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Balance:   row.Balance,
		UpdatedAt: row.UpdatedAt,
	}
}
