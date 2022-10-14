package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	mockdb "github.com/peacewalker122/project/db/mock"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

type EqMatcherPass struct {
	user db.CreateUserParams
	pass string
}

func (e EqMatcherPass) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.pass, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.user.HashedPassword = arg.HashedPassword

	return reflect.DeepEqual(e.user, arg)
}

func (e EqMatcherPass) String() string {
	return fmt.Sprintf("matches arg %v amd password %v", e.user, e.pass)
}

func Eq(user db.CreateUserParams, pass string) gomock.Matcher {
	return EqMatcherPass{
		user: user,
		pass: pass,
	}
}

func TestCreateUser(t *testing.T) {
	user, pass := NewUser(t)

	testCases := []struct {
		name          string
		Body          H
		buildstubs    func(mockdb *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			Body: H{
				"username":  user.Username,
				"password":  pass,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildstubs: func(mockdb *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
					HashedPassword: user.HashedPassword,
				}
				mockdb.EXPECT().
					CreateUser(gomock.Any(), Eq(arg, pass)).Times(1).Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				BodycheckUser(t, recorder.Body, user)
			},
		},
		{
			name: "DuplicateUsername",
			Body: H{
				"username":  user.Username,
				"password":  pass,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildstubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidUsername",
			Body: H{
				"username":  "test123*",
				"password":  pass,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildstubs: func(mock *mockdb.MockStore) {
				//stubs
				mock.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "Internal Error",
			Body: H{
				"username":  user.Username,
				"password":  pass,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildstubs: func(mockdb *mockdb.MockStore) {
				mockdb.EXPECT().CreateUser(gomock.Any(), gomock.Any()).
					Times(1).Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildstubs(store)

			server := Newserver(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.Body)
			require.NoError(t, err)

			url := "/user"
			request := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			//without this code above, your testing won't work.

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}

}

func NewUser(t *testing.T) (user db.User, password string) {
	password = util.Randomstring(6)
	hash, err := util.HashPassword(password)
	require.NoError(t, err)
	user = db.User{
		Username:       util.Randomusername(),
		HashedPassword: hash,
		FullName:       util.Randomusername(),
		Email:          util.Randomemail(),
	}
	return
}

func BodycheckUser(t *testing.T, body *bytes.Buffer, account db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, account.Username, gotUser.Username)
	require.Equal(t, account.FullName, gotUser.FullName)
	require.Equal(t, account.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
