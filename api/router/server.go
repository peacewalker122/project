package api

import (
	"fmt"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	auth "github.com/peacewalker122/project/api/auth"
	handler "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/api/oauth"
	apiutil "github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/db/repository/redis"
	db "github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

type Server struct {
	Store   db.PostgresStore
	Redis   redis.Store
	Config  util.Config
	handler handler.HandlerService
	Auth    *Util
	Router  *echo.Echo
	Oauth   oauth.OauthService
	apiutil.UtilTools
	Token      token.Maker
	FileString string
}

func Newserver(c util.Config, store db.PostgresStore, redisStore redis.Store) (*Server, error) {
	newtoken, err := token.NewJwt(c.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token %v", err.Error())
	}
	server := &Server{
		Store:  store,
		Redis:  redisStore,
		Config: c,
		Auth:   NewUtil(validator.New()),
		Token:  newtoken,
	}
	server.UtilTools = apiutil.NewApiUtil(store, redisStore, c)
	server.handler, server.FileString = handler.NewHandler(store, redisStore, c, newtoken, server.UtilTools)
	server.Oauth = oauth.NewHandler(store, redisStore, c, newtoken, server.UtilTools)
	server.routerhandle()
	return server, nil
}

func (s *Server) routerhandle() {
	router := echo.New()
	router.Use(middlewareLogging)
	//router.Use(middleware.HTTPSRedirectWithConfig(Redirect()))
	router.Validator = s.Auth.Validator
	router.HTTPErrorHandler = s.Auth.HTTPErrorHandler

	router.POST("/user", s.handler.CreateRequestUser)
	router.POST("/user/signup/:uuid", s.handler.CreateUser)
	router.POST("/token/renew", s.handler.RenewToken)
	router.POST("/user/login", s.handler.Login)

	router.POST("/user/forget", s.handler.AuthUser)
	router.POST("/user/request/forget/:uid", s.handler.ChangePassword)

	OauthRouter := router.Group("/oauth")
	OauthRouter.GET("/google", s.Oauth.GoogleVerif)
	OauthRouter.GET("/google/callback", s.Oauth.GoogleToken)

	authRouter := router.Group("/auth", auth.AuthMiddleware(s.Token))
	authRouter.GET("/account/:id", s.handler.GetAccounts)
	authRouter.GET("/account", s.handler.ListAccount)
	authRouter.POST("/account/private/:id", s.handler.UpdatePrivate)
	authRouter.POST("/account/profile/photo/:id", s.handler.UpdatePhotoProfile)
	authRouter.POST("/account/follow", s.handler.FollowAccount)
	authRouter.PUT("/account/follow", s.handler.AcceptFollower)
	authRouter.POST("/post", s.handler.CreatePost, middleware.TimeoutWithConfig(s.TimeoutPost()))
	authRouter.GET("/post/:id", s.handler.GetPost)
	authRouter.POST("/post/like", s.handler.LikePost)
	authRouter.POST("/post/comment", s.handler.CommentPost)
	authRouter.POST("/post/retweet", s.handler.RetweetPost)
	authRouter.GET("/post/image/:id", s.handler.GetPostImage, middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	authRouter.POST("/post/qoute/retweet", s.handler.QouteretweetPost)

	s.Router = router
}

func (s *Server) StartHTTPS(path string) error {
	return s.Router.StartAutoTLS(path)
}
func (s *Server) StartHTTP(path string) error {
	return s.Router.Start(path)
}
func (s *Server) timeout(c echo.Context) error {
	return c.JSON(echo.ErrRequestTimeout.Code, "timeout")
}
func (s *Server) TimeoutPost() middleware.TimeoutConfig {
	return middleware.TimeoutConfig{
		ErrorMessage: "timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			// we delete the file if its already timeout
			if _, err := os.Stat(s.FileString); err == nil {
				os.Remove(s.FileString)
			}
			c.Error(err)
			c.SetHandler(s.timeout)
		},
		Timeout: 8 * time.Second,
	}
}

func (s *Server) Testrouterhandle() {
	router := echo.New()
	router.Validator = NewValidator(validator.New())
	// router.HTTPErrorHandler = HTTPErrorHandler
	// router.Use(middleware.LoggerWithConfig(Logger()))
	// router.Binder = new(CustomBinder)
	router.POST("/user", s.handler.CreateUser)

	AuthMethod := router.Group("", auth.AuthMiddleware(s.Token))

	//AuthMethod.POST("/account", s.createAccount)
	AuthMethod.GET("/account/:id", s.handler.GetAccounts)
	AuthMethod.GET("/account", s.handler.ListAccount)
	AuthMethod.POST("/post", s.handler.CreatePost)
	AuthMethod.GET("/post/:id", s.handler.GetPost)

	s.Router = router
}
