package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// MemberReferralRepository は member.MemberReferralRepository の実装。
type MemberReferralRepository struct {
	q *sqlc.Queries
}

func NewMemberReferralRepository(db *sql.DB) *MemberReferralRepository {
	return &MemberReferralRepository{q: sqlc.New(db)}
}

var _ member.MemberReferralRepository = (*MemberReferralRepository)(nil)

func (r *MemberReferralRepository) Create(ctx context.Context, ref *member.MemberReferral) error {
	row, err := r.q.CreateMemberReferral(ctx, sqlc.CreateMemberReferralParams{
		ID:               ref.ID,
		ReferrerMemberID: ref.ReferrerMemberID,
		ReferredMemberID: ref.ReferredMemberID,
		RewardYen:        ref.RewardYen,
		CreatedAt:        ref.CreatedAt,
	})
	if err != nil {
		return err
	}
	*ref = toMemberReferralDomain(row)
	return nil
}

func (r *MemberReferralRepository) Get(ctx context.Context, id uuid.UUID) (*member.MemberReferral, error) {
	row, err := r.q.GetMemberReferral(ctx, id)
	if err != nil {
		return nil, err
	}
	ref := toMemberReferralDomain(row)
	return &ref, nil
}

func (r *MemberReferralRepository) List(ctx context.Context) ([]*member.MemberReferral, error) {
	rows, err := r.q.ListMemberReferrals(ctx)
	if err != nil {
		return nil, err
	}
	referrals := make([]*member.MemberReferral, 0, len(rows))
	for _, row := range rows {
		ref := toMemberReferralDomain(row)
		referrals = append(referrals, &ref)
	}
	return referrals, nil
}

func (r *MemberReferralRepository) ListByReferrerMemberID(ctx context.Context, referrerMemberID uuid.UUID) ([]*member.MemberReferral, error) {
	rows, err := r.q.ListMemberReferralsByReferrerMemberID(ctx, referrerMemberID)
	if err != nil {
		return nil, err
	}
	referrals := make([]*member.MemberReferral, 0, len(rows))
	for _, row := range rows {
		ref := toMemberReferralDomain(row)
		referrals = append(referrals, &ref)
	}
	return referrals, nil
}

func (r *MemberReferralRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteMemberReferral(ctx, id)
}

func toMemberReferralDomain(row sqlc.MemberReferral) member.MemberReferral {
	return member.MemberReferral{
		ID:               row.ID,
		ReferrerMemberID: row.ReferrerMemberID,
		ReferredMemberID: row.ReferredMemberID,
		RewardYen:        row.RewardYen,
		CreatedAt:        row.CreatedAt,
	}
}
