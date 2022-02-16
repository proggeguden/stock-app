package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/proggeguden/stock-app/backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTicker(t *testing.T) StockTicker {
	arg := AddTickerParams{
		CompanyName: util.RandomCompanyName(),
		Ticker:      util.RandomTicker(),
	}
	stock_ticker, err := testQueries.AddTicker(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, stock_ticker)

	require.Equal(t, arg.CompanyName, stock_ticker.CompanyName)
	require.Equal(t, arg.Ticker, stock_ticker.Ticker)

	require.NotZero(t, stock_ticker.TickerID)
	require.NotZero(t, stock_ticker.AddedDateTime)

	return stock_ticker
}

func TestAddTicker(t *testing.T) {
	createRandomTicker(t)
}

func TestGetTicker(t *testing.T) {
	ticker1 := createRandomTicker(t)
	ticker2, err := testQueries.GetTicker(context.Background(), ticker1.TickerID)
	require.NoError(t, err)
	require.NotEmpty(t, ticker2)

	require.Equal(t, ticker1.TickerID, ticker2.TickerID)
	require.Equal(t, ticker1.CompanyName, ticker2.CompanyName)
	require.Equal(t, ticker1.Ticker, ticker2.Ticker)

	require.WithinDuration(t, ticker1.AddedDateTime, ticker2.AddedDateTime, time.Second)
}

func TestUpdateTicker(t *testing.T) {
	ticker1 := createRandomTicker(t)

	arg := UpdateTickerParams{
		TickerID:    ticker1.TickerID,
		CompanyName: util.RandomCompanyName(),
	}

	ticker2, err := testQueries.UpdateTicker(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, ticker2)

	require.Equal(t, ticker1.TickerID, ticker2.TickerID)
	require.Equal(t, arg.CompanyName, ticker2.CompanyName)
	require.Equal(t, ticker1.Ticker, ticker2.Ticker)

	require.WithinDuration(t, ticker1.AddedDateTime, ticker2.AddedDateTime, time.Second)
}

func TestDeleteTicker(t *testing.T) {
	ticker1 := createRandomTicker(t)
	err := testQueries.DeleteTicker(context.Background(), ticker1.TickerID)
	require.NoError(t, err)

	ticker2, err := testQueries.GetTicker(context.Background(), ticker1.TickerID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, ticker2)
}

func TestListTickers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTicker(t)
	}

	arg := ListTickersParams{
		Limit:  5,
		Offset: 5,
	}

	tickers, err := testQueries.ListTickers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tickers, 5)

	for _, ticker := range tickers {
		require.NotEmpty(t, ticker)
	}
}

func TestAddToFavorites(t *testing.T) {
	ticker1 := createRandomTicker(t)
	require.Equal(t, ticker1.Favorited, bool(false))

	ticker2, err := testQueries.AddToFavorites(context.Background(), ticker1.TickerID)
	require.NoError(t, err)
	require.Equal(t, ticker1.TickerID, ticker2.TickerID)
	require.Equal(t, ticker2.Favorited, true)
}

func TestRemoveFromFavorites(t *testing.T) {
	ticker1 := createRandomTicker(t)
	_, err := testQueries.AddToFavorites(context.Background(), ticker1.TickerID)
	require.NoError(t, err)

	ticker1, err2 := testQueries.GetTicker(context.Background(), ticker1.TickerID)
	require.NoError(t, err2)
	require.Equal(t, ticker1.Favorited, bool(true))

	ticker1, err3 := testQueries.RemoveFromFavorites(context.Background(), ticker1.TickerID)
	require.NoError(t, err3)
	require.Equal(t, ticker1.Favorited, bool(false))
}
