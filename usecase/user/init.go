package user

import (
	db "github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/db/repository/redis"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/peacewalker122/project/util/email"
)

type UserUsecase struct {
	postgre db.PostgresStore
	redis   redis.Store
	email   *email.EmailHelper
	token   token.Maker
	config  util.Config
}

func NewUserUsecase(postgre db.PostgresStore, redis redis.Store, cfg util.Config) *UserUsecase {

	token, _ := token.NewJwt(cfg.TokenKey)

	return &UserUsecase{
		postgre: postgre,
		redis:   redis,
		email:   email.NewEmailHelper(postgre, redis, cfg),
		token:   token,
		config:  cfg,
	}
}