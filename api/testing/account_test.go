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
	"github.com/lib/pq"
	mockdb "github.com/peacewalker122/project/db/repository/postgres/mock"
	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	user, _ := NewUser(t)
	acc := NewAcc(user.Username)

	TestCases := []struct {
		name      string
		body      H
		setupAuth func(t *testing.T, request *http.Request, token token.Maker)
		stubs     func(mock *mockdb.MockPostgresStore)
		recorder  func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			body: H{"account_type": acc.IsPrivate},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				arg := db.CreateAccountsParams{Owner: acc.Owner, IsPrivate: acc.IsPrivate}
				mock.EXPECT().CreateAccounts(gomock.Any(), gomock.Eq(arg)).Times(1).Return(acc, nil)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
				BodycheckAccount(t, record.Body, acc)
			},
		},
		{
			name: "NonAccount",
			body: H{
				"account_type": acc.IsPrivate,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().
					CreateAccounts(gomock.Any(), gomock.Any()).
					Times(1).Return(db.Account{}, &pq.Error{Code: "23503"})
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, record.Code)
			},
		},
		{
			name: "DuplicateAccount",
			body: H{

				"account_type": acc.IsPrivate,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().
					CreateAccounts(gomock.Any(), gomock.Any()).
					Times(1).Return(db.Account{}, &pq.Error{Code: "23505"})
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, record.Code)
			},
		},
		{
			name: "InternalServerError",
			body: H{

				"account_type": acc.IsPrivate,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().
					CreateAccounts(gomock.Any(), gomock.Any()).
					Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "WrongRequest-account_type",
			body: H{
				"account_type": "acc.IsPrivate",
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().
					CreateAccounts(gomock.Any(), gomock.Any()).
					Times(0)
			},
			recorder: func(record *httptest.ResponseRecorder) {
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
			tc.stubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/account"
			request := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			tc.setupAuth(t, request, server.Token)
			server.Router.ServeHTTP(recorder, request)
			tc.recorder(recorder)
		})
	}
}

func TestGetAccount(t *testing.T) {
	user, _ := NewUser(t)
	acc := NewAcc(user.Username)

	Testcases := []struct {
		name      string
		id        int
		setupAuth func(t *testing.T, request *http.Request, token token.Maker)
		stubs     func(mock *mockdb.MockPostgresStore)
		recorder  func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			id:   int(acc.ID),
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(acc, nil)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, record.Code)
				BodycheckAccount(t, record.Body, acc)
			},
		},
		{
			name: "No-Account",
			id:   int(acc.ID),
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, record.Code)
			},
		},
		{
			name: "Internal-Error",
			id:   int(acc.ID),
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(acc.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name: "InvalidID",
			id:   0,
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().GetAccounts(gomock.Any(), gomock.Eq(acc.ID)).Times(0)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
			},
		},
	}

	for i := range Testcases {
		tc := Testcases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockPostgresStore(ctrl)
			tc.stubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/account/%v", tc.id)
			request := httptest.NewRequest(http.MethodGet, url, nil)
			tc.setupAuth(t, request, server.Token)
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			server.Router.ServeHTTP(recorder, request)
			tc.recorder(recorder)
		})
	}
}

func TestListAccount(t *testing.T) {
	user, _ := NewUser(t)

	n := 5
	acc := make([]db.Account, n)
	for i := 0; i < n; i++ {
		acc[i] = NewAcc(user.Username)
	}

	type Query struct {
		pageID   int
		pageSize int
	}

	testCases := []struct {
		name      string
		query     Query
		setupAuth func(t *testing.T, request *http.Request, token token.Maker)
		stubs     func(mock *mockdb.MockPostgresStore)
		recorder  func(record *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				var account string
				for i := range acc {
					account = acc[i].Owner
				}

				arg := db.ListAccountsParams{
					Owner:  account,
					Limit:  int32(n),
					Offset: 0,
				}
				mock.EXPECT().ListAccounts(gomock.Any(), gomock.Eq(arg)).Times(1).Return(acc, nil)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				log.Println(record.Body)
				require.Equal(t, http.StatusOK, record.Code)
				BodycheckAccounts(t, record.Body, acc)
			},
		},
		{
			name: "Internal-Error",
			query: Query{
				pageID:   1,
				pageSize: n,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {

				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).Times(1).Return([]db.Account{}, sql.ErrConnDone)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, record.Code)
			},
		},
		{
			name:  "Wrong-PageID",
			query: Query{pageID: -1, pageSize: n},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				arg := db.ListAccountsParams{Limit: int32(n), Offset: 0}
				mock.EXPECT().ListAccounts(gomock.Any(), gomock.Eq(arg)).Times(0)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, record.Code)
			},
		},
		{
			name: "Wrong-PageSize",
			query: Query{
				pageID:   1,
				pageSize: 1,
			},
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, user.Username, AuthTypeBearer, time.Minute)
			},
			stubs: func(mock *mockdb.MockPostgresStore) {
				mock.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).Times(0)
			},
			recorder: func(record *httptest.ResponseRecorder) {
				log.Println(record.Body)
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
			tc.stubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			urls := "/account"
			req := httptest.NewRequest(http.MethodGet, urls, nil)
			tc.setupAuth(t, req, server.Token)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			q := req.URL.Query()
			q.Add("page_id", fmt.Sprintf("%d", tc.query.pageID))
			q.Add("page_size", fmt.Sprintf("%d", tc.query.pageSize))
			req.URL.RawQuery = q.Encode()

			server.Router.ServeHTTP(recorder, req)
			tc.recorder(recorder)
		})
	}
}

func NewAcc(random string) db.Account {
	return db.Account{
		ID:        util.Randomint(1, 1000),
		Owner:     random,
		IsPrivate: false,
	}
}

func BodycheckAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)

	require.NoError(t, err)
	require.Equal(t, account.Owner, gotAccount.Owner)
	require.Equal(t, account.IsPrivate, gotAccount.IsPrivate)
	// Just Create The Testing Like The Returning API
}

func BodycheckAccounts(t *testing.T, body *bytes.Buffer, account []db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount []db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}
