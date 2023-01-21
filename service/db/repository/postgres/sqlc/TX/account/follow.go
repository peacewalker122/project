package account

import (
	"context"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (a *AccountTx) FollowTX(ctx context.Context, isQueue bool, FromAccountID, ToAccountID int64) error {
	err := a.DBTx(ctx, func(q *db.Queries) error {
		var err error
		if isQueue {
			err = a.DeleteAccountQueue(ctx, db.DeleteAccountQueueParams{
				Fromaccountid: FromAccountID,
				Toaccountid:   ToAccountID,
			})
			if err != nil {
				return err
			}
		}
		_, _, _, err = a.UpdateFollowing(ctx, FromAccountID, ToAccountID, int64(1))
		if err != nil {
			return err
		}

		return err
	})
	return err
}
