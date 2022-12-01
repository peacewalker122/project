package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
)

const (
	AuthRefresh    = "RefreshToken"
	AuthHeaderkey  = "authorization"
	AuthTypeBearer = "bearer"
	AuthPayload    = "authorization_payload"
)

func authMiddleware(token token.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get(AuthHeaderkey)
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
			if authType != AuthTypeBearer {
				err := fmt.Errorf("invalid authorization %v type", authType)
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			authToken := authString[1]

			payload, err := token.VerifyToken(authToken)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, err.Error())
			}

			c.Set(AuthPayload, payload)
			return next(c)
		}
	}
}
