package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	user := CreateRandomUser(t)
	arg := CreateAccountsParams{
		Owner:     user.Username,
		IsPrivate: false,
	}
	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func TestCreateAccount(t *testing.T) {
	user := CreateRandomUser(t)
	arg := CreateAccountsParams{
		Owner:     user.Username,
		IsPrivate: true,
	}
	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.IsPrivate, account.IsPrivate)
}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	result, err := testQueries.GetAccounts(context.Background(), account.AccountsID)
	require.NoError(t, err)

	require.Equal(t, account.AccountsID, result.AccountsID)
	require.Equal(t, account.Owner, result.Owner)
	require.Equal(t, account.IsPrivate, result.IsPrivate)
	require.Equal(t, account.CreatedAt, result.CreatedAt)
}
func TestGetAccountOwner(t *testing.T) {
	account := CreateRandomAccount(t)

	result, err := testQueries.GetAccountsOwner(context.Background(), account.Owner)
	require.NoError(t, err)

	require.Equal(t, account.AccountsID, result.AccountsID)
	require.Equal(t, account.Owner, result.Owner)
	require.Equal(t, account.IsPrivate, result.IsPrivate)
	require.Equal(t, account.CreatedAt, result.CreatedAt)
}

func TestListAccount(t *testing.T) {
	var acc Account
	for i := 0; i < 10; i++ {
		acc = CreateRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  acc.Owner,
		Limit:  5,
		Offset: 0,
	}

	result, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	for _, output := range result {
		require.NotEmpty(t, output)
		require.Equal(t,acc.Owner,output.Owner)
	}
}
