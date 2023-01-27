package postgres

import (
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/tx"
	sqlcTX "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/TX"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

type PostgresStore interface {
	db.Querier
	db.Model
	sqlcTX.SQLCTX
	tx.ENTTX
}

type postgreStruct struct {
	*sqlcTX.Tx
	*model.Models
	*db.SQLStore
}

func NewPostgresStore(projectDB *sql.DB) PostgresStore {
	var res postgreStruct
	res.Tx = sqlcTX.NewTx(projectDB)
	res.Models = model.NewModel(projectDB)
	res.SQLStore = db.NewStore(projectDB)
	return &res
}
