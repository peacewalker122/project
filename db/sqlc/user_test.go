package db

import (
	"context"
	"testing"

	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User{
	arg := CreateUserParams{
		Username:       util.Randomusername(),
		HashedPassword: util.Randomstring(6),
		FullName:       util.Randomusername(),
		Email:          util.Randomemail(),
	}

	account,err := testQueries.CreateUser(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,account)

	return account
}

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:       util.Randomusername(),
		HashedPassword: util.Randomstring(6),
		FullName:       util.Randomusername(),
		Email:          util.Randomemail(),
	}

	account,err := testQueries.CreateUser(context.Background(),arg)
	require.NoError(t,err)
	require.Equal(t,arg.Username,account.Username)
	require.Equal(t,arg.HashedPassword,account.HashedPassword)
	require.Equal(t,arg.FullName,account.FullName)
	require.Equal(t,arg.Email,account.Email)
}

func TestGetUser(t *testing.T){
	user := CreateRandomUser(t)

	result,err := testQueries.GetUser(context.Background(),user.Username)
	require.NoError(t,err)
	require.NotEmpty(t,result)
	require.Equal(t,user.Username,result.Username)
	require.Equal(t,user.HashedPassword,result.HashedPassword)
	require.Equal(t,user.FullName,result.FullName)
	require.Equal(t,user.Email,result.Email)
	require.Equal(t,user.PasswordChangedAt,result.PasswordChangedAt)
	require.Equal(t,user.CreatedAt,result.CreatedAt)
}

func TestListUser(t *testing.T){
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}
	arg := ListUserParams{
		Limit:  5,
		Offset: 5,
	}
	result,err := testQueries.ListUser(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,result)

	for _,output := range result{
		require.NotEmpty(t,output)
	} 
}