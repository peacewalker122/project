package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddFollow(t *testing.T) {
	store := Newstore(testDB, config.BucketAccount)
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

	update, err := store.UnFollowtx(ctx, UnfollowTXParam{
		Fromaccid: account2.AccountsID,
		Toaccid:   account1.AccountsID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, update)

	require.Equal(t, UF, update.FeatureType)
	require.True(t, update.Status)

	fromacc = update.FromAcc
	require.Equal(t, account2.AccountsID, fromacc.AccountsID)
	require.Equal(t, int64(0), fromacc.Following)

	toacc = update.ToAcc
	require.Equal(t, account1.AccountsID, toacc.AccountsID)
	require.Equal(t, int64(0), toacc.Follower)
}

func TestIndexingFile(t *testing.T) {
	store := Newstore(testDB, config.BucketAccount)
	filename, err := store.CreateFileIndex("/home/servumtopia/Pictures/Project/1/", "golang.png")
	require.NoError(t, err)

	require.Equal(t, "golang(1)(1).png", filename)
}

func TestMap(t *testing.T) {
	m := make(map[string][]string)
	maps := []string{"babi", "pig", "rw"}
	m["name"] = maps

	log.Print(m["name"][0])

	for _, i := range m["name"] {
		log.Println(i)
	}
}
