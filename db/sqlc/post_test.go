package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomPost(t *testing.T) Post {
	account := CreateRandomAccount(t)
	arg := CreatePostParams{
		AccountID: account.ID,
		PostWord: sql.NullString{
			String: util.Randomstring(20),
			Valid:  true,
		},
		PostPicture: util.Randombyte(),
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	return post
}

func TestCreatePost(t *testing.T) {
	account := CreateRandomAccount(t)
	arg := CreatePostParams{
		AccountID: account.ID,
		PostWord: sql.NullString{
			String: util.Randomstring(20),
			Valid:  true,
		},
		PostPicture: util.Randombyte(),
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.AccountID, post.AccountID)
	require.Equal(t, arg.PostWord.String, post.PostWord.String)
	require.Equal(t, arg.PostPicture, post.PostPicture)
}

func TestGetPost(t *testing.T){
	post := CreateRandomPost(t)
	result,err := testQueries.GetPost(context.Background(),post.ID)
	require.NoError(t,err)
	require.Equal(t,post.ID,result.ID)
	require.Equal(t,post.AccountID,result.AccountID)
	require.Equal(t,post.PostWord.String,result.PostWord.String)
	require.Equal(t,post.PostPicture,result.PostPicture)
}

func TestListPost(t *testing.T){
	for i := 0; i < 10; i++ {
		CreateRandomPost(t)
	}
	arg := ListPostParams{
		Limit:  5,
		Offset: 5,
	}
	result,err := testQueries.ListPost(context.Background(),arg)
	require.NoError(t,err)

	for _,output := range result{
		require.NotEmpty(t,output)
	}
}

func TestUpdatePost(t *testing.T) {
	NewCaption := util.Randomstring(10)
	NewPhoto := util.Randombyte()
	Post := CreateRandomPost(t)
	arg := UpdatePostParams{
		ID:          Post.ID,
		PostWord:    NullString(NewCaption),
		PostPicture: NewPhoto,
	}
	result,err := testQueries.UpdatePost(context.Background(),arg)
	require.NoError(t,err)
	require.Equal(t,Post.ID,result.ID)
	require.Equal(t,Post.AccountID,result.AccountID)
	require.Equal(t,NewCaption,result.PostWord.String)
	require.Equal(t,NewPhoto,result.PostPicture)
}


func NullString(caption string) sql.NullString{
	validity := true
	if len(caption) < 1 {
		validity = false
		return sql.NullString{}
	}
	
	return sql.NullString{
		String: caption,
		Valid:  validity,
	}
}