#!/bin/sh
migrate create -ext sql -dir db/migration/$1 -seq $2
#migrate -path db/$1 -database "postgresql://root:rootpass@localhost:4321/$2?sslmode=disable" $3