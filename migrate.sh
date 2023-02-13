#!/bin/sh
#migrate create -ext sql -dir db/$1 -seq $2
migrate -path service/db/repository/postgres/migration/project -database "postgresql://root:rootpass@localhost:5432/project?sslmode=disable" $1