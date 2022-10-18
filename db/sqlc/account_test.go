package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	user := CreateRandomUser(t)
	arg := CreateAccountsParams{
		Owner:       user.Username,
		AccountType: true,
	}
	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func TestCreateAccount(t *testing.T) {
	user := CreateRandomUser(t)
	arg := CreateAccountsParams{
		Owner:       user.Username,
		AccountType: true,
	}
	account, err := testQueries.CreateAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.AccountType, account.AccountType)
}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	result, err := testQueries.GetAccounts(context.Background(), account.ID)
	require.NoError(t, err)

	require.Equal(t, account.ID, result.ID)
	require.Equal(t, account.Owner, result.Owner)
	require.Equal(t, account.AccountType, result.AccountType)
	require.Equal(t, account.CreatedAt, result.CreatedAt)
}
func TestGetAccountOwner(t *testing.T) {
	account := CreateRandomAccount(t)

	result, err := testQueries.GetAccountsOwner(context.Background(), account.Owner)
	require.NoError(t, err)

	require.Equal(t, account.ID, result.ID)
	require.Equal(t, account.Owner, result.Owner)
	require.Equal(t, account.AccountType, result.AccountType)
	require.Equal(t, account.CreatedAt, result.CreatedAt)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	result, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	for _, output := range result {
		require.NotEmpty(t, output)
	}
}
