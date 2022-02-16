# StockApp

## Prerequisites
* [Docker](https://www.docker.com/)
* [Goose](https://github.com/pressly/goose)
* [SQLC](https://github.com/kyleconroy/sqlc)

## Commands

Creating a docker container with PostgreSQL image
```bash
docker run --name ticker_container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

Creating databae inside \db\migration
```bash
docker exec -it ticker_container createdb --username=root --owner=root stock_app
```

Migrate up / down
```bash
goose postgres "user=root password=secret dbname=stock_app sslmode=disable" up
goose postgres "user=root password=secret dbname=stock_app sslmode=disable" down
```

## Repository

* [SQLC] (https://github.com/kyleconroy/sqlc)

Create sqlc.yaml file
```bash
init sqlc
```
Lim inn dette i sqlc.yaml
```bash
version: 1
packages:
  - path: "./db/sqlc"
    name: "db"
    engine: "postgresql"
    schema: "./db/migration"
    queries: "./db/query"
```
