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
	router.HTTPErrorHandler = HTTPErrorHandler
	// router.Use(middleware.LoggerWithConfig(Logger()))
	// router.Binder = new(CustomBinder)
	router.POST("/user", s.createUser)

	authRouter := router.Group("/", authMiddleware(s.token))

	authRouter.POST("/account", s.createAccount)
	authRouter.GET("/account/:id", s.getAccounts)
	authRouter.GET("/account", s.listAccount)
	authRouter.POST("/post", s.createPost)
	authRouter.GET("/post/:id", s.getPost)

	s.router = router
}
