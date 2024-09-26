-- name: AddStock :exec
INSERT INTO stock (sku, total_count, reserved)
VALUES ($1, $2, $3);

-- name: UpdateStock :exec
UPDATE stock
SET total_count = $2,
    reserved = $3
WHERE sku = $1;

-- name: GetStock :exec
SELECT * FROM stock
WHERE sku = $1;

-- name: UpdateOrder :exec
UPDATE order
SET order_status = $2
WHERE order_id = $1;