-- name: CreateInventory :one
INSERT INTO inventories (id, product_id, quantity, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetInventory :one
SELECT * FROM inventories WHERE id = $1;

-- name: ListInventories :many
SELECT * FROM inventories ORDER BY updated_at DESC;

-- name: CreateWarehouse :one
INSERT INTO warehouses (id, name, address, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetWarehouse :one
SELECT * FROM warehouses WHERE id = $1;

-- name: ListWarehouses :many
SELECT * FROM warehouses ORDER BY created_at DESC;

-- name: DeleteWarehouse :exec
DELETE FROM warehouses WHERE id = $1;

-- name: CreateStockMovement :one
INSERT INTO stock_movements (id, inventory_id, quantity_delta, reason, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetStockMovement :one
SELECT * FROM stock_movements WHERE id = $1;

-- name: ListStockMovements :many
SELECT * FROM stock_movements ORDER BY created_at DESC;

-- name: ListStockMovementsByInventoryID :many
SELECT * FROM stock_movements WHERE inventory_id = $1 ORDER BY created_at DESC;

-- name: DeleteStockMovement :exec
DELETE FROM stock_movements WHERE id = $1;

-- name: CreateStockReservation :one
INSERT INTO stock_reservations (id, inventory_id, order_id, quantity, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetStockReservation :one
SELECT * FROM stock_reservations WHERE id = $1;

-- name: ListStockReservations :many
SELECT * FROM stock_reservations ORDER BY created_at DESC;

-- name: ListStockReservationsByInventoryID :many
SELECT * FROM stock_reservations WHERE inventory_id = $1 ORDER BY created_at DESC;

-- name: ListStockReservationsByOrderID :many
SELECT * FROM stock_reservations WHERE order_id = $1 ORDER BY created_at DESC;

-- name: DeleteStockReservation :exec
DELETE FROM stock_reservations WHERE id = $1;

-- name: CreateSupplier :one
INSERT INTO suppliers (id, name, contact_email, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSupplier :one
SELECT * FROM suppliers WHERE id = $1;

-- name: ListSuppliers :many
SELECT * FROM suppliers ORDER BY created_at DESC;

-- name: DeleteSupplier :exec
DELETE FROM suppliers WHERE id = $1;

-- name: CreatePurchaseOrder :one
INSERT INTO purchase_orders (id, supplier_id, status, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPurchaseOrder :one
SELECT * FROM purchase_orders WHERE id = $1;

-- name: ListPurchaseOrders :many
SELECT * FROM purchase_orders ORDER BY created_at DESC;

-- name: ListPurchaseOrdersBySupplierID :many
SELECT * FROM purchase_orders WHERE supplier_id = $1 ORDER BY created_at DESC;

-- name: DeletePurchaseOrder :exec
DELETE FROM purchase_orders WHERE id = $1;

-- name: CreatePurchaseOrderItem :one
INSERT INTO purchase_order_items (id, purchase_order_id, product_id, quantity, unit_price_yen)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPurchaseOrderItem :one
SELECT * FROM purchase_order_items WHERE id = $1;

-- name: ListPurchaseOrderItems :many
SELECT * FROM purchase_order_items ORDER BY id DESC;

-- name: ListPurchaseOrderItemsByPurchaseOrderID :many
SELECT * FROM purchase_order_items WHERE purchase_order_id = $1 ORDER BY id DESC;

-- name: ListPurchaseOrderItemsByProductID :many
SELECT * FROM purchase_order_items WHERE product_id = $1 ORDER BY id DESC;

-- name: DeletePurchaseOrderItem :exec
DELETE FROM purchase_order_items WHERE id = $1;

-- name: CreateSupplierProduct :one
INSERT INTO supplier_products (id, supplier_id, product_id, lead_time_days)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetSupplierProduct :one
SELECT * FROM supplier_products WHERE id = $1;

-- name: ListSupplierProducts :many
SELECT * FROM supplier_products ORDER BY id DESC;

-- name: ListSupplierProductsBySupplierID :many
SELECT * FROM supplier_products WHERE supplier_id = $1 ORDER BY id DESC;

-- name: ListSupplierProductsByProductID :many
SELECT * FROM supplier_products WHERE product_id = $1 ORDER BY id DESC;

-- name: DeleteSupplierProduct :exec
DELETE FROM supplier_products WHERE id = $1;

-- name: CreateStockAdjustment :one
INSERT INTO stock_adjustments (id, inventory_id, adjusted_quantity, note, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetStockAdjustment :one
SELECT * FROM stock_adjustments WHERE id = $1;

-- name: ListStockAdjustments :many
SELECT * FROM stock_adjustments ORDER BY created_at DESC;

-- name: ListStockAdjustmentsByInventoryID :many
SELECT * FROM stock_adjustments WHERE inventory_id = $1 ORDER BY created_at DESC;

-- name: DeleteStockAdjustment :exec
DELETE FROM stock_adjustments WHERE id = $1;

-- name: CreateInventoryAudit :one
INSERT INTO inventory_audits (id, warehouse_id, performed_at, result)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetInventoryAudit :one
SELECT * FROM inventory_audits WHERE id = $1;

-- name: ListInventoryAudits :many
SELECT * FROM inventory_audits ORDER BY performed_at DESC;

-- name: ListInventoryAuditsByWarehouseID :many
SELECT * FROM inventory_audits WHERE warehouse_id = $1 ORDER BY performed_at DESC;

-- name: DeleteInventoryAudit :exec
DELETE FROM inventory_audits WHERE id = $1;
