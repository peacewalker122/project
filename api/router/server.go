package api

import (
	"fmt"
	"github.com/peacewalker122/project/api/handler/account"
	"github.com/peacewalker122/project/api/handler/post"
	tokenhandler "github.com/peacewalker122/project/api/handler/token"
	"github.com/peacewalker122/project/api/handler/user"
	db "github.com/peacewalker122/project/service/db/repository/postgres"
	"github.com/peacewalker122/project/service/db/repository/redis"
	"github.com/peacewalker122/project/service/gcp"
	account2 "github.com/peacewalker122/project/usecase/account"
	auth2 "github.com/peacewalker122/project/usecase/auth"
	post2 "github.com/peacewalker122/project/usecase/post"
	tokenusecase "github.com/peacewalker122/project/usecase/token"
	user2 "github.com/peacewalker122/project/usecase/user"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	auth "github.com/peacewalker122/project/api/auth"
	handler "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/api/oauth"
	apiutil "github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

type Server struct {
	Store      db.PostgresStore
	Redis      redis.Store
	Config     util.Config
	handler    handler.HandlerService
	Auth       *Util
	Router     *echo.Echo
	Oauth      oauth.OauthService
	Token      token.Maker
	FileString string
	account    *account.AccountHandler
	user       *user.UserHandler
	post       *post.PostHandler
	token      *tokenhandler.TokenHandler
	apiutil.UtilTools
	gcp.GCPService
}

func Newserver(c util.Config, store db.PostgresStore, redisStore redis.Store, service gcp.GCPService) (*Server, error) {
	newtoken, err := token.NewJwt(c.TokenKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token %v", err.Error())
	}
	server := &Server{
		Store:      store,
		Redis:      redisStore,
		Config:     c,
		GCPService: service,
		Auth:       NewUtil(validator.New()),
		Token:      newtoken,
	}
	server.UtilTools = apiutil.NewApiUtil(store, redisStore, c)
	server.handler = handler.NewHandler(store, service, redisStore, c, newtoken, server.UtilTools)
	server.Oauth = oauth.NewHandler(store, redisStore, c, newtoken, server.UtilTools)

	server.account = account.NewAccountHandler(
		account2.NewAccountUseCase(store, redisStore, c, service),
		server.handler,
		server.Router,
		newtoken,
	)

	server.user = user.NewUserHandler(
		auth2.NewAuthUsecase(store, redisStore, c),
		user2.NewUserUsecase(store, redisStore, c),
	)

	server.post = post.NewPostHandler(
		post2.NewPostUsecase(store, redisStore, c, service),
		server.handler,
	)

	server.token = tokenhandler.NewTokenHandler(
		tokenusecase.NewTokenUsecase(newtoken, store, c),
	)

	server.Routerhandle()

	return server, nil
}

func (s *Server) Routerhandle() {
	router := echo.New()
	router.Use(middlewareLogging)
	router.Validator = s.Auth.Validator
	router.HTTPErrorHandler = s.Auth.HTTPErrorHandler

	userGroup := router.Group("/user")
	s.user.Router(userGroup)
	s.post.PostRouter(router)
	s.token.TokenRouter(router)

	OauthRouter := router.Group("/oauth")
	OauthRouter.GET("/google", s.Oauth.GoogleVerif)
	OauthRouter.GET("/google/callback", s.Oauth.GoogleToken)

	authRouter := router.Group("/auth", auth.AuthMiddleware(s.Token))

	s.account.Router(authRouter)

	//authRouter.POST("/post", s.handler.CreatePost, middleware.TimeoutWithConfig(s.TimeoutPost()))
	//authRouter.GET("/post/:id", s.handler.GetPost)
	//authRouter.POST("/post/like", s.handler.LikePost)
	//authRouter.POST("/post/comment", s.handler.CommentPost)
	//authRouter.POST("/post/retweet", s.handler.RetweetPost)
	//authRouter.GET("/post/image/:id", s.handler.GetPostImage, middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	//authRouter.POST("/post/qoute/retweet", s.handler.QouteretweetPost)

	s.Router = router
}

func (s *Server) StartHTTPS(path string) error {
	return s.Router.StartAutoTLS(path)
}
func (s *Server) StartHTTP(path string) error {
	s.Router.Logger.Info("server is running on ", path)
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
	//router.POST("/user", s.handler.CreateUser)

	//AuthMethod := router.Group("", auth.AuthMiddleware(s.Token))

	//AuthMethod.POST("/account", s.createAccount)
	//AuthMethod.POST("/post", s.handler.CreatePost)
	//AuthMethod.GET("/post/:id", s.handler.GetPost)

	s.Router = router
}
