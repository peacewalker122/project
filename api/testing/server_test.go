package api

import (
	"testing"
	"time"

	router "github.com/peacewalker122/project/api/router"
	db "github.com/peacewalker122/project/db/sqlc"
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

func NewTestServer(t *testing.T, store db.Store) *router.Server {
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
