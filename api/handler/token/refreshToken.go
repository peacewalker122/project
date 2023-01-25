package token

import (
	"errors"
	"github.com/labstack/echo/v4"
	auth "github.com/peacewalker122/project/api/auth"
	"net/http"
)

func (t *TokenHandler) RefreshToken(c echo.Context) error {
	token := c.Request().Header.Get(auth.AuthRefresh)
	if token == "" {
		err := errors.New("invalid, no token")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	refreshToken, errToken := t.token.RefreshToken(c.Request().Context(), token)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, errToken.Error())
	}

	return c.JSON(http.StatusOK, refreshToken)
}
