package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/peacewalker122/project/util"
)

var testQueries *Queries
var testDB *sql.DB
var config util.Config

func TestMain(m *testing.M) {

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't definef config ", err.Error())
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Panic("unable to connect into db: ", err)
	}
	// storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(config.ClientOption))
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	testQueries = New(testDB)

	os.Exit(m.Run())
}
