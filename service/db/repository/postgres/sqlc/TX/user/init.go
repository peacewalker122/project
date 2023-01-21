package user

import (
	"context"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/user"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/user"
)

type UserDBTX interface {
	CreateUserTX(ctx context.Context, arg *request.CreateUserParamsTx) (*result.CreateUserResult, error)
}

type UserTx struct {
	*db.SQLStore
}

func NewUserTx(db *db.SQLStore) *UserTx {
	return &UserTx{
		SQLStore: db,
	}
}
