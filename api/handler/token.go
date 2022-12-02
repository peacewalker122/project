package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	auth "github.com/peacewalker122/project/api/auth"
)

type tokenService interface{
	RenewToken(c echo.Context) error
}

func (s *Handler) RenewToken(c echo.Context) error {
	token := c.Request().Header.Get(auth.AuthRefresh)
	if token == "" {
		err := errors.New("invalid, no token")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	payload, err := s.token.VerifyToken(token)
	if err != nil {
		errs := fmt.Errorf("invalid token %v", err.Error())
		c.JSON(http.StatusBadRequest, errs.Error())
	}

	session, err := s.store.GetSession(c.Request().Context(), payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if session.IsBlocked {
		err := errors.New("session is blocked")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if session.ID != payload.ID {
		err := errors.New("incorrect session user")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if session.RefreshToken != token {
		err := fmt.Errorf("mismatch session token")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	accesToken, accesPayload, err := s.token.CreateToken(payload.Username, s.config.RefreshToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	rsp := AccesTokenResp{
		AccesToken:          accesToken,
		AccesTokenExpiresAt: accesPayload.ExpiredAt.Local().UTC(),
	}
	return c.JSON(http.StatusOK, rsp)
}
