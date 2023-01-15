package user

import (
	"context"

	"github.com/google/uuid"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/util"
	"github.com/peacewalker122/project/usecase/user"
)

type UserContract interface {
	CreateUser(ctx context.Context, requid string, token int) (*db.CreateUserTXResult, *util.Error)
	CreateNewUserRequest(ctx context.Context, req db.CreateUserParams) (uuid.UUID, *util.Error)
	Login(ctx context.Context, params user.SessionParams) (*user.SessionResult, *util.Error)
}
