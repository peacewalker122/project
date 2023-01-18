package model

import (
	"database/sql"
	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/account"
	not "github.com/peacewalker122/project/service/db/repository/postgres/payload/model/notif_query"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/users"

	entsql "entgo.io/ent/dialect/sql"
)

func NewModel(sql *sql.DB) Model {
	notifdrv := entsql.OpenDB("postgres", sql)
	//defer sql.Close()

	project := ent2.NewClient(ent2.Driver(notifdrv))
	return &Models{
		Client:         project,
		NotifsQueries:  not.NewNotifQuery(project),
		TokenQueries:   tokens.NewTokenQuery(project),
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
	*ent2.Client
	*not.NotifsQueries
	*tokens.TokenQueries
	*users.UserQueries
	*account.AccountQueries
}
