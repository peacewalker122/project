package account

import (
	"context"
	api "github.com/peacewalker122/project/api/handler"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
)

func (a *AccountUseCase) PrivateAccount(ctx context.Context, AccountID int64) (api.BasicResponse, error) {
	var resultData map[string]interface{}
	var accStatus string
	acc, err := a.postgre.GetAccounts(ctx, AccountID)
	if err != nil {
		return nil, err
	}

	err = a.postgre.PrivateAccount(ctx, db2.PrivateAccountParams{
		IsPrivate: !acc.IsPrivate,
		Username:  acc.Owner,
	})
	if err != nil {
		return nil, err
	}

	switch acc.IsPrivate {
	case true:
		accStatus = "private"
	case false:
		accStatus = "public"
	}

	resultData = map[string]interface{}{
		"status": accStatus,
	}

	return resultData, nil
}
