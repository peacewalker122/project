package tx

import (
	"context"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/likefeature"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/qouteRetweet"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/retweet_feature"

	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/account"
	not "github.com/peacewalker122/project/service/db/repository/postgres/payload/model/notif_query"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/params"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/users"
)

type ENTTX interface {
	not.NotifQuery
	tokens.TokensQuery
	users.UsersQuery
	account.AccountQuery
	retweet_feature.RetweetQuery
	qouteRetweet.QouteRetweetQuery
	likefeature.LikeQuery
	SetUsersOauth(ctx context.Context, params *params.CreateUsersOauthParam) (*OauthUserResponse, error)
	ChangePasswordAuth(ctx context.Context, params params.ChangePasswordParam) error
}

type Tx struct {
	*ent.Client
	*not.NotifsQueries
	*tokens.TokenQueries
	*users.UserQueries
	*account.AccountQueries
	*retweet_feature.RetweetQueries
	*qouteRetweet.QouteRetweetQueries
	*likefeature.LikeQueries
}

func NewTx(client *ent.Client) *Tx {
	return &Tx{
		Client:              client,
		NotifsQueries:       not.NewNotifQuery(client),
		TokenQueries:        tokens.NewTokenQuery(client),
		UserQueries:         users.NewUserQuery(client),
		AccountQueries:      account.NewAccountQuery(client),
		RetweetQueries:      retweet_feature.NewRetweetQuery(client),
		QouteRetweetQueries: qouteRetweet.NewQouteRetweetQuery(client),
		LikeQueries:         likefeature.NewLikeQuery(client),
	}
}
