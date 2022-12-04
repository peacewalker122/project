package api

import (
	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

func NewHandler(store db.Store, redis redis.Store, config util.Config, token token.Maker) (HandlerService, string) {
	return &Handler{
		store:  store,
		redis:  redis,
		config: config,
		token:  token,
	}, FileName
}

type Handler struct {
	store    db.Store
	redis    redis.Store
	config   util.Config
	token    token.Maker
}

type HandlerService interface {
	postService
	tokenService
	userService
	accountService
}

const (
	AuthRefresh    = "RefreshToken"
	AuthHeaderkey  = "authorization"
	AuthTypeBearer = "bearer"
	AuthPayload    = "authorization_payload"
)
