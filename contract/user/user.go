package user

import (
	"context"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/usecase/user"
	"github.com/peacewalker122/project/util"

	"github.com/google/uuid"
)

type UserContract interface {
	CreateUser(ctx context.Context, requid string, token int) (*db2.CreateUserTXResult, error)
	CreateNewUserRequest(ctx context.Context, req db2.CreateUserParams) (uuid.UUID, error)
	Login(ctx context.Context, params *user.LoginParams) (*user.SessionResult, *util.Error)
}
