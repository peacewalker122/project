package account

import (
	"context"
	"errors"
	api "github.com/peacewalker122/project/api/handler"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
)

func (a *AccountUseCase) UnFollowAccount(ctx context.Context, FromAccount, AccountID int64) (api.BasicResponse, error) {

	var resultData map[string]interface{}

	num, err := a.postgre.GetAccountsFollowRows(ctx, db2.GetAccountsFollowRowsParams{
		Fromid: FromAccount,
		Toid:   AccountID,
	})
	if err != nil {
		return nil, err
	}

	if num == 0 {
		return nil, errors.New("not requested")
	}

	_, err = a.postgre.UnFollowtx(ctx, db2.UnfollowTXParam{
		Fromaccid: FromAccount,
		Toaccid:   AccountID,
	})
	if err != nil {
		return nil, err
	}

	resultData = map[string]interface{}{
		"status": "unfollowed",
	}

	return resultData, nil
}
