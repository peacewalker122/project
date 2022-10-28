package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/token"
)

const (
	authRefresh    = "RefreshToken"
	authHeaderkey  = "authorization"
	authTypeBearer = "bearer"
	authPayload    = "authorization_payload"
)

func authMiddleware(token token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(authHeaderkey)
			if len(authHeader) == 0 {
				err := errors.New("authorization header is empty")
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			authString := strings.Fields(authHeader)
			if len(authString) < 2 {
				err := errors.New("unknown authorization type")
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			authType := strings.ToLower(authString[0])
			if authType != authTypeBearer {
				err := fmt.Errorf("invalid authorization %v type", authType)
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			authToken := authString[1]

			payload, err := token.VerifyToken(authToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			c.Set(authPayload, payload)
			return next(c)
		}
	}
}

func (s *Server) AuthAccount(c echo.Context, accountID int64) error {
	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	acc, err := s.store.GetAccounts(c.Request().Context(), accountID)
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}
	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized Username for this account")
		c.JSON(http.StatusUnauthorized, err.Error())
		return c.Redirect(http.StatusPermanentRedirect, s.config.AuthErrorAddres)
	}
	return nil
}
