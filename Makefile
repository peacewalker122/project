postgres:
	docker run --name postgres12 -p 4321:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootpass -d postgres:latest
migrate:
	migrate create -ext sql -dir db/migration -seq account_feature
dropdb:
	docker exec -it postgres12 dropdb project
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root project
ent:
	entc generate ./service/db/repository/postgres/ent/schema/*.go
schema:
	go run -mod=mod entgo.io/ent/cmd/ent init --target ./service/db/repository/postgres/schema $1
sqlc:
	sqlc generate
startdb:
	docker container start postgres12 && docker container start redisContainer
server:
	go run main.go
mock:
	mockgen -package mockdb -destination service/db/repository/postgres/mock/store.go github.com/peacewalker122/project/service/db/repository/postgres PostgresStore
psql:
	docker exec -it postgres12 psql -U root project

.PHONY: migrateup2 migratedown2 migratedown1 migrateup1 server startcontainer sqlc postgres migrate dropdb migrateup migratedown createdb 