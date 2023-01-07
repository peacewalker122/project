package model

import (
	"github.com/peacewalker122/project/db/payload/model/tokens"
	"github.com/peacewalker122/project/db/payload/model/users"
)

type CreateUsersOauthParam struct {
	User       *users.UsersParam
	OauthToken *tokens.TokensParams
}
