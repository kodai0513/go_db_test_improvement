package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// BirthdayRewardRepository は loyalty.BirthdayRewardRepository の実装。
type BirthdayRewardRepository struct {
	q *sqlc.Queries
}

func NewBirthdayRewardRepository(db *sql.DB) *BirthdayRewardRepository {
	return &BirthdayRewardRepository{q: sqlc.New(db)}
}

var _ loyalty.BirthdayRewardRepository = (*BirthdayRewardRepository)(nil)

func (r *BirthdayRewardRepository) Create(ctx context.Context, b *loyalty.BirthdayReward) error {
	row, err := r.q.CreateBirthdayReward(ctx, sqlc.CreateBirthdayRewardParams{
		ID:        b.ID,
		MemberID:  b.MemberID,
		Points:    b.Points,
		GrantedAt: b.GrantedAt,
	})
	if err != nil {
		return err
	}
	*b = toBirthdayRewardDomain(row)
	return nil
}

func (r *BirthdayRewardRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.BirthdayReward, error) {
	row, err := r.q.GetBirthdayReward(ctx, id)
	if err != nil {
		return nil, err
	}
	b := toBirthdayRewardDomain(row)
	return &b, nil
}

func (r *BirthdayRewardRepository) List(ctx context.Context) ([]*loyalty.BirthdayReward, error) {
	rows, err := r.q.ListBirthdayRewards(ctx)
	if err != nil {
		return nil, err
	}
	rewards := make([]*loyalty.BirthdayReward, 0, len(rows))
	for _, row := range rows {
		b := toBirthdayRewardDomain(row)
		rewards = append(rewards, &b)
	}
	return rewards, nil
}

func (r *BirthdayRewardRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*loyalty.BirthdayReward, error) {
	rows, err := r.q.ListBirthdayRewardsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	rewards := make([]*loyalty.BirthdayReward, 0, len(rows))
	for _, row := range rows {
		b := toBirthdayRewardDomain(row)
		rewards = append(rewards, &b)
	}
	return rewards, nil
}

func (r *BirthdayRewardRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteBirthdayReward(ctx, id)
}

func toBirthdayRewardDomain(row sqlc.BirthdayReward) loyalty.BirthdayReward {
	return loyalty.BirthdayReward{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Points:    row.Points,
		GrantedAt: row.GrantedAt,
	}
}
