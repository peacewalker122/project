package user

import (
	"context"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/TX/account"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/user"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/user"
	"github.com/peacewalker122/project/util"
)

type UserDBTX interface {
	CreateUserTX(ctx context.Context, arg *request.CreateUserParamsTx) (*result.CreateUserResult, error)
	UpdateUserTX(ctx context.Context, arg *request.UpdateUserParamsTx) *util.MultiError
}

type UserTx struct {
	acc *account.AccountTx
	*db.SQLStore
}

func NewUserTx(db *db.SQLStore) *UserTx {
	return &UserTx{
		SQLStore: db,
		acc:      account.NewAccountTx(db),
	}
}
