package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mockdb "github.com/peacewalker122/project/db/mock"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	user, _ := NewUser(t)
	acc := NewAcc(user.Username)
	post := NewPost(int(acc.ID))

	TestCases := []struct {
		name       string
		Body       H
		BuildStubs func(mock *mockdb.MockStore)
		CodeRecord func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			Body: H{
				"account_id":          post.AccountID,
				"picture_description": post.PictureDescription.String,
				"pictureid":           post.PictureID,
			},
			BuildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(post.AccountID)).Times(1).Return(acc, nil)
				arg := db.CreatePostParams{
					AccountID:          acc.ID,
					PictureDescription: post.PictureDescription,
					PictureID:          post.PictureID,
				}

				mock.EXPECT().CreatePost(gomock.Any(), gomock.Eq(arg)).Times(1).Return(post, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
				//BodycheckPost(t, record.Body, post)
			},
		},
	}

	for i := range TestCases {
		tc := TestCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.BuildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/post"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.router.ServeHTTP(recorder, req)
			tc.CodeRecord(recorder)
		})
	}
}

func NewPost(AccID int) db.Post {
	return db.Post{
		ID:        util.Randomint(1, 100),
		AccountID: int64(AccID),
		PictureDescription: sql.NullString{
			String: util.Randomusername(),
			Valid:  true,
		},
		PictureID: util.Randomint(1, 1000),
	}
}

func BodycheckPost(t *testing.T, body *bytes.Buffer, Post db.Post) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotPost db.Post
	err = json.Unmarshal(data, &gotPost)

	require.NoError(t, err)
	require.Equal(t, Post.ID, gotPost.ID)
	require.Equal(t, Post.PictureDescription.String, gotPost.PictureDescription.String)
	require.Equal(t, Post.PictureID, gotPost.PictureID)
	// Just Create The Testing Like The Returning API
}
