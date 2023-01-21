package account

import (
	"context"
	"errors"

	api "github.com/peacewalker122/project/api/handler"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (a *AccountUseCase) FollowAccount(ctx context.Context, FromAccount, AccountID int64) (api.BasicResponse, error) {

	var resultData map[string]interface{}

	IsPrivate, err := a.postgre.GetAccountsInfo(ctx, AccountID)
	if err != nil {
		return nil, err
	}

	num, err := a.postgre.GetAccountsFollowRows(ctx, db2.GetAccountsFollowRowsParams{
		Fromid: FromAccount,
		Toid:   AccountID,
	})
	if err != nil {
		return nil, err
	}
	if IsPrivate.IsPrivate {

		isQueue, err := a.postgre.GetQueueRows(ctx, db2.GetQueueRowsParams{
			Fromaccountid: FromAccount,
			Toaccountid:   AccountID,
		})

		if isQueue != 0 {
			return nil, errors.New("already in queue")
		}

		ok, err := a.postgre.CreateAccountsQueueTX(ctx, db2.CreateAccountQueueParams{
			FromAccountID: FromAccount,
			ToAccountID:   AccountID,
		})
		if err != nil || !ok {
			if !ok {
				err = errors.New("can't proceed queue")
			}
			return nil, err
		}

		resultData = map[string]interface{}{
			"status": "queue",
		}

		return resultData, nil
	}

	if num != 0 {
		return nil, errors.New("already follow")
	}

	err = a.postgre.FollowTX(
		ctx,
		false,
		FromAccount,
		AccountID,
	)
	if err != nil {
		return nil, err
	}

	resultData = map[string]interface{}{
		"status": "followed",
	}

	return resultData, nil
}
