package auth

import (
	"context"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model/params"

	"github.com/peacewalker122/project/util/email"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthUsecase) ChangePasswordAuth(ctx context.Context, req ChangePassParams) error {
	var (
		payload *ent.Users
	)

	err := a.redis.GetRedisPayload(ctx, req.UUID, &payload)
	if err != nil {
		return err
	}

	accountID, _ := a.postgre.GetAccountsOwner(ctx, payload.Username)

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	errchan := make(chan error, 1)
	done := make(chan struct{})

	err = a.postgre.ChangePasswordAuth(ctx, params.ChangePasswordParam{
		UUID:     req.UUID,
		Password: string(pass),
		Username: payload.Username,
		RedisDel: a.redis.Del,
	})

	go func() {
		err = a.email.SendEmailWithNotif(ctx, email.SendEmail{
			AccountID: []int64{accountID.ID},
			Type:      "password-changing",
			Params:    []string{payload.Email, payload.Username, req.ClientIp},
		})
		if err != nil {
			errchan <- err
		}
		done <- struct{}{}
	}()
	select {
	case <-done:
	case err := <-errchan:
		return err
	}

	return nil
}
