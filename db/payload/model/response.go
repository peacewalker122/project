package model

import (
	"github.com/peacewalker122/project/db/ent"
)

type OauthUserResponse struct {
	User    *ent.Users
	Token   *ent.Tokens
	Account *ent.Account
}
