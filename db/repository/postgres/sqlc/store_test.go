package db

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func TestAddFollow(t *testing.T) {
	store := Newstore(testDB)
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	var res FollowTXResult
	var err error

	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err = store.Followtx(ctx, FollowTXParam{
			Fromaccid: account2.ID,
			Toaccid:   account1.ID,
		})
	}()
	wg.Wait()
	require.NoError(t, err)

	require.NotEmpty(t, res)

	require.Equal(t, F, res.FeatureType)

	follow := res.Follow
	require.Equal(t, account2.ID, follow.FromAccountID)
	require.Equal(t, account1.ID, follow.ToAccountID)
	require.Equal(t, true, follow.Follow)

	fromacc := res.FromAcc
	require.Equal(t, account2.ID, fromacc.ID)
	require.Equal(t, int64(1), fromacc.Following)

	toacc := res.ToAcc
	require.Equal(t, account1.ID, toacc.ID)
	require.Equal(t, int64(1), toacc.Follower)

	update, err := store.UnFollowtx(ctx, UnfollowTXParam{
		Fromaccid: account2.ID,
		Toaccid:   account1.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, update)

	require.Equal(t, UF, update.FeatureType)
	require.True(t, update.Status)

	fromacc = update.FromAcc
	require.Equal(t, account2.ID, fromacc.ID)
	require.Equal(t, int64(0), fromacc.Following)

	toacc = update.ToAcc
	require.Equal(t, account1.ID, toacc.ID)
	require.Equal(t, int64(0), toacc.Follower)
}

func TestCreatePostTX(t *testing.T) {
	store := Newstore(testDB)
	account := CreateRandomAccount(t)

	arg := CreatePostParams{
		AccountID:          account.ID,
		IsRetweet:          false,
		PictureDescription: "memento mori",
		PhotoDir:           util.InputSqlString(""),
	}

	Post, err := store.CreatePostTx(context.Background(), arg)

	require.NoError(t, err)

	require.Equal(t, account.ID, Post.Post.AccountID)
	require.Equal(t, "memento mori", Post.Post.PictureDescription)

	postFeat := Post.PostFeature
	require.Equal(t, int64(0), postFeat.SumComment)
	require.Equal(t, int64(0), postFeat.SumLike)
	require.Equal(t, int64(0), postFeat.SumRetweet)
	require.Equal(t, int64(0), postFeat.SumQouteRetweet)
}

func TestIndexingFile(t *testing.T) {
	store := Newstore(testDB)
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

func TestCreateRetweet(t *testing.T) {
	s := newTeststore(testDB)
	res, err := s.CreateRetweetTX(context.Background(), CreateRetweetParams{
		FromAccountID: 1,
		PostID:        2,
		IsRetweet:     true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)
}
