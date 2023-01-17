package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	api "github.com/peacewalker122/project/api/router"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/db/repository/redis"
	"github.com/peacewalker122/project/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config: ", err.Error())
	}

	projectConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}

	err = projectConn.Ping()
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
	log.Println("Connect into postgres project database")

	redisServer, err := redis.NewRedis(config.RedisSource)
	if err != nil {
		log.Fatal("can't establish the connection: ", err.Error())
	}
	log.Println("Connect into redis")

	log.Println("initialize store")
	store := db.Newstore(projectConn)

	server, err := api.Newserver(config, store, redisServer)
	if err != nil {
		log.Fatal("can't establish the connection")
	}

	chanerr := make(chan error)

	go func(server *api.Server, chanerr chan error) {
		chanerr <- server.StartHTTP(config.HTTPServerAddress)
	}(server, chanerr)
	err = <-chanerr
	log.Println("server is running on port: ", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("can't start the server: ", err.Error())
	}
}
