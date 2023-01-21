package account

import (
	"context"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

type AccountDBTX interface {
	FollowTX(ctx context.Context, isQueue bool, FromAccountID, ToAccountID int64) error
	UnFollowTX(ctx context.Context, FromAccountID, ToAccountID int64) error
}

type AccountTx struct {
	*db.SQLStore
}

func NewAccountTx(db *db.SQLStore) *AccountTx {
	return &AccountTx{
		SQLStore: db,
	}
}
