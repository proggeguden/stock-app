postgres:
	docker run --name ticker_container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
	
createdb:
	docker exec -it ticker_container createdb --username=root --owner=root stock_app

dropdb:
	docker exec -it ticker_container dropdb stock_app

migrateup: 
	goose postgres "user=root password=secret dbname=stock_app sslmode=disable" up

migratedown:
	goose postgres "user=root password=secret dbname=stock_app sslmode=disable" down

sqlc:
	docker run --rm -v C:\Users\jakob\Desktop\stock-app\backend:/src -w /src kjconroy/sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc