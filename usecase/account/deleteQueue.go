package account

import (
	"context"
	"errors"
	api "github.com/peacewalker122/project/api/handler"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (a *AccountUseCase) DeleteQueue(ctx context.Context, AccountID, FromAccount int64) (api.BasicResponse, error) {
	var resultData map[string]interface{}

	isQueue, err := a.postgre.GetQueueRows(ctx, db2.GetQueueRowsParams{
		Fromaccountid: FromAccount,
		Toaccountid:   AccountID,
	})

	if isQueue != 1 {
		return nil, errors.New("no queue")
	}

	err = a.postgre.DeleteAccountQueue(ctx, db2.DeleteAccountQueueParams{
		Fromaccountid: FromAccount,
		Toaccountid:   AccountID,
	})
	if err != nil {
		return nil, err
	}

	resultData = map[string]interface{}{
		"status": "deleted",
	}

	return resultData, nil
}
