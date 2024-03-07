postgres:
	docker run -d --name practice -p 5432:5432 -e POSTGRES_USER=pasan -e POSTGRES_PASSWORD=12345 postgres:16-alpine

createdb:
	docker exec -it practice createdb --username=pasan --owner=pasan practice

dropdb:
	docker exec -it practice dropdb practice

migrateup:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/practice?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://pasan:12345@localhost:5432/practice?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test