package model

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/peacewalker122/project/db/ent"
	not "github.com/peacewalker122/project/db/model/notif_query"
	"github.com/peacewalker122/project/db/model/tokens"
)

func NewModel(sql *sql.DB,) Model {
	drv := entsql.OpenDB("postgres", sql)
	//defer sql.Close()

	res := ent.NewClient(ent.Driver(drv))

	return &Models{
		NotifsQueries: not.NewNotifQuery(res),
		TokenQueries:  tokens.NewTokenQuery(res),
	}
}

type Model interface {
	not.NotifQuery
	tokens.TokensQuery
}

type Models struct {
	*not.NotifsQueries
	*tokens.TokenQueries
}
