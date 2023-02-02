package user

import (
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
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
		Account        db.Account
		User           db.User
		RefreshPayload *token.Payload
		Session        db.Session
	}
	PayloadCreateUser struct {
		Username string `json:"username"`
		Password string `json:"hashed_password"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}
	LoginParams struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		ClientIp  string `json:"client_ip"`
		UserAgent string `json:"user_agent"`
	}
	UpdateUserParam struct {
		Username string `json:"username"`
		Password string `json:"hashed_password"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}
)

func (u *UpdateUserParam) Validate() *util.MultiError {
	var errs util.MultiError

	if !util.AlphaNumCheck(u.Username) {
		err := errors.New("invalid username")
		errs.Add(err, 400)
	}

	if !util.AlphaCheck(u.FullName) {
		err := errors.New("invalid fullname")
		errs.Add(err, 400)
	}

	return &errs
}

func (p *PayloadCreateUser) Validate() *util.MultiError {
	var errs util.MultiError

	if !util.AlphaNumCheck(p.Username) {
		err := errors.New("invalid username")
		errs.Add(err, 400)
	}

	if !util.AlphaCheck(p.FullName) {
		err := errors.New("invalid fullname")
		errs.Add(err, 400)
	}

	return &errs
}
