-- name: GetProductByID :one
SELECT *
FROM products
WHERE id = ?
LIMIT 1;
-- name: ListProducts :many
SELECT *
FROM products;
-- name: CreateProduct :exec
INSERT INTO products (name, description, price)
VALUES (?, ?, ?);
-- name: UpdateProduct :exec
UPDATE products
SET name = ?,
  description = ?,
  price = ?
WHERE id = ?;
-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;