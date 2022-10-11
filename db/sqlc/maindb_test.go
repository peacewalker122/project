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

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("can't definef config ", err.Error())
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Panic("unable to connect into db: ", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
