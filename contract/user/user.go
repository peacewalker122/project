package user

import (
	"context"

	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/user"
	"github.com/peacewalker122/project/usecase/user"
	"github.com/peacewalker122/project/util"

	"github.com/google/uuid"
)

type UserContract interface {
	CreateUser(ctx context.Context, requid string, token int) (*result.CreateUserResult, error)
	CreateNewUserRequest(ctx context.Context, req *user.PayloadCreateUser) (uuid.UUID, error)
	Login(ctx context.Context, params *user.LoginParams) (*user.SessionResult, *util.Error)
	UpdateUser(ctx context.Context, arg *user.UpdateUserParam) *util.MultiError
}
