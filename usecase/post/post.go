package post

import (
	db "github.com/peacewalker122/project/service/db/repository/postgres"
	"github.com/peacewalker122/project/service/db/repository/redis"
	"github.com/peacewalker122/project/service/gcp"
	"github.com/peacewalker122/project/util"
)

type PostUsecase struct {
	postgre db.PostgresStore
	redis   redis.Store
	config  util.Config
	gcp     gcp.GCPService
}

func NewPostUsecase(postgre db.PostgresStore, redis redis.Store, config util.Config, gcp gcp.GCPService) *PostUsecase {
	return &PostUsecase{
		postgre: postgre,
		redis:   redis,
		config:  config,
		gcp:     gcp,
	}
}
