package contract

import (
	"github.com/peacewalker122/project/contract/auth"
	"github.com/peacewalker122/project/contract/user"
	db "github.com/peacewalker122/project/db/repository/postgres"
	"github.com/peacewalker122/project/db/repository/redis"
	authuseCase "github.com/peacewalker122/project/usecase/auth"
	useruseCase "github.com/peacewalker122/project/usecase/user"
	"github.com/peacewalker122/project/util"
)

type Contract interface {
	user.UserContract
	auth.AuthContract
}

type contract struct {
	*useruseCase.UserUsecase
	*authuseCase.AuthUsecase
}

func NewContract(postgre db.PostgresStore, redis redis.Store, cfg util.Config) Contract {
	return &contract{
		UserUsecase: useruseCase.NewUserUsecase(postgre, redis, cfg),
		AuthUsecase: authuseCase.NewAuthUsecase(postgre, redis, cfg),
	}
}
