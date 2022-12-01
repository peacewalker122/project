package api

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenKey:      util.Randomstring(32),
		TokenDuration: time.Minute,
	}
	token, err := token.NewJwt(config.TokenKey)
	require.NoError(t, err)
	server := &Server{
		config: config,
		store:  store,
		token:  token,
	}

	server.testrouterhandle()
	return server
}

func (s *Server) testrouterhandle() {
	router := echo.New()
	router.Validator = &customValidator{
		validate: validator.New(),
	}
	// router.HTTPErrorHandler = HTTPErrorHandler
	// router.Use(middleware.LoggerWithConfig(Logger()))
	// router.Binder = new(CustomBinder)
	router.POST("/user", s.createUser)

	AuthMethod := router.Group("", authMiddleware(s.token))

	//AuthMethod.POST("/account", s.createAccount)
	AuthMethod.GET("/account/:id", s.getAccounts)
	AuthMethod.GET("/account", s.listAccount)
	AuthMethod.POST("/post", s.createPost)
	AuthMethod.GET("/post/:id", s.getPost)

	s.router = router
}
