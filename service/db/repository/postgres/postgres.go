package postgres

import (
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
)

type PostgresStore interface {
	db2.Querier
	db2.Model
	payload.Payload // we using this due tx not needed right now
}

func NewPostgresStore(projectDB *sql.DB) PostgresStore {
	return db2.Newstore(projectDB)
}
