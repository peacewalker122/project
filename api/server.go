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
	router.Use(middleware.TimeoutWithConfig(s.Timeout()))
	//router.Use(middleware.HTTPSRedirectWithConfig(Redirect()))
	router.Validator = &customValidator{
		validate: validator.New(),
	}
	router.HTTPErrorHandler = HTTPErrorHandler

	router.POST("/user", s.createUser)
	router.POST("/token/renew", s.renewToken)
	router.POST("/user/login", s.login)

	authRouter := router.Group("", authMiddleware(s.token))
	//authRouter.POST("/account", s.createAccount)
	authRouter.GET("/account/:id", s.getAccounts)
	authRouter.GET("/account", s.listAccount)
	authRouter.POST("/account/follow", s.followAccount)
	authRouter.POST("/post", s.createPost)
	authRouter.GET("/post/:id", s.getPost)
	authRouter.POST("/post/like", s.likePost)
	authRouter.POST("/post/comment", s.commentPost)
	authRouter.POST("/post/retweet", s.retweetPost)
	authRouter.GET("/post/image/:id", s.getPostImage, middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	authRouter.POST("/post/qoute/retweet", s.qouteretweetPost)

	s.router = router
}

func (s *Server) StartHTTPS(path string) error {
	return s.router.StartAutoTLS(path)
}
func (s *Server) StartHTTP(path string) error {
	return s.router.Start(path)
}
func (s *Server) timeout(c echo.Context) error {
	return c.JSON(echo.ErrRequestTimeout.Code, "timeout")
}
