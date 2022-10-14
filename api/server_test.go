package api

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
)

func NewTestServer(t *testing.T, store db.Store) *Server {
	server := &Server{store: store}
	server.testrouterhandle()
	return server
}

func (s *Server) testrouterhandle() {
	router := echo.New()
	router.Validator = &customValidator{
		validate: validator.New(),
	}
	router.HTTPErrorHandler = HTTPErrorHandler
	router.Binder = new(CustomBinder)

	router.POST("/user", s.createUser)
	router.POST("/account", s.createAccount)
	router.GET("/account/:id", s.getAccounts)
	router.GET("/account", s.listAccount)
	router.POST("/post", s.createPost)

	s.router = router
}
