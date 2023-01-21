package model

import (
	"database/sql"

	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/tx"

	entsql "entgo.io/ent/dialect/sql"
)

func NewModel(sql *sql.DB) *Models {
	notifdrv := entsql.OpenDB("postgres", sql)
	//defer sql.Close()

	project := ent2.NewClient(ent2.Driver(notifdrv))

	var res Models

	res.Tx = tx.NewTx(project)

	return &res
}
type Model interface {
	tx.ENTTX
}

type Models struct {
	*tx.Tx
}
