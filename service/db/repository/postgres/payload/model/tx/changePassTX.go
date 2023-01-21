package tx

import (
	"context"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/params"
)

func (t *Tx) ChangePasswordAuth(ctx context.Context, params params.ChangePasswordParam) error {
	err := t.WithTx(ctx, func(tx *Tx) error {
		var err error

		err = t.SetPassword(ctx, params.Username, params.Password)
		if err != nil {
			return err
		}

		err = params.RedisDel(ctx, params.UUID)
		if err != nil {
			return err
		}

		return err
	})

	return err
}