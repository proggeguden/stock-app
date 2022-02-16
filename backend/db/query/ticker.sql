-- name: AddTicker :one
INSERT INTO stock_tickers (
    ticker_id,
    company_name,
    ticker
) VALUES (
    uuid_generate_v4(), $1, $2
) RETURNING *;

-- name: GetTicker :one
SELECT * FROM stock_tickers
WHERE ticker_id = $1 LIMIT 1;

-- name: ListTickers :many
SELECT * FROM stock_tickers
ORDER BY ticker_id
LIMIT $1
OFFSET $2;

-- name: UpdateTicker :one
UPDATE stock_tickers
SET company_name = $2
WHERE ticker_id = $1
RETURNING *;

-- name: DeleteTicker :exec
DELETE FROM stock_tickers
WHERE ticker_id = $1;

-- name: AddToFavorites :one
UPDATE stock_tickers
SET favorited = TRUE
WHERE ticker_id = $1
RETURNING *;

-- name: RemoveFromFavorites :one
UPDATE stock_tickers
SET favorited = FALSE
WHERE ticker_id = $1
RETURNING *;