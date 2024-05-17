// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package product

import (
	"database/sql"
)

type Product struct {
	PrdID          uint64
	PrdName        string
	PrdPrice       string
	PrdDescription sql.NullString
	PrdStock       int32
	PrdActive      bool
	PrdCreatedAt   sql.NullTime
	PrdUpdatedAt   sql.NullTime
}