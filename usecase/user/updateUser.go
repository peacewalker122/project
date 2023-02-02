package user

import (
	"context"
	"errors"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/user"
	"github.com/peacewalker122/project/util"
)

func (u *UserUsecase) UpdateUser(ctx context.Context, arg *UpdateUserParam) *util.MultiError {
	var errs *util.MultiError

	if util.AlphaNumCheck(arg.Username) {
		errs.Add(errors.New("username must be alphanumeric"))
	}

	if util.AlphaCheck(arg.FullName) {
		errs.Add(errors.New("name must be alphabet"))
	}

	_, err := u.postgre.GetAccountByEmail(ctx, arg.Email)
	if err == nil {
		errs.Add(errors.New("email already exist"))
	}
	_, err = u.postgre.GetUser(ctx, arg.Username)
	if err == nil {
		errs.Add(errors.New("username already exist"))
	}

	err = u.postgre.UpdateUserTX(ctx, &request.UpdateUserParamsTx{
		Username: arg.Username,
		Password: arg.Password,
		Email:    arg.Email,
		FullName: arg.FullName,
	})
	if err != nil {
		errs.Add(err)
	}

	return errs
}
