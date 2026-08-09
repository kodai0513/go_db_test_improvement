-- name: CreateOrder :one
INSERT INTO orders (id, member_id, total_amount_yen, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY created_at DESC;

-- name: CreateOrderItem :one
INSERT INTO order_items (id, order_id, product_id, quantity, unit_price_yen)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetOrderItem :one
SELECT * FROM order_items WHERE id = $1;

-- name: ListOrderItems :many
SELECT * FROM order_items ORDER BY id DESC;

-- name: ListOrderItemsByOrderID :many
SELECT * FROM order_items WHERE order_id = $1 ORDER BY id DESC;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE id = $1;

-- name: CreateOrderStatusHistory :one
INSERT INTO order_status_histories (id, order_id, status, changed_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderStatusHistory :one
SELECT * FROM order_status_histories WHERE id = $1;

-- name: ListOrderStatusHistories :many
SELECT * FROM order_status_histories ORDER BY changed_at DESC;

-- name: ListOrderStatusHistoriesByOrderID :many
SELECT * FROM order_status_histories WHERE order_id = $1 ORDER BY changed_at DESC;

-- name: DeleteOrderStatusHistory :exec
DELETE FROM order_status_histories WHERE id = $1;

-- name: CreateOrderNote :one
INSERT INTO order_notes (id, order_id, note, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderNote :one
SELECT * FROM order_notes WHERE id = $1;

-- name: ListOrderNotes :many
SELECT * FROM order_notes ORDER BY created_at DESC;

-- name: ListOrderNotesByOrderID :many
SELECT * FROM order_notes WHERE order_id = $1 ORDER BY created_at DESC;

-- name: DeleteOrderNote :exec
DELETE FROM order_notes WHERE id = $1;

-- name: CreateOrderDiscount :one
INSERT INTO order_discounts (id, order_id, coupon_id, discount_amount_yen)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderDiscount :one
SELECT * FROM order_discounts WHERE id = $1;

-- name: ListOrderDiscounts :many
SELECT * FROM order_discounts ORDER BY id DESC;

-- name: ListOrderDiscountsByOrderID :many
SELECT * FROM order_discounts WHERE order_id = $1 ORDER BY id DESC;

-- name: DeleteOrderDiscount :exec
DELETE FROM order_discounts WHERE id = $1;

-- name: CreateCart :one
INSERT INTO carts (id, member_id, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCart :one
SELECT * FROM carts WHERE id = $1;

-- name: ListCarts :many
SELECT * FROM carts ORDER BY created_at DESC;

-- name: ListCartsByMemberID :many
SELECT * FROM carts WHERE member_id = $1 ORDER BY created_at DESC;

-- name: DeleteCart :exec
DELETE FROM carts WHERE id = $1;

-- name: CreateCartItem :one
INSERT INTO cart_items (id, cart_id, product_id, quantity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCartItem :one
SELECT * FROM cart_items WHERE id = $1;

-- name: ListCartItems :many
SELECT * FROM cart_items ORDER BY id DESC;

-- name: ListCartItemsByCartID :many
SELECT * FROM cart_items WHERE cart_id = $1 ORDER BY id DESC;

-- name: DeleteCartItem :exec
DELETE FROM cart_items WHERE id = $1;

-- name: CreateOrderAddress :one
INSERT INTO order_addresses (id, order_id, postal_code, address)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderAddress :one
SELECT * FROM order_addresses WHERE id = $1;

-- name: ListOrderAddresses :many
SELECT * FROM order_addresses ORDER BY id DESC;

-- name: ListOrderAddressesByOrderID :many
SELECT * FROM order_addresses WHERE order_id = $1 ORDER BY id DESC;

-- name: DeleteOrderAddress :exec
DELETE FROM order_addresses WHERE id = $1;

-- name: CreateOrderTax :one
INSERT INTO order_taxes (id, order_id, tax_amount_yen, tax_rate)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderTax :one
SELECT * FROM order_taxes WHERE id = $1;

-- name: ListOrderTaxes :many
SELECT * FROM order_taxes ORDER BY id DESC;

-- name: ListOrderTaxesByOrderID :many
SELECT * FROM order_taxes WHERE order_id = $1 ORDER BY id DESC;

-- name: DeleteOrderTax :exec
DELETE FROM order_taxes WHERE id = $1;

-- name: CreateOrderCancellation :one
INSERT INTO order_cancellations (id, order_id, reason, cancelled_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderCancellation :one
SELECT * FROM order_cancellations WHERE id = $1;

-- name: ListOrderCancellations :many
SELECT * FROM order_cancellations ORDER BY cancelled_at DESC;

-- name: ListOrderCancellationsByOrderID :many
SELECT * FROM order_cancellations WHERE order_id = $1 ORDER BY cancelled_at DESC;

-- name: DeleteOrderCancellation :exec
DELETE FROM order_cancellations WHERE id = $1;

-- name: ListOrderSummaryByMember :many
-- 会員ごとの注文件数・合計金額・平均注文金額を集計する。
SELECT
    o.member_id AS member_id,
    COUNT(o.id) AS order_count,
    COALESCE(SUM(o.total_amount_yen), 0)::bigint AS total_amount_yen,
    COALESCE(AVG(o.total_amount_yen), 0)::float8 AS avg_amount_yen
FROM orders o
GROUP BY o.member_id
ORDER BY total_amount_yen DESC;

-- name: ListOrderCountByStatus :many
-- ステータスごとの注文件数を集計する。
SELECT
    o.status AS status,
    COUNT(o.id) AS order_count
FROM orders o
GROUP BY o.status
ORDER BY order_count DESC, o.status;

-- name: ListSalesQuantityAndAmountByProduct :many
-- 商品ごとの販売数量・売上合計を集計する。
SELECT
    oi.product_id AS product_id,
    COALESCE(SUM(oi.quantity), 0)::bigint AS total_quantity,
    COALESCE(SUM(oi.quantity * oi.unit_price_yen), 0)::bigint AS total_amount_yen
FROM order_items oi
GROUP BY oi.product_id
ORDER BY total_amount_yen DESC;
