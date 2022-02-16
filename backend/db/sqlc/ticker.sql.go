// Code generated by sqlc. DO NOT EDIT.
// source: ticker.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addTicker = `-- name: AddTicker :one
INSERT INTO stock_tickers (
    ticker_id,
    company_name,
    ticker
) VALUES (
    uuid_generate_v4(), $1, $2
) RETURNING ticker_id, company_name, ticker, favorited, tags, added_date_time
`

type AddTickerParams struct {
	CompanyName string
	Ticker      string
}

func (q *Queries) AddTicker(ctx context.Context, arg AddTickerParams) (StockTicker, error) {
	row := q.db.QueryRowContext(ctx, addTicker, arg.CompanyName, arg.Ticker)
	var i StockTicker
	err := row.Scan(
		&i.TickerID,
		&i.CompanyName,
		&i.Ticker,
		&i.Favorited,
		&i.Tags,
		&i.AddedDateTime,
	)
	return i, err
}

const addToFavorites = `-- name: AddToFavorites :one
UPDATE stock_tickers
SET favorited = TRUE
WHERE ticker_id = $1
RETURNING ticker_id, company_name, ticker, favorited, tags, added_date_time
`

func (q *Queries) AddToFavorites(ctx context.Context, tickerID uuid.UUID) (StockTicker, error) {
	row := q.db.QueryRowContext(ctx, addToFavorites, tickerID)
	var i StockTicker
	err := row.Scan(
		&i.TickerID,
		&i.CompanyName,
		&i.Ticker,
		&i.Favorited,
		&i.Tags,
		&i.AddedDateTime,
	)
	return i, err
}

const deleteTicker = `-- name: DeleteTicker :exec
DELETE FROM stock_tickers
WHERE ticker_id = $1
`

func (q *Queries) DeleteTicker(ctx context.Context, tickerID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTicker, tickerID)
	return err
}

const getTicker = `-- name: GetTicker :one
SELECT ticker_id, company_name, ticker, favorited, tags, added_date_time FROM stock_tickers
WHERE ticker_id = $1 LIMIT 1
`

func (q *Queries) GetTicker(ctx context.Context, tickerID uuid.UUID) (StockTicker, error) {
	row := q.db.QueryRowContext(ctx, getTicker, tickerID)
	var i StockTicker
	err := row.Scan(
		&i.TickerID,
		&i.CompanyName,
		&i.Ticker,
		&i.Favorited,
		&i.Tags,
		&i.AddedDateTime,
	)
	return i, err
}

const listTickers = `-- name: ListTickers :many
SELECT ticker_id, company_name, ticker, favorited, tags, added_date_time FROM stock_tickers
ORDER BY ticker_id
LIMIT $1
OFFSET $2
`

type ListTickersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) ListTickers(ctx context.Context, arg ListTickersParams) ([]StockTicker, error) {
	rows, err := q.db.QueryContext(ctx, listTickers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StockTicker
	for rows.Next() {
		var i StockTicker
		if err := rows.Scan(
			&i.TickerID,
			&i.CompanyName,
			&i.Ticker,
			&i.Favorited,
			&i.Tags,
			&i.AddedDateTime,
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

const removeFromFavorites = `-- name: RemoveFromFavorites :one
UPDATE stock_tickers
SET favorited = FALSE
WHERE ticker_id = $1
RETURNING ticker_id, company_name, ticker, favorited, tags, added_date_time
`

func (q *Queries) RemoveFromFavorites(ctx context.Context, tickerID uuid.UUID) (StockTicker, error) {
	row := q.db.QueryRowContext(ctx, removeFromFavorites, tickerID)
	var i StockTicker
	err := row.Scan(
		&i.TickerID,
		&i.CompanyName,
		&i.Ticker,
		&i.Favorited,
		&i.Tags,
		&i.AddedDateTime,
	)
	return i, err
}

const updateTicker = `-- name: UpdateTicker :one
UPDATE stock_tickers
SET company_name = $2
WHERE ticker_id = $1
RETURNING ticker_id, company_name, ticker, favorited, tags, added_date_time
`

type UpdateTickerParams struct {
	TickerID    uuid.UUID
	CompanyName string
}

func (q *Queries) UpdateTicker(ctx context.Context, arg UpdateTickerParams) (StockTicker, error) {
	row := q.db.QueryRowContext(ctx, updateTicker, arg.TickerID, arg.CompanyName)
	var i StockTicker
	err := row.Scan(
		&i.TickerID,
		&i.CompanyName,
		&i.Ticker,
		&i.Favorited,
		&i.Tags,
		&i.AddedDateTime,
	)
	return i, err
}
