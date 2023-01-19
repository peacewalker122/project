package tx

import (
	"context"
	ent2 "github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/params"
)

func (t *Tx) ChangePasswordAuth(ctx context.Context, params params.ChangePasswordParam) error {
	err := t.WithTx(ctx, func(tx *ent2.Tx) error {
		var err error

		err = t.ChangePasswordAuth(ctx, params)
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