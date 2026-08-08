-- name: CreateStaffMember :one
INSERT INTO staff_members (id, name, email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetStaffMember :one
SELECT * FROM staff_members WHERE id = $1;

-- name: ListStaffMembers :many
SELECT * FROM staff_members ORDER BY created_at DESC;

-- name: DeleteStaffMember :exec
DELETE FROM staff_members WHERE id = $1;

-- name: CreateRole :one
INSERT INTO roles (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: ListRoles :many
SELECT * FROM roles ORDER BY id DESC;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;

-- name: CreatePermission :one
INSERT INTO permissions (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetPermission :one
SELECT * FROM permissions WHERE id = $1;

-- name: ListPermissions :many
SELECT * FROM permissions ORDER BY id DESC;

-- name: DeletePermission :exec
DELETE FROM permissions WHERE id = $1;

-- name: CreateRolePermission :one
INSERT INTO role_permissions (id, role_id, permission_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRolePermission :one
SELECT * FROM role_permissions WHERE id = $1;

-- name: ListRolePermissions :many
SELECT * FROM role_permissions ORDER BY id DESC;

-- name: ListRolePermissionsByRoleID :many
SELECT * FROM role_permissions WHERE role_id = $1 ORDER BY id DESC;

-- name: DeleteRolePermission :exec
DELETE FROM role_permissions WHERE id = $1;

-- name: CreateStaffRole :one
INSERT INTO staff_roles (id, staff_member_id, role_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetStaffRole :one
SELECT * FROM staff_roles WHERE id = $1;

-- name: ListStaffRoles :many
SELECT * FROM staff_roles ORDER BY id DESC;

-- name: ListStaffRolesByStaffMemberID :many
SELECT * FROM staff_roles WHERE staff_member_id = $1 ORDER BY id DESC;

-- name: DeleteStaffRole :exec
DELETE FROM staff_roles WHERE id = $1;

-- name: CreateStaffActivityLog :one
INSERT INTO staff_activity_logs (id, staff_member_id, action, occurred_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetStaffActivityLog :one
SELECT * FROM staff_activity_logs WHERE id = $1;

-- name: ListStaffActivityLogs :many
SELECT * FROM staff_activity_logs ORDER BY occurred_at DESC;

-- name: ListStaffActivityLogsByStaffMemberID :many
SELECT * FROM staff_activity_logs WHERE staff_member_id = $1 ORDER BY occurred_at DESC;

-- name: DeleteStaffActivityLog :exec
DELETE FROM staff_activity_logs WHERE id = $1;

-- name: CreateDepartment :one
INSERT INTO departments (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetDepartment :one
SELECT * FROM departments WHERE id = $1;

-- name: ListDepartments :many
SELECT * FROM departments ORDER BY id DESC;

-- name: DeleteDepartment :exec
DELETE FROM departments WHERE id = $1;

-- name: CreateStaffDepartment :one
INSERT INTO staff_departments (id, staff_member_id, department_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetStaffDepartment :one
SELECT * FROM staff_departments WHERE id = $1;

-- name: ListStaffDepartments :many
SELECT * FROM staff_departments ORDER BY id DESC;

-- name: ListStaffDepartmentsByStaffMemberID :many
SELECT * FROM staff_departments WHERE staff_member_id = $1 ORDER BY id DESC;

-- name: DeleteStaffDepartment :exec
DELETE FROM staff_departments WHERE id = $1;

-- name: CreateStaffSchedule :one
INSERT INTO staff_schedules (id, staff_member_id, work_date, shift)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetStaffSchedule :one
SELECT * FROM staff_schedules WHERE id = $1;

-- name: ListStaffSchedules :many
SELECT * FROM staff_schedules ORDER BY work_date DESC;

-- name: ListStaffSchedulesByStaffMemberID :many
SELECT * FROM staff_schedules WHERE staff_member_id = $1 ORDER BY work_date DESC;

-- name: DeleteStaffSchedule :exec
DELETE FROM staff_schedules WHERE id = $1;

-- name: CreateStaffPerformanceReview :one
INSERT INTO staff_performance_reviews (id, staff_member_id, score, reviewed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetStaffPerformanceReview :one
SELECT * FROM staff_performance_reviews WHERE id = $1;

-- name: ListStaffPerformanceReviews :many
SELECT * FROM staff_performance_reviews ORDER BY reviewed_at DESC;

-- name: ListStaffPerformanceReviewsByStaffMemberID :many
SELECT * FROM staff_performance_reviews WHERE staff_member_id = $1 ORDER BY reviewed_at DESC;

-- name: DeleteStaffPerformanceReview :exec
DELETE FROM staff_performance_reviews WHERE id = $1;
