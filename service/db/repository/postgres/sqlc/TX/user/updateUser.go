package user

import (
	"context"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/user"
	"github.com/peacewalker122/project/util"
)

func (u *UserTx) UpdateUserTX(ctx context.Context, arg *request.UpdateUserParamsTx) *util.MultiError {
	var errs *util.MultiError

	err := u.DBTx(ctx, func(q *db.Queries) error {
		var err error

		if arg.Username != "" {
			err = q.UpdateAccountOwner(ctx, util.InputSqlString(arg.Username))
			if err != nil {
				errs.Add(err)
				return err
			}
		}

		err = q.UpdateUserData(ctx, db.UpdateUserDataParams{
			FullName: util.InputSqlString(arg.FullName),
			Email:    util.InputSqlString(arg.Email),
			Username: util.InputSqlString(arg.Username),
		})
		if err != nil {
			errs.Add(err)
			return err
		}

		return err
	})

	if err != nil {
		errs.Add(err)
	}

	return errs
}
