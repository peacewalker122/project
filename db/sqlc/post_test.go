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
		PictureDescription: sql.NullString{
			String: util.Randomstring(20),
			Valid:  true,
		},
		PictureID: util.Randomint(1,1000),
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
		PictureDescription: sql.NullString{
			String: util.Randomstring(20),
			Valid:  true,
		},
		PictureID: util.Randomint(1,1000),
	}
	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.AccountID, post.AccountID)
	require.Equal(t, arg.PictureDescription.String, post.PictureDescription.String)
	require.Equal(t, arg.PictureID, post.PictureID)
}

func TestGetPost(t *testing.T){
	post := CreateRandomPost(t)
	result,err := testQueries.GetPost(context.Background(),post.ID)
	require.NoError(t,err)
	require.Equal(t,post.ID,result.ID)
	require.Equal(t,post.AccountID,result.AccountID)
	require.Equal(t,post.PictureDescription.String,result.PictureDescription.String)
	require.Equal(t,post.PictureID,result.PictureID)
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
	NewPhoto := util.Randomint(1,1000)
	Post := CreateRandomPost(t)
	arg := UpdatePostParams{
		ID:          Post.ID,
		PictureDescription:    NullString(NewCaption),
		PictureID: NewPhoto,
	}
	result,err := testQueries.UpdatePost(context.Background(),arg)
	require.NoError(t,err)
	require.Equal(t,Post.ID,result.ID)
	require.Equal(t,Post.AccountID,result.AccountID)
	require.Equal(t,NewCaption,result.PictureDescription.String)
	require.Equal(t,NewPhoto,result.PictureID)
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