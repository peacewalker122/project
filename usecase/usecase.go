package usecase

import (
	"github.com/peacewalker122/project/contract"
	"github.com/peacewalker122/project/service/db/repository/postgres"
	"github.com/peacewalker122/project/service/db/repository/redis"
	"github.com/peacewalker122/project/service/gcp"
	"github.com/peacewalker122/project/usecase/account"
	"github.com/peacewalker122/project/usecase/auth"
	"github.com/peacewalker122/project/usecase/user"
	"github.com/peacewalker122/project/util"
)

type usecase struct {
	*user.UserUsecase
	*auth.AuthUsecase
	*account.AccountUseCase
}

func NewUsecase(postgre postgres.PostgresStore, redis redis.Store, cfg util.Config, gcpClient gcp.GCPService) contract.Contract {
	return &usecase{
		UserUsecase:    user.NewUserUsecase(postgre, redis, cfg),
		AuthUsecase:    auth.NewAuthUsecase(postgre, redis, cfg),
		AccountUseCase: account.NewAccountUseCase(postgre, redis, cfg, gcpClient),
	}
}
