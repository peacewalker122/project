package params

import (
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/account"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/users"
)

type CreateUsersOauthParam struct {
	User       *users.UsersParam
	OauthToken *tokens.TokensParams
	Account    *account.AccountParam
}

type UpdatePasswordParam struct {
}
