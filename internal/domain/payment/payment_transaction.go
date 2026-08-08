package payment

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// PaymentTransaction は決済処理の実行履歴を表す。
type PaymentTransaction struct {
	ID          uuid.UUID
	PaymentID   uuid.UUID
	Status      string
	ProcessedAt time.Time
}

// PaymentTransactionRepository は payment_transactions テーブルへのアクセスを抽象化する。
type PaymentTransactionRepository interface {
	Create(ctx context.Context, t *PaymentTransaction) error
	Get(ctx context.Context, id uuid.UUID) (*PaymentTransaction, error)
	List(ctx context.Context) ([]*PaymentTransaction, error)
	ListByPaymentID(ctx context.Context, paymentID uuid.UUID) ([]*PaymentTransaction, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
