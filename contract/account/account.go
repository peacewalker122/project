package account

import (
	"context"
	api "github.com/peacewalker122/project/api/handler"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/usecase/account"
)

type AccountContract interface {
	AcceptFollower(ctx context.Context, AccountID, FromAccount int64) (api.BasicResponse, error)
	DeleteQueue(ctx context.Context, AccountID, FromAccount int64) (api.BasicResponse, error)
	FollowAccount(ctx context.Context, FromAccount, AccountID int64) (api.BasicResponse, error)
	GetAccount(ctx context.Context, AccountID int64) (*db2.Account, error)
	ListQueuedAccount(ctx context.Context, param *account.GetAccountParams) (*[]db2.ListQueueRow, error)
	ListAccount(ctx context.Context, param *account.GetAccountParams) (*[]db2.Account, error)
	UnFollowAccount(ctx context.Context, FromAccount, AccountID int64) (api.BasicResponse, error)
	PrivateAccount(ctx context.Context, AccountID int64) (api.BasicResponse, error)
}
