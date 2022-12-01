package api

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "github.com/peacewalker122/project/api/handler"
	apiUtil "github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

type Server struct {
	handler    *handler.Handler
	router     *echo.Echo
	token      token.Maker
	fileString string
}

func Newserver(c util.Config, store db.Store, redisStore redis.Store) (*Server, error) {
	newtoken, err := token.NewJwt(c.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token %v", err.Error())
	}
	newHandler := handler.NewHandler(store, redisStore, c, newtoken)
	server := &Server{
		handler: newHandler,
		token:   newtoken,
	}
	server.routerhandle()
	return server, nil
}

func (s *Server) routerhandle() {
	router := echo.New()
	router.Use(middleware.LoggerWithConfig(apiUtil.Logger()))

	//router.Use(middleware.HTTPSRedirectWithConfig(Redirect()))
	router.Validator = apiUtil.NewValidator(validator.New())
	router.HTTPErrorHandler = apiUtil.HTTPErrorHandler

	router.POST("/user", s.createUser)
	router.POST("/token/renew", s.renewToken)
	router.POST("/user/login", s.login)

	authRouter := router.Group("", authMiddleware(s.token))
	//authRouter.POST("/account", s.createAccount)
	authRouter.GET("/account/:id", s.getAccounts)
	authRouter.GET("/account", s.listAccount)
	authRouter.POST("/account/follow", s.followAccount)
	authRouter.PUT("/account/follow", s.acceptFollower)
	authRouter.POST("/post", s.createPost, middleware.TimeoutWithConfig(s.TimeoutPost()))
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
func (s *Server) TimeoutPost() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		ErrorMessage: "timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			// we delete the file if its already timeout
			if _, err := os.Stat(s.fileString); err == nil {
				os.Remove(s.fileString)
			}
			c.Error(err)
			c.SetHandler(s.timeout)
		},
		Timeout: 8 * time.Second,
	}
}
