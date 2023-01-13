package main

import (
	"database/sql"
	"log"
	"runtime"

	_ "github.com/golang/mock/mockgen/model"
	api "github.com/peacewalker122/project/api/router"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/db/repository/redis"
	"github.com/peacewalker122/project/util"
)

func main() {
	runtime.GOMAXPROCS(2)

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config: ", err.Error())
	}
	log.Println("Connect into postgres")

	projectConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
	defer projectConn.Close()

	notifConn, err := sql.Open(config.DBDriver, config.NotifDBSource)
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
	defer notifConn.Close()

	redis, err := redis.NewRedis(config.RedisSource)
	if err != nil {
		log.Fatal("can't establish the connection: ", err.Error())
	}

	log.Println("initialize store")
	store := db.Newstore(projectConn, notifConn)

	server, err := api.Newserver(config, store, redis)
	if err != nil {
		log.Fatal("can't establish the connection")
	}

	chanerr := make(chan error)

	go func(server *api.Server, chanerr chan error) {
		chanerr <- server.StartHTTP(config.HTTPServerAddress)
	}(server, chanerr)

	err = <-chanerr
	if err != nil {
		log.Fatal("can't start the server: ", err.Error())
	}
}
