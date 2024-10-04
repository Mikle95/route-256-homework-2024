-- name: AddUpdateStock :exec
INSERT INTO stock (sku, total_count, reserved)
VALUES (@sku, @total_count, @reserved)
ON CONFLICT(sku) 
DO UPDATE SET total_count = @total_count, reserved = @reserved;

-- name: GetStock :one
SELECT * FROM stock
WHERE sku = @sku;

-- name: InsertOrder :one
INSERT INTO user_order (user_id, current_status)
VALUES (@user_id, @current_status)
RETURNING order_id;


-- name: UpdateOrderStatus :exec
UPDATE user_order
SET current_status = @current_status
WHERE order_id = @order_id;


-- name: SelectOrder :one
SELECT * FROM user_order
WHERE order_id = @order_id;


-- name: InsertItem :exec
INSERT INTO item (sku, order_id, count)
VALUES (@sku, @order_id, @count);


-- name: SelectItem :many
SELECT * FROM item
WHERE order_id = @order_id;