package result

import db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"

type CreateUserResult struct {
	User    db.User    `json:"user"`
	Account db.Account `json:"account"`
	Error   error      `json:"error"`
	ErrCode int        `json:"err_code"`
}
