package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// ReferralBonusRepository は loyalty.ReferralBonusRepository の実装。
type ReferralBonusRepository struct {
	q *sqlc.Queries
}

func NewReferralBonusRepository(db *sql.DB) *ReferralBonusRepository {
	return &ReferralBonusRepository{q: sqlc.New(db)}
}

var _ loyalty.ReferralBonusRepository = (*ReferralBonusRepository)(nil)

func (r *ReferralBonusRepository) Create(ctx context.Context, b *loyalty.ReferralBonus) error {
	row, err := r.q.CreateReferralBonus(ctx, sqlc.CreateReferralBonusParams{
		ID:        b.ID,
		MemberID:  b.MemberID,
		Points:    b.Points,
		GrantedAt: b.GrantedAt,
	})
	if err != nil {
		return err
	}
	*b = toReferralBonusDomain(row)
	return nil
}

func (r *ReferralBonusRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.ReferralBonus, error) {
	row, err := r.q.GetReferralBonus(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toReferralBonusDomain(row)
	return &b, nil
}

func (r *ReferralBonusRepository) List(ctx context.Context) ([]*loyalty.ReferralBonus, error) {
	rows, err := r.q.ListReferralBonuses(ctx)
	if err != nil {
		return nil, err
	}
	bonuses := make([]*loyalty.ReferralBonus, 0, len(rows))
	for _, row := range rows {
		b := toReferralBonusDomain(row)
		bonuses = append(bonuses, &b)
	}
	return bonuses, nil
}

func (r *ReferralBonusRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*loyalty.ReferralBonus, error) {
	rows, err := r.q.ListReferralBonusesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	bonuses := make([]*loyalty.ReferralBonus, 0, len(rows))
	for _, row := range rows {
		b := toReferralBonusDomain(row)
		bonuses = append(bonuses, &b)
	}
	return bonuses, nil
}

func (r *ReferralBonusRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteReferralBonus(ctx, id)
}

func toReferralBonusDomain(row sqlc.ReferralBonus) loyalty.ReferralBonus {
	return loyalty.ReferralBonus{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Points:    row.Points,
		GrantedAt: row.GrantedAt,
	}
}
