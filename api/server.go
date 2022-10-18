package api

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

type Server struct {
	config util.Config
	store  db.Store
	router *echo.Echo
	token  token.Maker
}

func Newserver(c util.Config, store db.Store) (*Server, error) {
	newtoken, err := token.NewJwt(c.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token %v", err.Error())
	}
	server := &Server{
		config: c,
		store:  store,
		token:  newtoken,
	}
	server.routerhandle()
	return server, nil
}

func (s *Server) routerhandle() {
	router := echo.New()
	router.Use(middleware.LoggerWithConfig(Logger()))
	router.Validator = &customValidator{
		validate: validator.New(),
	}
	router.HTTPErrorHandler = HTTPErrorHandler
	router.POST("/user", s.createUser)
	router.POST("/token/renew", s.renewToken)
	router.POST("/user/login", s.login)

	authRouter := router.Group("", authMiddleware(s.token))

	authRouter.POST("/account", s.createAccount)
	authRouter.GET("/account/:id", s.getAccounts)
	authRouter.GET("/account", s.listAccount)
	authRouter.POST("/post", s.createPost)
	authRouter.GET("/post/:id", s.getPost)

	s.router = router
}

func (s *Server) Start(path string) error {
	return s.router.Start(path)
}
