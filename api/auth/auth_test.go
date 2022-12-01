package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AddAuthorization(
	t *testing.T,
	req *http.Request,
	tokenMaker token.Maker,
	username string,
	AuthType string,
	Duration time.Duration,
) {
	token, payload, err := tokenMaker.CreateToken(username, Duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	AuthHeader := fmt.Sprintf("%s %s", AuthType, token)
	assert.NoError(t, err)
	req.Header.Set(authHeaderkey, AuthHeader)
}
func TestAuth(t *testing.T) {

	testCases := []struct {
		name      string
		setupAuth func(t *testing.T, request *http.Request, token token.Maker)
		recorder  func(t *testing.T, rec *httptest.ResponseRecorder)
	}{
		{
			name: "Ok",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, "test", authTypeBearer, time.Minute)
			},
			recorder: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			name: "UnsupportedAuth",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, "test", "authTypeBearer", time.Minute)
			},
			recorder: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
		{
			name: "Invalid-Auth",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, "", authTypeBearer, time.Minute)
			},
			recorder: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, rec.Code)
			},
		},
		{
			name: "Expired-Token",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, "test", authTypeBearer, -time.Minute)
			},
			recorder: func(t *testing.T, rec *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, rec.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := NewTestServer(t, nil)
			url := "/auth"

			server.router.GET(url, func(c echo.Context) error {
				return c.JSON(http.StatusOK, echo.Map{})
			}, authMiddleware(server.token))

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, url, nil)
			tc.setupAuth(t, req, server.token)
			server.router.ServeHTTP(recorder, req)
			tc.recorder(t, recorder)

		})
	}
}
