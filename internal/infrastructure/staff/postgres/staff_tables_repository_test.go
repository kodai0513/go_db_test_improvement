package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"example.com/go-db-test-improvement/internal/domain/staff"
	"example.com/go-db-test-improvement/internal/infrastructure/staff/postgres"
)

func TestRoleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewRoleRepository(testDB)
	ctx := context.Background()

	role := &staff.Role{
		ID:   uuid.New(),
		Name: "マネージャー",
	}
	require.NoError(t, repo.Create(ctx, role))

	got, err := repo.Get(ctx, role.ID)
	require.NoError(t, err)
	assert.Equal(t, role.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, role.ID))
	_, err = repo.Get(ctx, role.ID)
	assert.Error(t, err)
}

func TestPermissionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewPermissionRepository(testDB)
	ctx := context.Background()

	p := &staff.Permission{
		ID:   uuid.New(),
		Name: "orders:write",
	}
	require.NoError(t, repo.Create(ctx, p))

	got, err := repo.Get(ctx, p.ID)
	require.NoError(t, err)
	assert.Equal(t, p.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, p.ID))
	_, err = repo.Get(ctx, p.ID)
	assert.Error(t, err)
}

func TestRolePermissionRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewRolePermissionRepository(testDB)
	ctx := context.Background()

	rp := &staff.RolePermission{
		ID:           uuid.New(),
		RoleID:       uuid.New(),
		PermissionID: uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, rp))

	got, err := repo.Get(ctx, rp.ID)
	require.NoError(t, err)
	assert.Equal(t, rp.PermissionID, got.PermissionID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byRole, err := repo.ListByRoleID(ctx, rp.RoleID)
	require.NoError(t, err)
	assert.Len(t, byRole, 1)

	require.NoError(t, repo.Delete(ctx, rp.ID))
	_, err = repo.Get(ctx, rp.ID)
	assert.Error(t, err)
}

func TestStaffRoleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStaffRoleRepository(testDB)
	ctx := context.Background()

	sr := &staff.StaffRole{
		ID:            uuid.New(),
		StaffMemberID: uuid.New(),
		RoleID:        uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, sr))

	got, err := repo.Get(ctx, sr.ID)
	require.NoError(t, err)
	assert.Equal(t, sr.RoleID, got.RoleID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByStaffMemberID(ctx, sr.StaffMemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, sr.ID))
	_, err = repo.Get(ctx, sr.ID)
	assert.Error(t, err)
}

func TestStaffActivityLogRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStaffActivityLogRepository(testDB)
	ctx := context.Background()

	l := &staff.StaffActivityLog{
		ID:            uuid.New(),
		StaffMemberID: uuid.New(),
		Action:        "login",
		OccurredAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, l))

	got, err := repo.Get(ctx, l.ID)
	require.NoError(t, err)
	assert.Equal(t, l.Action, got.Action)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByStaffMemberID(ctx, l.StaffMemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, l.ID))
	_, err = repo.Get(ctx, l.ID)
	assert.Error(t, err)
}

func TestDepartmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewDepartmentRepository(testDB)
	ctx := context.Background()

	d := &staff.Department{
		ID:   uuid.New(),
		Name: "カスタマーサポート部",
	}
	require.NoError(t, repo.Create(ctx, d))

	got, err := repo.Get(ctx, d.ID)
	require.NoError(t, err)
	assert.Equal(t, d.Name, got.Name)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	require.NoError(t, repo.Delete(ctx, d.ID))
	_, err = repo.Get(ctx, d.ID)
	assert.Error(t, err)
}

func TestStaffDepartmentRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStaffDepartmentRepository(testDB)
	ctx := context.Background()

	sd := &staff.StaffDepartment{
		ID:            uuid.New(),
		StaffMemberID: uuid.New(),
		DepartmentID:  uuid.New(),
	}
	require.NoError(t, repo.Create(ctx, sd))

	got, err := repo.Get(ctx, sd.ID)
	require.NoError(t, err)
	assert.Equal(t, sd.DepartmentID, got.DepartmentID)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByStaffMemberID(ctx, sd.StaffMemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, sd.ID))
	_, err = repo.Get(ctx, sd.ID)
	assert.Error(t, err)
}

func TestStaffScheduleRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStaffScheduleRepository(testDB)
	ctx := context.Background()

	s := &staff.StaffSchedule{
		ID:            uuid.New(),
		StaffMemberID: uuid.New(),
		WorkDate:      time.Now().UTC().Truncate(24 * time.Hour),
		Shift:         "早番",
	}
	require.NoError(t, repo.Create(ctx, s))

	got, err := repo.Get(ctx, s.ID)
	require.NoError(t, err)
	assert.Equal(t, s.Shift, got.Shift)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByStaffMemberID(ctx, s.StaffMemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, s.ID))
	_, err = repo.Get(ctx, s.ID)
	assert.Error(t, err)
}

func TestStaffPerformanceReviewRepository_CreateGetListDelete(t *testing.T) {
	repo := postgres.NewStaffPerformanceReviewRepository(testDB)
	ctx := context.Background()

	rv := &staff.StaffPerformanceReview{
		ID:            uuid.New(),
		StaffMemberID: uuid.New(),
		Score:         85,
		ReviewedAt:    time.Now().UTC().Truncate(time.Microsecond),
	}
	require.NoError(t, repo.Create(ctx, rv))

	got, err := repo.Get(ctx, rv.ID)
	require.NoError(t, err)
	assert.Equal(t, rv.Score, got.Score)

	list, err := repo.List(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 1)

	byMember, err := repo.ListByStaffMemberID(ctx, rv.StaffMemberID)
	require.NoError(t, err)
	assert.Len(t, byMember, 1)

	require.NoError(t, repo.Delete(ctx, rv.ID))
	_, err = repo.Get(ctx, rv.ID)
	assert.Error(t, err)
}
