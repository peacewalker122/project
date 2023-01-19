package account

import (
	"github.com/labstack/echo/v4"
	handler "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/contract/account"
	"github.com/peacewalker122/project/token"
)

type AccountHandler struct {
	account account.AccountContract
	helper  handler.Helper
	token   token.Maker
}

func NewAccountHandler(account account.AccountContract, helper handler.Helper, router *echo.Echo, token token.Maker) *AccountHandler {
	return &AccountHandler{
		account: account,
		helper:  helper,
		token:   token,
	}
}

func (a *AccountHandler) Router(e *echo.Group) {
	e.GET("/account", a.GetAccount)
	e.GET("/account/queue", a.GetAllQueue)
	e.GET("/accounts", a.ListAccount)
	e.POST("/account/follow", a.FollowAccount)
	e.PATCH("/account/unfollow", a.UnfollowAccount)
	e.PATCH("/account/private", a.UpdateAccountStatus)
	e.DELETE("/account/queue", a.DeleteQueue)
}
