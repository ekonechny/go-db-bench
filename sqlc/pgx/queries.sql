-- name: Products :many
SELECT * FROM products LIMIT $1;

-- name: Insert :copyfrom
INSERT INTO products (name, description, categories, price, features, color, material, upc)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: ClearProducts :exec
TRUNCATE products;