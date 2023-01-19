package tx

import (
	"context"

	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/account"
	not "github.com/peacewalker122/project/service/db/repository/postgres/payload/model/notif_query"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/params"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/users"
)

type TX interface {
	not.NotifQuery
	tokens.TokensQuery
	users.UsersQuery
	account.AccountQuery
	SetUsersOauth(ctx context.Context, params *params.CreateUsersOauthParam) (*OauthUserResponse, error)
	ChangePasswordAuth(ctx context.Context, params params.ChangePasswordParam) error
}

type Tx struct {
	*ent.Client
	*not.NotifsQueries
	*tokens.TokenQueries
	*users.UserQueries
	*account.AccountQueries
}

func NewTx(client *ent.Client) *Tx {
	return &Tx{
		Client:         client,
		NotifsQueries:  not.NewNotifQuery(client),
		TokenQueries:   tokens.NewTokenQuery(client),
		UserQueries:    users.NewUserQuery(client),
		AccountQueries: account.NewAccountQuery(client),
	}
}
