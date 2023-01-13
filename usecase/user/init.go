package user

import (
	"github.com/peacewalker122/project/api/util"
	db "github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/db/repository/redis"
)

type UserUsecase struct {
	Postgre db.PostgresStore
	Util    util.UtilTools
	Redis   redis.Store
}

func NewUserUsecase(postgre db.PostgresStore, util util.UtilTools, redis redis.Store) *UserUsecase {
	return &UserUsecase{
		Postgre: postgre,
		Util:    util,
		Redis:   redis,
	}
}
