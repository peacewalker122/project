package account

import (
	"context"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc"
)

func (a *AccountUseCase) GetAccount(ctx context.Context, AccountID int64) (*db2.Account, error) {
	accountData, err := a.postgre.GetAccounts(ctx, AccountID)
	if err != nil {
		return nil, err
	}

	return &accountData, nil
}

func (a *AccountUseCase) ListQueuedAccount(ctx context.Context, param *GetAccountParams) (*[]db2.ListQueueRow, error) {
	accounts, err := a.postgre.ListQueue(ctx, db2.ListQueueParams{
		Limit:     param.Limit,
		Offset:    (param.Offset - 1) * param.Limit,
		Accountid: param.ToAccountID,
	})
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}

func (a *AccountUseCase) ListAccount(ctx context.Context, param *GetAccountParams) (*[]db2.Account, error) {
	accounts, err := a.postgre.ListAccounts(ctx, db2.ListAccountsParams{
		Limit:  param.Limit,
		Offset: (param.Offset - 1) * param.Limit,
		Owner:  param.Username,
	})
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}
