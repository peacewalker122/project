#!/bin/sh
migrate create -ext sql -dir service/db/repository/postgres/migration/project -seq $1
#migrate -path db/$1 -database "postgresql://root:rootpass@localhost:4321/$2?sslmode=disable" $3