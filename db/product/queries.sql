-- name: InsertProduct :exec
INSERT INTO products ( prd_name
                     , prd_price
                     , prd_description
                     , prd_stock
                     , prd_active)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateProduct :exec
UPDATE products
SET prd_name        = $2
  , prd_price       = $3
  , prd_description = $4
  , prd_stock       = $5
  , prd_active      = $6
WHERE prd_id = $1;

-- name: InactivateProduct :exec
UPDATE products
SET prd_active = false
WHERE prd_id = $1;

-- name: ListProducts :one
SELECT prd_id
     , prd_name
     , prd_price
     , prd_description
     , prd_stock
     , prd_active
FROM products;

-- name: GetProductById :one
SELECT prd_id
     , prd_name
     , prd_price
     , prd_description
     , prd_stock
     , prd_active
FROM products
WHERE prd_id = $1;