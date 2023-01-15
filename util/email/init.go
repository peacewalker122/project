package email

import (
	db "github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/db/repository/redis"
	"github.com/peacewalker122/project/util"
)

type EmailHelper struct {
	cfg     util.Config
	redis   redis.Store
	postgre db.PostgresStore
}

func NewEmailHelper(store db.PostgresStore, redis redis.Store, cfg util.Config) *EmailHelper {
	return &EmailHelper{
		cfg:     cfg,
		redis:   redis,
		postgre: store,
	}
}
