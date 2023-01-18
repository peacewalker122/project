package main

import (
	"cloud.google.com/go/storage"
	"context"
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/service/db/repository/redis"
	"github.com/peacewalker122/project/service/gcp"
	"google.golang.org/api/option"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	api "github.com/peacewalker122/project/api/router"
	"github.com/peacewalker122/project/util"
)

var ctx context.Context

func main() {
	ctx = context.Background()

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

	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.ClientOption))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	gcpService := gcp.NewGCPService(client)

	log.Println("initialize store")
	store := db.Newstore(projectConn)

	server, err := api.Newserver(config, store, redisServer, gcpService)
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
