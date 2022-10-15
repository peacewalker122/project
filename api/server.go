package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/peacewalker122/project/db/sqlc"
)

type Server struct {
	store  db.Store
	router *echo.Echo
}

func Newserver(store db.Store) *Server {
	server := &Server{store: store}
	server.routerhandle()
	return server
}

func (s *Server) routerhandle() {
	router := echo.New()
	router.Use(middleware.LoggerWithConfig(Logger()))
	router.Validator = &customValidator{
		validate: validator.New(),
	}
	router.HTTPErrorHandler = HTTPErrorHandler

	router.POST("/user", s.createUser)
	router.POST("/account", s.createAccount)
	router.GET("/account/:id", s.getAccounts)
	router.GET("/account", s.listAccount)
	router.POST("/post", s.createPost)
	router.GET("/post/:id", s.getPost)

	s.router = router
}

func (s *Server) Start(path string) error {
	return s.router.Start(path)
}
