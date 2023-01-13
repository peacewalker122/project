package model

import (
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/tokens"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/users"
	"github.com/peacewalker122/project/db/repository/postgres/payload/model/account"
)

type CreateUsersOauthParam struct {
	User       *users.UsersParam
	OauthToken *tokens.TokensParams
	Account    *account.AccountParam
}
