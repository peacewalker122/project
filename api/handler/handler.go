package api

import (
	apiutil "github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/contract"
	db "github.com/peacewalker122/project/service/db/repository/postgres"
	"github.com/peacewalker122/project/service/db/repository/redis"
	"github.com/peacewalker122/project/service/gcp"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

func NewHandler(store db.PostgresStore, gcpClient gcp.GCPService, redis redis.Store, config util.Config, token token.Maker, apiutil apiutil.UtilTools) (HandlerService, string) {
	return &Handler{
		store:  store,
		redis:  redis,
		config: config,
		token:  token,
		util:   apiutil,
	}, FileName
}

type Handler struct {
	util     apiutil.UtilTools
	store    db.PostgresStore
	redis    redis.Store
	config   util.Config
	contract contract.Contract
	token    token.Maker
}

type HandlerService interface {
	postService
	tokenService
	userService
	accountService
	AuthHandler
	Helper
}

const (
	AuthRefresh    = "RefreshToken"
	AuthHeaderkey  = "authorization"
	AuthTypeBearer = "bearer"
	AuthPayload    = "authorization_payload"
)
