package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/peacewalker122/project/api"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
	store := db.Newstore(conn)
	server := api.Newserver(store)

	if err := server.Start(config.HTTPServerAddress); err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
}
