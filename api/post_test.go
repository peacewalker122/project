package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mockdb "github.com/peacewalker122/project/db/mock"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	user, _ := NewUser(t)
	acc := NewAcc(user.Username)
	post := NewPost(int(acc.AccountsID))

	TestCases := []struct {
		name       string
		Body       H
		setupAuth  func(t *testing.T, request *http.Request, token token.Maker)
		BuildStubs func(mock *mockdb.MockStore)
		CodeRecord func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			Body: H{
				"account_id":          post.AccountID,
				"picture_description": post.PictureDescription,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(post.AccountID)).Times(1).Return(acc, nil)
				arg := db.CreatePostParams{
					AccountID:          acc.AccountsID,
					PictureDescription: post.PictureDescription,
				}

				mock.EXPECT().CreatePost(gomock.Any(), gomock.Eq(arg)).Times(1).Return(post, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
			},
		},
		{
			name: "Internal-Error(CreatePost)",
			Body: H{
				"account_id":          post.AccountID,
				"picture_description": post.PictureDescription,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(post.AccountID)).Times(1).Return(acc, nil)
				arg := db.CreatePostParams{
					AccountID:          acc.AccountsID,
					PictureDescription: post.PictureDescription,
				}

				mock.EXPECT().CreatePost(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Post{}, sql.ErrConnDone)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "Wrong-Request(account_id)",
			Body: H{
				"account_id":          "post.AccountID",
				"picture_description": post.PictureDescription,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Any()).Times(0)

			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
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
			tc.setupAuth(t, req, server.token)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.router.ServeHTTP(recorder, req)
			tc.CodeRecord(recorder)
		})
	}
}

func TestGetPost(t *testing.T) {
	user, _ := NewUser(t)
	user2, _ := NewUser(t)
	acc := NewAcc(user.Username)
	acc2 := NewAcc(user2.Username)
	post := NewPost(int(acc.AccountsID))

	testCases := []struct {
		name       string
		id         int64
		setupAuth  func(t *testing.T, request *http.Request, token token.Maker)
		buildStubs func(mock *mockdb.MockStore)
		CodeRecord func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			id:   post.PostID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.PostID)).Times(1).Return(post, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
				log.Println(record.Code)
			},
		},
		{
			name: "Unauthorized-Username",
			id:   post.PostID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user2.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user2.Username)).Times(1).Return(acc2, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.PostID)).Times(1).Return(post, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, record.Code)
			},
		},
		{
			name: "Nonexist-Username",
			id:   post.PostID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Any()).Times(1).Return(db.Account{}, sql.ErrNoRows)
				//mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(post, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "Internal-Error",
			id:   post.PostID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.PostID)).Times(1).Return(db.Post{}, sql.ErrConnDone)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "Internal-Error(GetOwner)",
			id:   post.PostID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "No-Post",
			id:   post.PostID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.PostID)).Times(1).Return(db.Post{}, sql.ErrNoRows)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "Wrong-Request",
			id:   0,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, authTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(0)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.PostID)).Times(0)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/post/%v", tc.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			tc.setupAuth(t, req, server.token)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.router.ServeHTTP(recorder, req)
			tc.CodeRecord(recorder)
		})
	}
}

func NewPost(AccID int) db.Post {
	return db.Post{
		PostID:             util.Randomint(1, 100),
		AccountID:          int64(AccID),
		PictureDescription: util.Randomusername(),
	}
}

func BodycheckPost(t *testing.T, body *bytes.Buffer, Post db.Post) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotPost db.Post
	err = json.Unmarshal(data, &gotPost)

	require.NoError(t, err)
	require.Equal(t, Post.PostID, gotPost.PostID)
	require.Equal(t, Post.PictureDescription, gotPost.PictureDescription)
	// Just Create The Testing Like The Returning API
}
