package payment

import (
	"context"

	"github.com/google/uuid"
)

// PaymentProviderAccount は会員が決済代行事業者側に保有するアカウントを表す。
type PaymentProviderAccount struct {
	ID                uuid.UUID
	PaymentProviderID uuid.UUID
	MemberID          uuid.UUID
	ExternalAccountID string
}

// PaymentProviderAccountRepository は payment_provider_accounts テーブルへのアクセスを抽象化する。
type PaymentProviderAccountRepository interface {
	Create(ctx context.Context, a *PaymentProviderAccount) error
	Get(ctx context.Context, id uuid.UUID) (*PaymentProviderAccount, error)
	List(ctx context.Context) ([]*PaymentProviderAccount, error)
	ListByPaymentProviderID(ctx context.Context, paymentProviderID uuid.UUID) ([]*PaymentProviderAccount, error)
	ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*PaymentProviderAccount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
