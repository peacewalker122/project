package api

import (
	db "github.com/peacewalker122/project/service/db/repository/postgres"
	"testing"
	"time"

	router "github.com/peacewalker122/project/api/router"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

var (
	AuthRefresh    = "RefreshToken"
	AuthHeaderkey  = "authorization"
	AuthTypeBearer = "bearer"
	AuthPayload    = "authorization_payload"
)

func NewTestServer(t *testing.T, store db.PostgresStore) *router.Server {
	config := util.Config{
		TokenKey:      util.Randomstring(32),
		TokenDuration: time.Minute,
	}
	token, err := token.NewJwt(config.TokenKey)
	require.NoError(t, err)
	server := &router.Server{
		Token:  token,
		Config: config,
		Store:  store,
	}

	server.Testrouterhandle()
	return server
}
