package account

import (
	"context"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (a *AccountTx) UnFollowTX(ctx context.Context, FromAccountID, ToAccountID int64) error {
	err := a.DBTx(ctx, func(q *db.Queries) error {
		var err error
		_, _, _, err = a.DeleteFollowing(ctx, FromAccountID, ToAccountID, int64(-1))
		if err != nil {
			return err
		}
		return err
	})
	return err
}
