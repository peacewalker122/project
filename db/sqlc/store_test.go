package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddFollow(t *testing.T) {
	store := Newstore(testDB)
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	ctx := context.Background()
	results, err := store.Followtx(ctx, FollowTXParam{
		Fromaccid: account2.AccountsID,
		Toaccid:   account1.AccountsID,
	})
	require.NoError(t, err)

	res := results
	require.NotEmpty(t, res)

	require.Equal(t, F, res.FeatureType)

	follow := res.Follow
	require.Equal(t, account2.AccountsID, follow.FromAccountID)
	require.Equal(t, account1.AccountsID, follow.ToAccountID)
	require.Equal(t, true, follow.Follow)

	fromacc := res.FromAcc
	require.Equal(t, account2.AccountsID, fromacc.AccountsID)
	require.Equal(t, int64(1), fromacc.Following)

	toacc := res.ToAcc
	require.Equal(t, account1.AccountsID, toacc.AccountsID)
	require.Equal(t, int64(1), toacc.Follower)
}

func BenchmarkFollow(b *testing.B) {
	ctx := context.Background()
	store := Newstore(testDB)
	for i := 0; i < 100; i++ {
		store.Followtx(ctx, FollowTXParam{
			Fromaccid: int64(2),
			Toaccid:   int64(1),
		})
	}
}
