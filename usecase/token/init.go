package token

import (
	"github.com/peacewalker122/project/service/db/repository/postgres"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

type TokenUsecase struct {
	token token.Maker
	store postgres.PostgresStore
	cfg   *util.Config
}

func NewTokenUsecase(m token.Maker, p postgres.PostgresStore, c util.Config) *TokenUsecase {
	return &TokenUsecase{
		token: m,
		store: p,
		cfg:   &c,
	}
}
