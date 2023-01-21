package user

import (
	"context"
	"errors"
	"github.com/lib/pq"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/user"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/user"
	"net/http"
)

func (u *UserTx) CreateUserTX(ctx context.Context, arg *request.CreateUserParamsTx) (*result.CreateUserResult, error) {
	var res result.CreateUserResult
	err := u.DBTx(ctx, func(q *db.Queries) error {
		var err error

		res.User, err = q.CreateUser(ctx, db.CreateUserParams{
			Username:       arg.Username,
			HashedPassword: arg.Password,
			Email:          arg.Email,
			FullName:       arg.FullName,
		})
		if err != nil {
			if pqerr, ok := err.(*pq.Error); ok {
				switch pqerr.Code.Name() {
				case "unique_violation":
					res.ErrCode = http.StatusForbidden
					res.Error = errors.New("username or email already exists")
				}
			}
			res.ErrCode = http.StatusInternalServerError
			res.Error = err
		}

		res.Account, err = q.CreateAccounts(ctx, db.CreateAccountsParams{
			Owner: arg.Username,
		})
		if err != nil {
			if pqerr, ok := err.(*pq.Error); ok {
				switch pqerr.Code.Name() {
				case "unique_violation", "foreign_key_violation":
					res.ErrCode = http.StatusForbidden
					res.Error = err
				}
			}
			res.ErrCode = http.StatusInternalServerError
			res.Error = err
		}
		return err
	})
	return &res, err
}
