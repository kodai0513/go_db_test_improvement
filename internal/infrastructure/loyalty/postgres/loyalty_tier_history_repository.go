package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// LoyaltyTierHistoryRepository は loyalty.LoyaltyTierHistoryRepository の実装。
type LoyaltyTierHistoryRepository struct {
	q *sqlc.Queries
}

func NewLoyaltyTierHistoryRepository(db *sql.DB) *LoyaltyTierHistoryRepository {
	return &LoyaltyTierHistoryRepository{q: sqlc.New(db)}
}

var _ loyalty.LoyaltyTierHistoryRepository = (*LoyaltyTierHistoryRepository)(nil)

func (r *LoyaltyTierHistoryRepository) Create(ctx context.Context, h *loyalty.LoyaltyTierHistory) error {
	row, err := r.q.CreateLoyaltyTierHistory(ctx, sqlc.CreateLoyaltyTierHistoryParams{
		ID:            h.ID,
		MemberID:      h.MemberID,
		LoyaltyTierID: h.LoyaltyTierID,
		ChangedAt:     h.ChangedAt,
	})
	if err != nil {
		return err
	}
	*h = toLoyaltyTierHistoryDomain(row)
	return nil
}

func (r *LoyaltyTierHistoryRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.LoyaltyTierHistory, error) {
	row, err := r.q.GetLoyaltyTierHistory(ctx, id)
	if err != nil {
		return nil, err
	}
	h := toLoyaltyTierHistoryDomain(row)
	return &h, nil
}

func (r *LoyaltyTierHistoryRepository) List(ctx context.Context) ([]*loyalty.LoyaltyTierHistory, error) {
	rows, err := r.q.ListLoyaltyTierHistories(ctx)
	if err != nil {
		return nil, err
	}
	histories := make([]*loyalty.LoyaltyTierHistory, 0, len(rows))
	for _, row := range rows {
		h := toLoyaltyTierHistoryDomain(row)
		histories = append(histories, &h)
	}
	return histories, nil
}

func (r *LoyaltyTierHistoryRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*loyalty.LoyaltyTierHistory, error) {
	rows, err := r.q.ListLoyaltyTierHistoriesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	histories := make([]*loyalty.LoyaltyTierHistory, 0, len(rows))
	for _, row := range rows {
		h := toLoyaltyTierHistoryDomain(row)
		histories = append(histories, &h)
	}
	return histories, nil
}

func (r *LoyaltyTierHistoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteLoyaltyTierHistory(ctx, id)
}

func toLoyaltyTierHistoryDomain(row sqlc.LoyaltyTierHistory) loyalty.LoyaltyTierHistory {
	return loyalty.LoyaltyTierHistory{
		ID:            row.ID,
		MemberID:      row.MemberID,
		LoyaltyTierID: row.LoyaltyTierID,
		ChangedAt:     row.ChangedAt,
	}
}
