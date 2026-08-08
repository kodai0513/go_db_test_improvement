package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// RewardItemRepository は loyalty.RewardItemRepository の実装。
type RewardItemRepository struct {
	q *sqlc.Queries
}

func NewRewardItemRepository(db *sql.DB) *RewardItemRepository {
	return &RewardItemRepository{q: sqlc.New(db)}
}

var _ loyalty.RewardItemRepository = (*RewardItemRepository)(nil)

func (r *RewardItemRepository) Create(ctx context.Context, i *loyalty.RewardItem) error {
	row, err := r.q.CreateRewardItem(ctx, sqlc.CreateRewardItemParams{
		ID:         i.ID,
		Name:       i.Name,
		PointsCost: i.PointsCost,
	})
	if err != nil {
		return err
	}
	*i = toRewardItemDomain(row)
	return nil
}

func (r *RewardItemRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.RewardItem, error) {
	row, err := r.q.GetRewardItem(ctx, id)
	if err != nil {
		return nil, err
	}
	i := toRewardItemDomain(row)
	return &i, nil
}

func (r *RewardItemRepository) List(ctx context.Context) ([]*loyalty.RewardItem, error) {
	rows, err := r.q.ListRewardItems(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*loyalty.RewardItem, 0, len(rows))
	for _, row := range rows {
		i := toRewardItemDomain(row)
		items = append(items, &i)
	}
	return items, nil
}

func (r *RewardItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteRewardItem(ctx, id)
}

func toRewardItemDomain(row sqlc.RewardItem) loyalty.RewardItem {
	return loyalty.RewardItem{
		ID:         row.ID,
		Name:       row.Name,
		PointsCost: row.PointsCost,
	}
}
