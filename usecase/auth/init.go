package auth

import (
	db "github.com/peacewalker122/project/service/db/repository/postgres"
	"github.com/peacewalker122/project/service/db/repository/redis"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/peacewalker122/project/util/email"
)

type AuthUsecase struct {
	postgre db.PostgresStore
	redis   redis.Store
	email   *email.EmailHelper
	token   token.Maker
	config  util.Config
}

func NewAuthUsecase(postgre db.PostgresStore, redis redis.Store, cfg util.Config) *AuthUsecase {

	token, _ := token.NewJwt(cfg.TokenKey)

	return &AuthUsecase{
		postgre: postgre,
		redis:   redis,
		email:   email.NewEmailHelper(postgre, redis, cfg),
		token:   token,
		config:  cfg,
	}
}
