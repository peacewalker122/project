package db

import (
	"context"
	"testing"

	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

const (
	L  = "like"
	R  = "retweet"
	C  = "comment"
	QR = "qoute-retweet"
)

func createRandomEntry(t *testing.T) Entry {
	acc := CreateRandomAccount(t)
	post := CreateRandomPost(t)
	entry := CreateEntriesParams{
		FromAccountID: acc.AccountsID,
		PostID:        post.PostID,
		TypeEntries:   util.RandomType(),
	}
	result, err := testQueries.CreateEntries(context.Background(), entry)
	require.NoError(t, err)
	require.Equal(t, acc.AccountsID, result.FromAccountID)
	require.Equal(t, result.PostID, post.PostID)
	require.Equal(t, result.TypeEntries, entry.TypeEntries)

	return result
}

func TestCreateEntry(t *testing.T) {
	acc := CreateRandomAccount(t)
	post := CreateRandomPost(t)
	entry := CreateEntriesParams{
		FromAccountID: acc.AccountsID,
		PostID:        post.PostID,
		TypeEntries:   util.RandomType(),
	}
	result, err := testQueries.CreateEntries(context.Background(), entry)
	require.NoError(t, err)
	require.Equal(t, acc.AccountsID, result.FromAccountID)
	require.Equal(t, result.PostID, post.PostID)
	require.Equal(t, result.TypeEntries, entry.TypeEntries)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)

	result, err := testQueries.GetEntries(context.Background(), entry.EntriesID)
	require.NoError(t, err)
	require.Equal(t, entry.FromAccountID, result.FromAccountID)
	require.Equal(t, result.PostID, entry.PostID)
	require.Equal(t, result.TypeEntries, entry.TypeEntries)
}

func TestListEntry(t *testing.T) {
	var entry Entry
	for i := 0; i <= 5; i++ {
		entry = createRandomEntry(t)
	}
	arg := ListEntriesParams{
		PostID: entry.PostID,
		Limit:  5,
		Offset: 0,
	}
	result, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	for _, output := range result {
		require.NotEmpty(t, output)
		require.Equal(t, entry.PostID, output.PostID)
	}
}
