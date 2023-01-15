package auth

import (
	"context"
	"time"

	"github.com/peacewalker122/project/util/email"
)

func (a *AuthUsecase) CreateRequest(ctx context.Context, params AuthParams) error {
	errchan := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		err := a.email.ChangePasswordAuth(ctx, email.SendEmail{
			Params:   []string{params.Email, params.ClientIp, params.UUID.String()},
			Type:     "change_password",
			TimeSend: time.Now(),
		})
		if err != nil {
			errchan <- err
		}

		tempVal, err := a.postgre.GetAllWithEmail(ctx, params.Email)
		if err != nil {
			errchan <- err
		}

		err = a.redis.Set(ctx, params.UUID.String(), tempVal, 5*time.Minute)
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
