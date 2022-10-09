package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:rootpass@localhost:4321/project?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Panic("unable to connect into db: ", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
