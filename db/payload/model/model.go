package model

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/peacewalker122/project/db/ent"
	not "github.com/peacewalker122/project/db/payload/model/notif_query"
	"github.com/peacewalker122/project/db/payload/model/tokens"
	"github.com/peacewalker122/project/db/payload/model/users"
)

func NewModel(sql *sql.DB) Model {
	drv := entsql.OpenDB("postgres", sql)
	//defer sql.Close()

	res := ent.NewClient(ent.Driver(drv))

	return &Models{
		Client:        res,
		NotifsQueries: not.NewNotifQuery(res),
		TokenQueries:  tokens.NewTokenQuery(res),
		UserQueries:   users.NewUserQuery(res),
	}
}

type Model interface {
	TX
	not.NotifQuery
	tokens.TokensQuery
	users.UsersQuery
}

type Models struct {
	*ent.Client
	*not.NotifsQueries
	*tokens.TokenQueries
	*users.UserQueries
}
