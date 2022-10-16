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
