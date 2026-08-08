package merchant

import (
	"context"

	"github.com/google/uuid"
)

// VendorBankAccount はベンダーの振込先口座を表す。
type VendorBankAccount struct {
	ID            uuid.UUID
	VendorID      uuid.UUID
	AccountNumber string
}

// VendorBankAccountRepository は vendor_bank_accounts テーブルへのアクセスを抽象化する。
type VendorBankAccountRepository interface {
	Create(ctx context.Context, a *VendorBankAccount) error
	Get(ctx context.Context, id uuid.UUID) (*VendorBankAccount, error)
	List(ctx context.Context) ([]*VendorBankAccount, error)
	ListByVendorID(ctx context.Context, vendorID uuid.UUID) ([]*VendorBankAccount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
