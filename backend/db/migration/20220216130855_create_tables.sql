-- +goose Up
CREATE TABLE stock_tickers (
  ticker_id uuid NOT NULL,
  company_name varchar(50) NOT NULL,
  ticker varchar(10) NOT NULL,
  favorited bool NOT NULL DEFAULT (false),
  tags json,
  added_date_time timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY(ticker_id)
);

-- +goose Down
DROP TABLE IF EXISTS stock_tickers;
