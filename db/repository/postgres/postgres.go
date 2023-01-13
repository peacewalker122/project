package postgres

import (
	"database/sql"

	"github.com/peacewalker122/project/db/repository/postgres/payload"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
)

type PostgresStore interface {
	db.Querier
	db.Model
	payload.Payload // we using this due tx not needed right now
}

func NewPostgresStore(projectDB, NotifDB *sql.DB) PostgresStore {
	return db.Newstore(projectDB, NotifDB)
}
