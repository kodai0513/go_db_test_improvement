package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// RewardRedemptionRepository は loyalty.RewardRedemptionRepository の実装。
type RewardRedemptionRepository struct {
	q *sqlc.Queries
}

func NewRewardRedemptionRepository(db *sql.DB) *RewardRedemptionRepository {
	return &RewardRedemptionRepository{q: sqlc.New(db)}
}

var _ loyalty.RewardRedemptionRepository = (*RewardRedemptionRepository)(nil)

func (r *RewardRedemptionRepository) Create(ctx context.Context, rr *loyalty.RewardRedemption) error {
	row, err := r.q.CreateRewardRedemption(ctx, sqlc.CreateRewardRedemptionParams{
		ID:           rr.ID,
		MemberID:     rr.MemberID,
		RewardItemID: rr.RewardItemID,
		RedeemedAt:   rr.RedeemedAt,
	})
	if err != nil {
		return err
	}
	*rr = toRewardRedemptionDomain(row)
	return nil
}

func (r *RewardRedemptionRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.RewardRedemption, error) {
	row, err := r.q.GetRewardRedemption(ctx, id)
	if err != nil {
		return nil, err
	}
	rr := toRewardRedemptionDomain(row)
	return &rr, nil
}

func (r *RewardRedemptionRepository) List(ctx context.Context) ([]*loyalty.RewardRedemption, error) {
	rows, err := r.q.ListRewardRedemptions(ctx)
	if err != nil {
		return nil, err
	}
	redemptions := make([]*loyalty.RewardRedemption, 0, len(rows))
	for _, row := range rows {
		rr := toRewardRedemptionDomain(row)
		redemptions = append(redemptions, &rr)
	}
	return redemptions, nil
}

func (r *RewardRedemptionRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*loyalty.RewardRedemption, error) {
	rows, err := r.q.ListRewardRedemptionsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	redemptions := make([]*loyalty.RewardRedemption, 0, len(rows))
	for _, row := range rows {
		rr := toRewardRedemptionDomain(row)
		redemptions = append(redemptions, &rr)
	}
	return redemptions, nil
}

func (r *RewardRedemptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteRewardRedemption(ctx, id)
}

func toRewardRedemptionDomain(row sqlc.RewardRedemption) loyalty.RewardRedemption {
	return loyalty.RewardRedemption{
		ID:           row.ID,
		MemberID:     row.MemberID,
		RewardItemID: row.RewardItemID,
		RedeemedAt:   row.RedeemedAt,
	}
}
