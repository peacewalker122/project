package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomPostFeature(t *testing.T) PostFeature {
	post := CreateRandomPost(t)
	result, err := testQueries.CreatePost_feature(context.Background(), post.PostID)
	require.NoError(t, err)
	require.Equal(t, post.PostID, result.PostID)

	return result
}

func TestCreatePostFeature(t *testing.T) {
	post := CreateRandomPost(t)
	returning := PostFeature{
		PostID:          post.PostID,
		SumComment:      0,
		SumLike:         0,
		SumRetweet:      0,
		SumQouteRetweet: 0,
	}
	require.Equal(t, post.PostID, returning.PostID)
}

func TestGetPostFeature(t *testing.T) {
	post := CreateRandomPostFeature(t)

	result, err := testQueries.GetPost_feature(context.Background(), post.PostID)
	require.NoError(t, err)
	require.Equal(t, post.PostID, result.PostID)
	require.Equal(t, post.SumComment, result.SumComment)
	require.Equal(t, post.SumRetweet, result.SumRetweet)

	results, err := testQueries.GetPost_feature_Update(context.Background(), post.PostID)
	require.NoError(t, err)
	//require.Equal(t, post.PostID, results.PostID)
	require.Equal(t, post.SumComment, results.SumComment)
	require.Equal(t, post.SumRetweet, results.SumRetweet)

	update, err := testQueries.UpdatePost_feature(context.Background(), UpdatePost_featureParams{
		PostID:          post.PostID,
		SumComment:      result.SumComment + 1,
		SumLike:         result.SumLike + 1,
		SumRetweet:      result.SumRetweet + 1,
		SumQouteRetweet: result.SumQouteRetweet + 1,
	})

	require.NoError(t, err)
	require.Equal(t, post.PostID, update.PostID)
	require.Equal(t, int64(1), update.SumComment)
	require.Equal(t, int64(1), update.SumRetweet)
}

func TestJoinTable(t *testing.T) {
	join, err := testQueries.GetPostJoin(context.Background(), int64(69))
	require.NoError(t, err)
	require.Equal(t, int64(143), join.AccountID)
}
