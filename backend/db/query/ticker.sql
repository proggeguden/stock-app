-- name: AddTicker :one
INSERT INTO stock_tickers (
    company_name,
    ticker
) VALUES (
    $1, $2
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