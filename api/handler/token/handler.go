package token

import (
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/contract/token"
)

type TokenHandler struct {
	token token.TokenContract
}

func NewTokenHandler(token token.TokenContract) *TokenHandler {
	return &TokenHandler{
		token: token,
	}
}

func (t *TokenHandler) TokenRouter(e *echo.Echo) {
	e.POST("/token", t.RefreshToken)
}
