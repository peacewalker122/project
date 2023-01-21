package account

import (
	"context"

	api "github.com/peacewalker122/project/api/handler"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (a *AccountUseCase) AcceptFollower(ctx context.Context, AccountID, FromAccount int64) (api.BasicResponse, error) {
	var resultData map[string]interface{}

	isQueue, err := a.postgre.GetAccountsQueue(ctx, db2.GetAccountsQueueParams{
		Fromid: FromAccount,
		Toid:   AccountID,
	})
	if err != nil {
		return nil, err
	}
	err = a.postgre.FollowTX(
		ctx,
		isQueue,
		FromAccount,
		AccountID,
	)
	if err != nil {
		return nil, err
	}

	resultData = map[string]interface{}{
		"status": "accepted",
	}

	return resultData, nil
}
