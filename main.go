package main

import (
	"context"
	"database/sql"
	"log"
	"runtime"
	"time"

	"cloud.google.com/go/storage"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/peacewalker122/project/api/router"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
	"google.golang.org/api/option"
)

func main() {
	runtime.GOMAXPROCS(2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config")
	}
	log.Println("Connect into postgres")
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
	log.Println("Connect into Google Cloud")
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(config.ClientOption))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	log.Println("initialize store")
	store, redis := db.Newstore(conn, config.RedisSource)
	server, err := api.Newserver(config, store, redis)
	if err != nil {
		log.Fatal("can't establish the connection")
	}
	if err := server.StartHTTP(config.HTTPServerAddress); err != nil {
		log.Fatal("unable to establish the connection due: ", err.Error())
	}
}
