package account

import (
	"github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/db/repository/redis"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
	"github.com/peacewalker122/project/util/email"
)

type AccountUseCase struct {
	postgre postgres.PostgresStore
	redis   redis.Store
	email   *email.EmailHelper
	token   token.Maker
	config  util.Config
}

func NewAccountUseCase(postgre postgres.PostgresStore, redis redis.Store, cfg util.Config) *AccountUseCase {

	token, _ := token.NewJwt(cfg.TokenKey)

	return &AccountUseCase{
		postgre: postgre,
		redis:   redis,
		email:   email.NewEmailHelper(postgre, redis, cfg),
		token:   token,
		config:  cfg,
	}
}
