package main

import (
	"database/sql"
	"log"

	"github.com/peacewalker122/project/api"
	db "github.com/peacewalker122/project/db/sqlc"
)

const (
	driverName    = "postgres"
	DBsource      = "postgresql://root:rootpass@localhost:4321/project?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(driverName, DBsource)
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
	store := db.Newstore(conn)
	server := api.Newserver(store)

	if err := server.Start(serverAddress); err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
}
