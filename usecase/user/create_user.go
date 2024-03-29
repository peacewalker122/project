package user

import (
	"context"
	"encoding/json"

	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/user"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/user"

	"github.com/peacewalker122/project/util"
)

func (s *UserUsecase) CreateUser(ctx context.Context, requid string, token int) (*result.CreateUserResult, error) {
	var (
		val PayloadCreateUser
		// multierror *util.Error
	)
	_, err := s.email.VerifyEmailAuth(ctx, requid, int(token))
	if err != nil {
		return nil, err
	}

	// here we get the value that user input from redis
	result, err := s.redis.Get(ctx, requid+"h")
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(result), &val); err != nil {
		return nil, err
	}

	hashpass, err := util.HashPassword(val.Password)
	if err != nil {
		return nil, err
	}
	arg := request.CreateUserParamsTx{
		Username: val.Username,
		Password: hashpass,
		FullName: val.FullName,
		Email:    val.Email,
	}

	res, err := s.postgre.CreateUserTX(ctx, &arg)
	if err != nil {
		return nil, err
	}

	return res, nil
}
