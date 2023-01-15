package user

import (
	"context"

	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

func (u *UserUsecase) Login(ctx context.Context, params SessionParams) (*SessionResult, *util.Error) {

	var (
		errs *util.Error
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
			errs.Important(err.Error(), "get-account-id")
			return nil, errs
		}
		params.ID = &ID
	}

	accesstoken, Accespayload, err := u.token.CreateToken(&token.PayloadRequest{
		Username:  params.Username,
		AccountID: *params.ID,
		Duration:  u.config.TokenDuration,
	})
	if err != nil {
		errs.Important(err.Error(), "create-acctoken")
		return nil, errs
	}

	refreshToken, refreshPayload, err := u.token.CreateToken(&token.PayloadRequest{
		Username:  params.Username,
		AccountID: *params.ID,
		Duration:  u.config.RefreshToken,
	})
	if err != nil {
		errs.Important(err.Error(), "create-reftoken")
		return nil, errs
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
		errs.Important(err.Error(), "create-session")
		return nil, errs
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
