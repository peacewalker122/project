package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
)

type Server struct {
	store *db.Store  
	router *echo.Echo
}

func Newserver(store *db.Store) *Server{
	server := &Server{store: store}
	router := echo.New()
	router.Validator = &customValidator{
		validate: validator.New(),
	}
	router.HTTPErrorHandler = HTTPErrorHandler

	router.POST("/user",server.createUser)
	router.POST("/account",server.createAccount)

	server.router = router
	return server
}

func(s *Server) Start(path string) error{
	return s.router.Start(path)
}