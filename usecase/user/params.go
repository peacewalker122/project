package user

import (
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/token"
)

type (
	SessionParams struct {
		ID        *int64
		Username  string
		Email     string
		UserAgent string
		ClientIp  string
		IsBlocked bool
	}
	SessionResult struct {
		AccessToken    string
		RefreshToken   string
		AccessPayload  *token.Payload
		RefreshPayload *token.Payload
		Session        db.Session
	}
	PayloadCreateUser struct {
		Username string `json:"username"`
		Password string `json:"hashed_password"`
		FullName string `json:"full_name"`
		Email    string	`json:"email"`
	}
)
