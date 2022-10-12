postgres:
	docker run --name postgres12 -p 4321:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootpass -d postgres:latest
migrate:
	migrate create -ext sql -dir db/migration -seq 
dropdb:
	docker exec -it postgres12 dropdb project
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root project
migrateup:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose down
sqlc:
	sqlc generate
startdb:
	docker container start postgres12
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/peacewalker122/project/db/sqlc Store

.PHONY: server startcontainer sqlc postgres migrate dropdb migrateup migratedown createdb 