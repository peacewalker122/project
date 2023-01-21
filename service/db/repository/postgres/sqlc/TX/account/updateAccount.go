package account

import (
	"context"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (a *AccountTx) UpdateFollowing(
	ctx context.Context, fromAccount, toAccount, num int64,
) (
	acc db.AccountsFollow, Toacc, Fromacc db.Account, err error,
) {
	acc, err = a.CreateAccountsFollow(ctx, db.CreateAccountsFollowParams{
		FromAccountID: fromAccount,
		ToAccountID:   toAccount,
		Follow:        true,
	})
	if err != nil {
		return
	}

	Toacc, err = a.UpdateAccountFollower(ctx, db.UpdateAccountFollowerParams{Num: num, ID: toAccount})
	if err != nil {
		return
	}
	Fromacc, err = a.UpdateAccountFollowing(ctx, db.UpdateAccountFollowingParams{Num: num, ID: fromAccount})
	if err != nil {
		return
	}
	return
}

func (a *AccountTx) DeleteFollowing(
	ctx context.Context, fromAccount, toAccount, num int64,
) (
	Toacc, Fromacc db.Account, status bool, err error,
) {
	Toacc, err = a.UpdateAccountFollower(ctx, db.UpdateAccountFollowerParams{Num: num, ID: toAccount})
	if err != nil {
		return
	}
	Fromacc, err = a.UpdateAccountFollowing(ctx, db.UpdateAccountFollowingParams{Num: num, ID: fromAccount})
	if err != nil {
		return
	}

	err = a.DeleteAccountsFollow(ctx, db.DeleteAccountsFollowParams{
		Fromid: fromAccount,
		Toid:   toAccount,
	})
	if err != nil {
		return
	}
	status = true

	return
}
