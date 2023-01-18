package user

import (
	"context"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc"

	"github.com/peacewalker122/project/token"
)

func (u *UserUsecase) Login(ctx context.Context, params SessionParams) (*SessionResult, error) {

	var (
	// errs *util.Error
	)

	// go u.email.SendEmailWithNotif(ctx, email.SendEmail{
	// 	AccountID: []int64{params.ID},
	// 	Params:    []string{params.Email, params.ClientIp},
	// 	Type:      "login",
	// 	TimeSend:  time.Now().UTC().Local(),
	// })

	if params.ID == nil {
		ID, err := u.postgre.GetAccountID(ctx, params.Username)
		if err != nil {
			return nil, err
		}
		params.ID = &ID
	}

	accesstoken, Accespayload, err := u.token.CreateToken(&token.PayloadRequest{
		Username:  params.Username,
		AccountID: *params.ID,
		Duration:  u.config.TokenDuration,
	})
	if err != nil {
		return nil, err
	}

	refreshToken, refreshPayload, err := u.token.CreateToken(&token.PayloadRequest{
		Username:  params.Username,
		AccountID: *params.ID,
		Duration:  u.config.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     params.Username,
		RefreshToken: refreshToken,
		UserAgent:    params.UserAgent,
		ClientIp:     params.ClientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := u.postgre.CreateSession(ctx, arg)
	if err != nil {
		return nil, err
	}

	res := &SessionResult{
		AccessToken:    accesstoken,
		RefreshToken:   refreshToken,
		AccessPayload:  Accespayload,
		RefreshPayload: refreshPayload,
		Session:        session,
	}
	return res, nil
}
