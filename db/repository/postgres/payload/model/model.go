package model

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/peacewalker122/project/db/repository/postgres/ent"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/account"
	not "github.com/peacewalker122/project/db/repository/postgres/payload/model/notif_query"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/users"
)

func NewModel(sql ...*sql.DB) Model {
	drv := entsql.OpenDB("postgres", sql[0])
	notifdrv := entsql.OpenDB("postgres", sql[1])
	//defer sql.Close()

	res := ent.NewClient(ent.Driver(drv))
	project := ent.NewClient(ent.Driver(notifdrv))
	return &Models{
		Client:         res,
		NotifsQueries:  not.NewNotifQuery(res),
		TokenQueries:   tokens.NewTokenQuery(res),
		UserQueries:    users.NewUserQuery(project),
		AccountQueries: account.NewAccountQuery(project),
	}
}

type Model interface {
	TX
	not.NotifQuery
	tokens.TokensQuery
	users.UsersQuery
	account.AccountQuery
}

type Models struct {
	*ent.Client
	*not.NotifsQueries
	*tokens.TokenQueries
	*users.UserQueries
	*account.AccountQueries
}
