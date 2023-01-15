package auth

import (
	"context"

	"github.com/peacewalker122/project/db/repository/postgres/ent"
	"github.com/peacewalker122/project/util"
	"github.com/peacewalker122/project/util/email"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthUsecase) ChangePasswordAuth(ctx context.Context, req ChangePassParams) *util.Error {
	var (
		payload *ent.Users
		errs    *util.Error
	)

	err := a.redis.GetRedisPayload(ctx, req.UUID, &payload)
	if err != nil {
		errs.Important(err.Error(), "get-redis-payload")
		return errs
	}

	accountID, _ := a.postgre.GetAccountsOwner(ctx, payload.Username)

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errs.Important(err.Error(), "bcrypt")
		return errs
	}

	errchan := make(chan error, 1)
	done := make(chan struct{})

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
		errs.Important(err.Error(), "error")
		return errs
	}

	err = a.postgre.SetPassword(ctx, payload.Username, string(pass))
	if err != nil {
		errs.Important(err.Error(), "set-password")
		return errs
	}

	return nil
}
