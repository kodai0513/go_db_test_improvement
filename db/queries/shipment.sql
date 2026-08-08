-- name: CreateShipment :one
INSERT INTO shipments (id, order_id, address, status, shipped_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetShipment :one
SELECT * FROM shipments WHERE id = $1;

-- name: ListShipments :many
SELECT * FROM shipments ORDER BY id;

-- name: CreateShipmentItem :one
INSERT INTO shipment_items (id, shipment_id, order_item_id, quantity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetShipmentItem :one
SELECT * FROM shipment_items WHERE id = $1;

-- name: ListShipmentItems :many
SELECT * FROM shipment_items ORDER BY id;

-- name: ListShipmentItemsByShipmentID :many
SELECT * FROM shipment_items WHERE shipment_id = $1 ORDER BY id;

-- name: DeleteShipmentItem :exec
DELETE FROM shipment_items WHERE id = $1;

-- name: CreateCarrier :one
INSERT INTO carriers (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetCarrier :one
SELECT * FROM carriers WHERE id = $1;

-- name: ListCarriers :many
SELECT * FROM carriers ORDER BY id;

-- name: DeleteCarrier :exec
DELETE FROM carriers WHERE id = $1;

-- name: CreateTrackingEvent :one
INSERT INTO tracking_events (id, shipment_id, status, occurred_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetTrackingEvent :one
SELECT * FROM tracking_events WHERE id = $1;

-- name: ListTrackingEvents :many
SELECT * FROM tracking_events ORDER BY occurred_at DESC;

-- name: ListTrackingEventsByShipmentID :many
SELECT * FROM tracking_events WHERE shipment_id = $1 ORDER BY occurred_at DESC;

-- name: DeleteTrackingEvent :exec
DELETE FROM tracking_events WHERE id = $1;

-- name: CreateShippingAddress :one
INSERT INTO shipping_addresses (id, shipment_id, postal_code, address)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetShippingAddress :one
SELECT * FROM shipping_addresses WHERE id = $1;

-- name: ListShippingAddresses :many
SELECT * FROM shipping_addresses ORDER BY id;

-- name: ListShippingAddressesByShipmentID :many
SELECT * FROM shipping_addresses WHERE shipment_id = $1 ORDER BY id;

-- name: DeleteShippingAddress :exec
DELETE FROM shipping_addresses WHERE id = $1;

-- name: CreateShippingRate :one
INSERT INTO shipping_rates (id, carrier_id, zone_name, price_yen)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetShippingRate :one
SELECT * FROM shipping_rates WHERE id = $1;

-- name: ListShippingRates :many
SELECT * FROM shipping_rates ORDER BY id;

-- name: ListShippingRatesByCarrierID :many
SELECT * FROM shipping_rates WHERE carrier_id = $1 ORDER BY id;

-- name: DeleteShippingRate :exec
DELETE FROM shipping_rates WHERE id = $1;

-- name: CreateShippingZone :one
INSERT INTO shipping_zones (id, name)
VALUES ($1, $2)
RETURNING *;

-- name: GetShippingZone :one
SELECT * FROM shipping_zones WHERE id = $1;

-- name: ListShippingZones :many
SELECT * FROM shipping_zones ORDER BY id;

-- name: DeleteShippingZone :exec
DELETE FROM shipping_zones WHERE id = $1;

-- name: CreateDeliveryAttempt :one
INSERT INTO delivery_attempts (id, shipment_id, attempted_at, succeeded)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetDeliveryAttempt :one
SELECT * FROM delivery_attempts WHERE id = $1;

-- name: ListDeliveryAttempts :many
SELECT * FROM delivery_attempts ORDER BY attempted_at DESC;

-- name: ListDeliveryAttemptsByShipmentID :many
SELECT * FROM delivery_attempts WHERE shipment_id = $1 ORDER BY attempted_at DESC;

-- name: DeleteDeliveryAttempt :exec
DELETE FROM delivery_attempts WHERE id = $1;

-- name: CreateReturnRequest :one
INSERT INTO return_requests (id, order_id, reason, status, created_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetReturnRequest :one
SELECT * FROM return_requests WHERE id = $1;

-- name: ListReturnRequests :many
SELECT * FROM return_requests ORDER BY created_at DESC;

-- name: ListReturnRequestsByOrderID :many
SELECT * FROM return_requests WHERE order_id = $1 ORDER BY created_at DESC;

-- name: DeleteReturnRequest :exec
DELETE FROM return_requests WHERE id = $1;

-- name: CreateReturnItem :one
INSERT INTO return_items (id, return_request_id, order_item_id, quantity)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetReturnItem :one
SELECT * FROM return_items WHERE id = $1;

-- name: ListReturnItems :many
SELECT * FROM return_items ORDER BY id;

-- name: ListReturnItemsByReturnRequestID :many
SELECT * FROM return_items WHERE return_request_id = $1 ORDER BY id;

-- name: DeleteReturnItem :exec
DELETE FROM return_items WHERE id = $1;
