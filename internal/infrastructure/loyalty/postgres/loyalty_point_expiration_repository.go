package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/loyalty"
	"example.com/go-db-test-improvement/internal/infrastructure/loyalty/postgres/sqlc"
)

// LoyaltyPointExpirationRepository は loyalty.LoyaltyPointExpirationRepository の実装。
type LoyaltyPointExpirationRepository struct {
	q *sqlc.Queries
}

func NewLoyaltyPointExpirationRepository(db *sql.DB) *LoyaltyPointExpirationRepository {
	return &LoyaltyPointExpirationRepository{q: sqlc.New(db)}
}

var _ loyalty.LoyaltyPointExpirationRepository = (*LoyaltyPointExpirationRepository)(nil)

func (r *LoyaltyPointExpirationRepository) Create(ctx context.Context, e *loyalty.LoyaltyPointExpiration) error {
	row, err := r.q.CreateLoyaltyPointExpiration(ctx, sqlc.CreateLoyaltyPointExpirationParams{
		ID:        e.ID,
		MemberID:  e.MemberID,
		Points:    e.Points,
		ExpiresAt: e.ExpiresAt,
	})
	if err != nil {
		return err
	}
	*e = toLoyaltyPointExpirationDomain(row)
	return nil
}

func (r *LoyaltyPointExpirationRepository) Get(ctx context.Context, id uuid.UUID) (*loyalty.LoyaltyPointExpiration, error) {
	row, err := r.q.GetLoyaltyPointExpiration(ctx, id)
	if err != nil {
		return nil, err
	}
	e := toLoyaltyPointExpirationDomain(row)
	return &e, nil
}

func (r *LoyaltyPointExpirationRepository) List(ctx context.Context) ([]*loyalty.LoyaltyPointExpiration, error) {
	rows, err := r.q.ListLoyaltyPointExpirations(ctx)
	if err != nil {
		return nil, err
	}
	expirations := make([]*loyalty.LoyaltyPointExpiration, 0, len(rows))
	for _, row := range rows {
		e := toLoyaltyPointExpirationDomain(row)
		expirations = append(expirations, &e)
	}
	return expirations, nil
}

func (r *LoyaltyPointExpirationRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*loyalty.LoyaltyPointExpiration, error) {
	rows, err := r.q.ListLoyaltyPointExpirationsByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	expirations := make([]*loyalty.LoyaltyPointExpiration, 0, len(rows))
	for _, row := range rows {
		e := toLoyaltyPointExpirationDomain(row)
		expirations = append(expirations, &e)
	}
	return expirations, nil
}

func (r *LoyaltyPointExpirationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteLoyaltyPointExpiration(ctx, id)
}

func toLoyaltyPointExpirationDomain(row sqlc.LoyaltyPointExpiration) loyalty.LoyaltyPointExpiration {
	return loyalty.LoyaltyPointExpiration{
		ID:        row.ID,
		MemberID:  row.MemberID,
		Points:    row.Points,
		ExpiresAt: row.ExpiresAt,
	}
}
