package user

import (
	"context"
	"encoding/json"

	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/util"
)

func (s *UserUsecase) CreateUser(ctx context.Context, requid string, token int) (*db.CreateUserTXResult, *util.Error) {
	var (
		val        map[string]interface{}
		multierror *util.Error
	)

	_, err := s.email.VerifyEmailAuth(ctx, requid, int(token))
	if err != nil {
		multierror.Important(err.Error(), "email")
		return nil, multierror
	}

	// here we get the value that user input from redis
	result, err := s.redis.Get(ctx, requid+"h")
	if err != nil {
		multierror.Important(err.Error(), "redis")
		return nil, multierror
	}
	if err := json.Unmarshal([]byte(result), &val); err != nil {
		multierror.Important(err.Error(), "marshal")
		return nil, multierror
	}

	hashpass, err := util.HashPassword(val["password"].(string))
	if err != nil {
		multierror.Important(err.Error(), "hash password")
		return nil, multierror
	}
	arg := db.CreateUserParamsTx{
		Username: val["username"].(string),
		Password: hashpass,
		FullName: val["full_name"].(string),
		Email:    val["email"].(string),
	}

	res, err := s.postgre.CreateUserTx(ctx, arg)
	if err != nil {
		multierror.Important(res.Error.Error(), "create user")
		return nil, multierror
	}

	return &res, nil
}
