// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: products.sql

package mysqldb

import (
	"context"
)

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products (name, description, price)
VALUES (?, ?, ?)
`

type CreateProductParams struct {
	Name        string
	Description string
	Price       int64
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct, arg.Name, arg.Description, arg.Price)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProductByID = `-- name: GetProductByID :one
SELECT id, name, description, price, created_at, updated_at
FROM products
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetProductByID(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description, price, created_at, updated_at
FROM products
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
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
SET name = ?,
  description = ?,
  price = ?
WHERE id = ?
`

type UpdateProductParams struct {
	Name        string
	Description string
	Price       int64
	ID          string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.ID,
	)
	return err
}
