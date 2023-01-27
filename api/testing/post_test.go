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

	"github.com/google/uuid"
	mockdb "github.com/peacewalker122/project/service/db/repository/postgres/mock"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePost(t *testing.T) {
	user, _ := NewUser(t)
	acc := NewAcc(user.Username)
	post := NewPost(int(acc.ID))
	postfeature := NewPostFeature(post.ID)

	TestCases := []struct {
		name       string
		Body       H
		setupAuth  func(t *testing.T, request *http.Request, token token.Maker)
		BuildStubs func(mock *mockdb.MockPostgresStore)
		CodeRecord func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			Body: H{
				"account_id":          post.AccountID,
				"picture_description": post.PictureDescription,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(post.AccountID)).Times(1).Return(acc, nil)
				arg := db2.CreatePostParams{
					AccountID:          acc.ID,
					PictureDescription: post.PictureDescription,
				}

				mock.EXPECT().CreatePost(gomock.Any(), gomock.Eq(arg)).Times(1).Return(post, nil)
				mock.EXPECT().CreatePost_feature(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(postfeature, nil)
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
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(post.AccountID)).Times(1).Return(acc, nil)

				mock.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Times(1).Return(db2.Post{}, sql.ErrConnDone)
				mock.EXPECT().CreatePost_feature(gomock.Any(), gomock.Any()).Times(0)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "Internal-Error(CreatePostFeature)",
			Body: H{
				"account_id":          post.AccountID,
				"picture_description": post.PictureDescription,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(post.AccountID)).Times(1).Return(acc, nil)

				mock.EXPECT().CreatePost(gomock.Any(), gomock.Any()).Times(1).Return(post, nil)
				mock.EXPECT().CreatePost_feature(gomock.Any(), gomock.Any()).Times(1).Return(db2.PostFeature{}, sql.ErrConnDone)
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
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			BuildStubs: func(mock *mockdb.MockPostgresStore) {
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

			store := mockdb.NewMockPostgresStore(ctrl)
			tc.BuildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/post"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			tc.setupAuth(t, req, server.Token)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.Router.ServeHTTP(recorder, req)
			tc.CodeRecord(recorder)
		})
	}
}

func TestGetPost(t *testing.T) {
	user, _ := NewUser(t)
	user2, _ := NewUser(t)
	acc := NewAcc(user.Username)
	acc2 := NewAcc(user2.Username)
	post := NewPost(int(acc.ID))
	postfeature := NewPostFeature(post.ID)
	testCases := []struct {
		name       string
		id         uuid.UUID
		setupAuth  func(t *testing.T, request *http.Request, token token.Maker)
		buildStubs func(mock *mockdb.MockPostgresStore)
		CodeRecord func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(post, nil)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(postfeature, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
				log.Println(record.Code)
			},
		},
		{
			name: "Unauthorized-Username",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user2.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user2.Username)).Times(1).Return(acc2, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(post, nil)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(postfeature, nil)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, record.Code)
			},
		},
		{
			name: "Nonexist-Username",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Any()).Times(1).Return(db2.Account{}, sql.ErrNoRows)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Any()).Times(0)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Any()).Times(0)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "Internal-Error(Get-Post)",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(db2.Post{}, sql.ErrConnDone)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Any()).Times(0)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "Internal-Error(Get-Post_feature)",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(post, nil)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Any()).Times(1).Return(db2.PostFeature{}, sql.ErrConnDone)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "Internal-Error(GetOwner)",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(db2.Account{}, sql.ErrNoRows)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "No-Post(Post)",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(db2.Post{}, sql.ErrNoRows)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Any()).Times(0)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "No-Post(Post-feature)",
			id:   post.ID,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(1).Return(acc, nil)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(1).Return(post, nil)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Any()).Times(1).Return(db2.PostFeature{}, sql.ErrNoRows)
			},
			CodeRecord: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "Wrong-Request",
			id:   uuid.Nil,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			buildStubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccountsOwner(gomock.Any(), gomock.Eq(user.Username)).Times(0)
				mock.EXPECT().GetPost(gomock.Any(), gomock.Eq(post.ID)).Times(0)
				mock.EXPECT().GetPost_feature(gomock.Any(), gomock.Any()).Times(0)
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

			store := mockdb.NewMockPostgresStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/post/%v", tc.id)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			tc.setupAuth(t, req, server.Token)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.Router.ServeHTTP(recorder, req)
			tc.CodeRecord(recorder)
		})
	}
}

func NewPost(AccID int) db2.Post {
	return db2.Post{
		ID:                 uuid.New(),
		AccountID:          int64(AccID),
		PictureDescription: util.Randomusername(),
	}
}

func NewPostFeature(UUID uuid.UUID) db2.PostFeature {
	return db2.PostFeature{
		PostID:          UUID,
		SumComment:      util.Randomint(1, 100),
		SumLike:         util.Randomint(1, 1000),
		SumRetweet:      util.Randomint(1, 100),
		SumQouteRetweet: util.Randomint(1, 100),
	}
}

func BodycheckPost(t *testing.T, body *bytes.Buffer, Post db2.Post) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotPost db2.Post
	err = json.Unmarshal(data, &gotPost)

	require.NoError(t, err)
	require.Equal(t, Post.ID, gotPost.ID)
	require.Equal(t, Post.PictureDescription, gotPost.PictureDescription)
	// Just Create The Testing Like The Returning API
}
