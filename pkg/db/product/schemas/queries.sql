-- name: InsertProduct :exec
INSERT INTO products ( prd_name
                     , prd_price
                     , prd_description
                     , prd_stock
                     , prd_active)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateProduct :exec
UPDATE products
SET prd_name        = ?
  , prd_price       = ?
  , prd_description = ?
  , prd_stock       = ?
  , prd_active      = ?
WHERE prd_id = ?;

-- name: InactivateProduct :exec
UPDATE products
SET prd_active = false
WHERE prd_id = ?;

-- name: ListProducts :many
SELECT prd_id
     , prd_name
     , prd_price
     , prd_description
     , prd_stock
     , prd_active
FROM products
WHERE prd_active = true;

-- name: GetProductById :one
SELECT prd_id
     , prd_name
     , prd_price
     , prd_description
     , prd_stock
     , prd_active
FROM products
WHERE prd_id = ?;