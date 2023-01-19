package tx

import (
	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
)

type OauthUserResponse struct {
	User    *ent2.Users
	Token   *ent2.Tokens
	Account *ent2.Account
}
