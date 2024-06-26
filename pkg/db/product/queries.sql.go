// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package product

import (
	"context"
	"database/sql"
)

const getProductById = `-- name: GetProductById :one
SELECT prd_id
     , prd_name
     , prd_price
     , prd_description
     , prd_stock
     , prd_active
FROM products
WHERE prd_id = ?
`

type GetProductByIdRow struct {
	PrdID          uint64
	PrdName        string
	PrdPrice       string
	PrdDescription sql.NullString
	PrdStock       int32
	PrdActive      bool
}

func (q *Queries) GetProductById(ctx context.Context, prdID uint64) (GetProductByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getProductById, prdID)
	var i GetProductByIdRow
	err := row.Scan(
		&i.PrdID,
		&i.PrdName,
		&i.PrdPrice,
		&i.PrdDescription,
		&i.PrdStock,
		&i.PrdActive,
	)
	return i, err
}

const inactivateProduct = `-- name: InactivateProduct :exec
UPDATE products
SET prd_active = false
WHERE prd_id = ?
`

func (q *Queries) InactivateProduct(ctx context.Context, prdID uint64) error {
	_, err := q.db.ExecContext(ctx, inactivateProduct, prdID)
	return err
}

const insertProduct = `-- name: InsertProduct :exec
INSERT INTO products ( prd_name
                     , prd_price
                     , prd_description
                     , prd_stock
                     , prd_active)
VALUES (?, ?, ?, ?, ?)
`

type InsertProductParams struct {
	PrdName        string
	PrdPrice       string
	PrdDescription sql.NullString
	PrdStock       int32
	PrdActive      bool
}

func (q *Queries) InsertProduct(ctx context.Context, arg InsertProductParams) error {
	_, err := q.db.ExecContext(ctx, insertProduct,
		arg.PrdName,
		arg.PrdPrice,
		arg.PrdDescription,
		arg.PrdStock,
		arg.PrdActive,
	)
	return err
}

const listProducts = `-- name: ListProducts :many
SELECT prd_id
     , prd_name
     , prd_price
     , prd_description
     , prd_stock
     , prd_active
FROM products
WHERE prd_active = true
`

type ListProductsRow struct {
	PrdID          uint64
	PrdName        string
	PrdPrice       string
	PrdDescription sql.NullString
	PrdStock       int32
	PrdActive      bool
}

func (q *Queries) ListProducts(ctx context.Context) ([]ListProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListProductsRow
	for rows.Next() {
		var i ListProductsRow
		if err := rows.Scan(
			&i.PrdID,
			&i.PrdName,
			&i.PrdPrice,
			&i.PrdDescription,
			&i.PrdStock,
			&i.PrdActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET prd_name        = ?
  , prd_price       = ?
  , prd_description = ?
  , prd_stock       = ?
  , prd_active      = ?
WHERE prd_id = ?
`

type UpdateProductParams struct {
	PrdName        string
	PrdPrice       string
	PrdDescription sql.NullString
	PrdStock       int32
	PrdActive      bool
	PrdID          uint64
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.PrdName,
		arg.PrdPrice,
		arg.PrdDescription,
		arg.PrdStock,
		arg.PrdActive,
		arg.PrdID,
	)
	return err
}
