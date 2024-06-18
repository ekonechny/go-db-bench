-- name: Products :many
SELECT * FROM products LIMIT $1;

-- name: Insert :exec
-- sqlc not support copyIn -> https://github.com/sqlc-dev/sqlc/issues/3264
INSERT INTO products (
    name, description, categories, price, features, color, material, upc
) SELECT name, description, categories, price, features, color, material, upc FROM UNNEST(@products::products[]);

-- name: ClearProducts :exec
TRUNCATE products;