postgres:
	docker run --name postgres12 -p 4321:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootpass -d postgres:latest
migrate:
	migrate create -ext sql -dir db/migration -seq account_feature
dropdb:
	docker exec -it postgres12 dropdb project
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root project
migrateup:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose up 1
migrateup2:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose up 2
migratedown:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose down 1
migratedown2:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:4321/project?sslmode=disable" -verbose down 2
sqlc:
	sqlc generate
startdb:
	docker container start postgres12
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/peacewalker122/project/db/sqlc Store

.PHONY: migrateup2 migratedown2 migratedown1 migrateup1 server startcontainer sqlc postgres migrate dropdb migrateup migratedown createdb 