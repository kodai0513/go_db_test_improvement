package member

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, m *Member) error
	Get(ctx context.Context, id uuid.UUID) (*Member, error)
	List(ctx context.Context) ([]*Member, error)

	// ListMemberCountAndAvgPointBalanceByTag は、タグごとの会員数と平均ポイント残高を集計する。
	ListMemberCountAndAvgPointBalanceByTag(ctx context.Context) ([]*MemberTagPointSummary, error)
	// ListTopMembersByAddressCountAndPointBalance は、ポイント残高・住所登録数の降順で上位limit件の会員を集計する。
	ListTopMembersByAddressCountAndPointBalance(ctx context.Context, limit int32) ([]*MemberAddressPointRanking, error)
	// ListReferralRewardSummaryByReferrer は、紹介者ごとの紹介成功件数と獲得報酬合計を集計する。
	ListReferralRewardSummaryByReferrer(ctx context.Context) ([]*ReferralRewardSummary, error)
}

// MemberTagPointSummary はタグ別の会員数・平均ポイント残高の集計結果を表す。
type MemberTagPointSummary struct {
	MemberTagID     uuid.UUID
	TagName         string
	MemberCount     int64
	AvgPointBalance float64
}

// MemberAddressPointRanking は会員別の住所登録数・ポイント残高の集計結果を表す。
type MemberAddressPointRanking struct {
	MemberID     uuid.UUID
	MemberName   string
	AddressCount int64
	PointBalance int32
}

// ReferralRewardSummary は紹介者別の紹介成功件数・獲得報酬合計の集計結果を表す。
type ReferralRewardSummary struct {
	ReferrerMemberID uuid.UUID
	ReferrerName     string
	ReferralCount    int64
	TotalRewardYen   int64
}
