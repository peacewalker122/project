package user

import (
	"context"
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"github.com/peacewalker122/project/util"
	"net/http"

	"github.com/peacewalker122/project/token"
)

func (u *UserUsecase) Login(ctx context.Context, params *LoginParams) (*SessionResult, *util.Error) {

	if params == nil {
		return nil, util.NewError(http.StatusBadRequest, "invalid params")
	}
	// go u.email.SendEmailWithNotif(ctx, email.SendEmail{
	// 	AccountID: []int64{params.ID},
	// 	Params:    []string{params.Email, params.ClientIp},
	// 	Type:      "login",
	// 	TimeSend:  time.Now().UTC().Local(),
	// })

	userData, err := u.postgre.GetUser(ctx, params.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.NewError(http.StatusNotFound, "user not found")
		}
		return nil, util.NewError(http.StatusInternalServerError, err.Error())
	}

	account, err := u.postgre.GetAccountsOwner(ctx, params.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, util.NewError(http.StatusNotFound, "account not found")
		}
		return nil, util.NewError(http.StatusInternalServerError, err.Error())
	}

	err = util.CheckPassword(params.Password, userData.HashedPassword)
	if err != nil {
		return nil, util.NewError(http.StatusUnauthorized, "password is incorrect")
	}

	accesstoken, Accespayload, err := u.token.CreateToken(&token.PayloadRequest{
		Username:  params.Username,
		AccountID: account.ID,
		Duration:  u.config.TokenDuration,
	})
	if err != nil {
		return nil, util.NewError(http.StatusInternalServerError, err.Error())
	}

	refreshToken, refreshPayload, err := u.token.CreateToken(&token.PayloadRequest{
		Username:  params.Username,
		AccountID: account.ID,
		Duration:  u.config.RefreshToken,
	})
	if err != nil {
		return nil, util.NewError(http.StatusInternalServerError, err.Error())
	}

	arg := db.CreateSessionParams{
		ID:           Accespayload.ID,
		Username:     params.Username,
		RefreshToken: refreshToken,
		UserAgent:    params.UserAgent,
		ClientIp:     params.ClientIp,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := u.postgre.CreateSession(ctx, arg)
	if err != nil {
		return nil, util.NewError(http.StatusInternalServerError, err.Error())
	}

	res := &SessionResult{
		AccessToken:    accesstoken,
		RefreshToken:   refreshToken,
		Account:        account,
		User:           userData,
		AccessPayload:  Accespayload,
		RefreshPayload: refreshPayload,
		Session:        session,
	}
	return res, nil
}
