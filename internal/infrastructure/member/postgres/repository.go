package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/member"
	"example.com/go-db-test-improvement/internal/infrastructure/member/postgres/sqlc"
)

// Repository implements member.Repository backed by the sqlc-generated
// query layer.
type Repository struct {
	q *sqlc.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{q: sqlc.New(db)}
}

var _ member.Repository = (*Repository)(nil)

func (r *Repository) Create(ctx context.Context, m *member.Member) error {
	row, err := r.q.CreateMember(ctx, sqlc.CreateMemberParams{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
	})
	if err != nil {
		return err
	}
	*m = toDomain(row)
	return nil
}

func (r *Repository) Get(ctx context.Context, id uuid.UUID) (*member.Member, error) {
	row, err := r.q.GetMember(ctx, id)
	if err != nil {
		return nil, err
	}
	m := toDomain(row)
	return &m, nil
}

func (r *Repository) List(ctx context.Context) ([]*member.Member, error) {
	rows, err := r.q.ListMembers(ctx)
	if err != nil {
		return nil, err
	}
	members := make([]*member.Member, 0, len(rows))
	for _, row := range rows {
		m := toDomain(row)
		members = append(members, &m)
	}
	return members, nil
}

func (r *Repository) ListMemberCountAndAvgPointBalanceByTag(ctx context.Context) ([]*member.MemberTagPointSummary, error) {
	rows, err := r.q.ListMemberCountAndAvgPointBalanceByTag(ctx)
	if err != nil {
		return nil, err
	}
	summaries := make([]*member.MemberTagPointSummary, 0, len(rows))
	for _, row := range rows {
		summaries = append(summaries, &member.MemberTagPointSummary{
			MemberTagID:     row.MemberTagID,
			TagName:         row.TagName,
			MemberCount:     row.MemberCount,
			AvgPointBalance: row.AvgPointBalance,
		})
	}
	return summaries, nil
}

func (r *Repository) ListTopMembersByAddressCountAndPointBalance(ctx context.Context, limit int32) ([]*member.MemberAddressPointRanking, error) {
	rows, err := r.q.ListTopMembersByAddressCountAndPointBalance(ctx, limit)
	if err != nil {
		return nil, err
	}
	rankings := make([]*member.MemberAddressPointRanking, 0, len(rows))
	for _, row := range rows {
		rankings = append(rankings, &member.MemberAddressPointRanking{
			MemberID:     row.MemberID,
			MemberName:   row.MemberName,
			AddressCount: row.AddressCount,
			PointBalance: row.PointBalance,
		})
	}
	return rankings, nil
}

func (r *Repository) ListReferralRewardSummaryByReferrer(ctx context.Context) ([]*member.ReferralRewardSummary, error) {
	rows, err := r.q.ListReferralRewardSummaryByReferrer(ctx)
	if err != nil {
		return nil, err
	}
	summaries := make([]*member.ReferralRewardSummary, 0, len(rows))
	for _, row := range rows {
		summaries = append(summaries, &member.ReferralRewardSummary{
			ReferrerMemberID: row.ReferrerMemberID,
			ReferrerName:     row.ReferrerName,
			ReferralCount:    row.ReferralCount,
			TotalRewardYen:   row.TotalRewardYen,
		})
	}
	return summaries, nil
}

func toDomain(row sqlc.Member) member.Member {
	return member.Member{
		ID:        row.ID,
		Name:      row.Name,
		Email:     row.Email,
		CreatedAt: row.CreatedAt,
	}
}
