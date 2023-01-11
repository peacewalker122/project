package model

import (
	"context"

	"github.com/peacewalker122/project/db/ent"
)

type TX interface {
	SetUsersOauth(ctx context.Context, params *CreateUsersOauthParam) (*OauthUserResponse, error)
}

func (s *Models) SetUsersOauth(ctx context.Context, params *CreateUsersOauthParam) (*OauthUserResponse, error) {
	var res *OauthUserResponse
	errTX := s.WithTx(ctx, func(tx *ent.Tx) error {
		var err error
		res.User, err = s.SetUser(ctx, params.User)
		if err != nil {
			return err
		}
		res.Token, err = s.SetToken(ctx, params.OauthToken)
		if err != nil {
			return err
		}
		res.Account, err = s.SetAccount(ctx, params.Account)
		if err != nil {
			return err
		}
		return nil
	})
	if errTX != nil {
		return nil, errTX
	}
	return res, nil
}
