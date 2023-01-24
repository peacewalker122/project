package token

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	jwttoken "github.com/peacewalker122/project/token"
	"time"
)

func (t *TokenUsecase) RefreshToken(ctx context.Context, token string) (*AccesTokenResp, error) {
	payload, err := t.token.VerifyToken(token)
	if err != nil {
		errs := fmt.Errorf("invalid token %v", err.Error())
		return nil, errs
	}

	session, err := t.store.GetSession(ctx, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	if session.IsBlocked {
		return nil, errors.New("session is blocked")
	}

	if session.ID != payload.ID {
		return nil, errors.New("incorrect session user")
	}

	if session.RefreshToken != token {
		return nil, fmt.Errorf("mismatch session token")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, fmt.Errorf("expired session")
	}

	accesToken, accesPayload, err := t.token.CreateToken(&jwttoken.PayloadRequest{
		Username:  payload.Username,
		AccountID: payload.AccountID,
		Duration:  t.cfg.RefreshToken,
	})

	if err != nil {
		return nil, err
	}
	rsp := AccesTokenResp{
		AccesToken:          accesToken,
		AccesTokenExpiresAt: accesPayload.ExpiredAt.Local().UTC(),
	}
	return &rsp, nil
}
