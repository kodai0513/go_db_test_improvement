package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"example.com/go-db-test-improvement/internal/domain/notification"
	"example.com/go-db-test-improvement/internal/infrastructure/notification/postgres/sqlc"
)

// PushDeviceRepository は notification.PushDeviceRepository の実装。
type PushDeviceRepository struct {
	q *sqlc.Queries
}

func NewPushDeviceRepository(db *sql.DB) *PushDeviceRepository {
	return &PushDeviceRepository{q: sqlc.New(db)}
}

var _ notification.PushDeviceRepository = (*PushDeviceRepository)(nil)

func (r *PushDeviceRepository) Create(ctx context.Context, d *notification.PushDevice) error {
	row, err := r.q.CreatePushDevice(ctx, sqlc.CreatePushDeviceParams{
		ID:           d.ID,
		MemberID:     d.MemberID,
		DeviceToken:  d.DeviceToken,
		RegisteredAt: d.RegisteredAt,
	})
	if err != nil {
		return err
	}
	*d = toPushDeviceDomain(row)
	return nil
}

func (r *PushDeviceRepository) Get(ctx context.Context, id uuid.UUID) (*notification.PushDevice, error) {
	row, err := r.q.GetPushDevice(ctx, id)
	if err != nil {
		return nil, err
	}
	d := toPushDeviceDomain(row)
	return &d, nil
}

func (r *PushDeviceRepository) List(ctx context.Context) ([]*notification.PushDevice, error) {
	rows, err := r.q.ListPushDevices(ctx)
	if err != nil {
		return nil, err
	}
	devices := make([]*notification.PushDevice, 0, len(rows))
	for _, row := range rows {
		d := toPushDeviceDomain(row)
		devices = append(devices, &d)
	}
	return devices, nil
}

func (r *PushDeviceRepository) ListByMemberID(ctx context.Context, memberID uuid.UUID) ([]*notification.PushDevice, error) {
	rows, err := r.q.ListPushDevicesByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	devices := make([]*notification.PushDevice, 0, len(rows))
	for _, row := range rows {
		d := toPushDeviceDomain(row)
		devices = append(devices, &d)
	}
	return devices, nil
}

func (r *PushDeviceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.q.DeletePushDevice(ctx, id)
}

func toPushDeviceDomain(row sqlc.PushDevice) notification.PushDevice {
	return notification.PushDevice{
		ID:           row.ID,
		MemberID:     row.MemberID,
		DeviceToken:  row.DeviceToken,
		RegisteredAt: row.RegisteredAt,
	}
}
