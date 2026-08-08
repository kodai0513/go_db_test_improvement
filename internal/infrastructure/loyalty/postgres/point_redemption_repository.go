package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// PointRedemptionRepository は loyalty.PointRedemptionRepository の実装。
type PointRedemptionRepository struct {
	q *sqlc.Queries
}

func NewPointRedemptionRepository(db *sql.DB) *PointRedemptionRepository {
	return &PointRedemptionRepository{q: sqlc.New(db)}
}

var _ loyalty.PointRedemptionRepository = (*PointRedemptionRepository)(nil)

func (r *PointRedemptionRepository) Create(ctx context.Context, p *loyalty.PointRedemption) error {
	row, err := r.q.CreatePointRedemption(ctx, sqlc.CreatePointRedemptionParams{
		ID:         p.ID,
		MemberID:   p.MemberID,
		Points:     p.Points,
		RedeemedAt: p.RedeemedAt,
	})
	if err != nil {
		return err
	}
	*p = toPointRedemptionDomain(row)
	return nil
}

func (r *PointRedemptionRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.PointRedemption, error) {
	row, err := r.q.GetPointRedemption(ctx, id)
	if err != nil {
		return nil, err
	}
	p := toPointRedemptionDomain(row)
	return &p, nil
}

func (r *PointRedemptionRepository) List(ctx context.Context) ([]*loyalty.PointRedemption, error) {
	rows, err := r.q.ListPointRedemptions(ctx)
	if err != nil {
		return nil, err
	}
	redemptions := make([]*loyalty.PointRedemption, 0, len(rows))
	for _, row := range rows {
		p := toPointRedemptionDomain(row)
		redemptions = append(redemptions, &p)
	}
	return redemptions, nil
}

func (r *PointRedemptionRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*loyalty.PointRedemption, error) {
	rows, err := r.q.ListPointRedemptionsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	redemptions := make([]*loyalty.PointRedemption, 0, len(rows))
	for _, row := range rows {
		p := toPointRedemptionDomain(row)
		redemptions = append(redemptions, &p)
	}
	return redemptions, nil
}

func (r *PointRedemptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePointRedemption(ctx, id)
}

func toPointRedemptionDomain(row sqlc.PointRedemption) loyalty.PointRedemption {
	return loyalty.PointRedemption{
		ID:         row.ID,
		MemberID:   row.MemberID,
		Points:     row.Points,
		RedeemedAt: row.RedeemedAt,
	}
}
