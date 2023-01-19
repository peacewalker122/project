package tx

import (
	"context"

	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/params"
)

func (t *Tx) SetUsersOauth(ctx context.Context, params *params.CreateUsersOauthParam) (*OauthUserResponse, error) {
	var res *OauthUserResponse
	errTX := t.WithTx(ctx, func(tx *ent.Tx) error {
		var err error
		res.User, err = t.SetUser(ctx, params.User)
		if err != nil {
			return err
		}
		res.Token, err = t.SetToken(ctx, params.OauthToken)
		if err != nil {
			return err
		}
		res.Account, err = t.SetAccount(ctx, params.Account)
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
