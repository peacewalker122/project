package oauth

import (
	"github.com/labstack/echo/v4"
	apiutil "github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

func NewHandler(store db.Store, redis redis.Store, config util.Config, token token.Maker, apiutil apiutil.UtilTools) OauthService {
	return &Handler{
		store:  store,
		redis:  redis,
		config: config,
		token:  token,
		util:   apiutil,
	}
}

type Handler struct {
	util   apiutil.UtilTools
	store  db.Store
	redis  redis.Store
	config util.Config
	token  token.Maker
}

type OauthService interface {
	GoogleVerif(c echo.Context) error
	GoogleToken(c echo.Context) error
}

const (
	AuthRefresh    = "RefreshToken"
	AuthHeaderkey  = "authorization"
	AuthTypeBearer = "bearer"
	AuthPayload    = "authorization_payload"
)
