package account

import (
	"github.com/labstack/echo/v4"
	auth "github.com/peacewalker122/project/api/auth"
	handler "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/contract/account"
	"github.com/peacewalker122/project/token"
)

type AccountHandler struct {
	account account.AccountContract
	helper  handler.Helper
	token   token.Maker
	Router  *echo.Echo
}

func NewAccountHandler(account account.AccountContract, helper handler.Helper, router *echo.Echo, token token.Maker) *AccountHandler {
	return &AccountHandler{
		account: account,
		helper:  helper,
		token:   token,
		Router:  router,
	}
}

func (a *AccountHandler) Register() {
	group := a.Router.Group("/account", auth.AuthMiddleware(a.token))
	group.GET("", a.GetAccount)
	group.GET("/queue", a.GetAllQueue)

	group.GET("s", a.ListAccount)
	group.POST("/follow", a.FollowAccount)
	group.PATCH("/unfollow", a.UnfollowAccount)
	group.PATCH("/private", a.UpdateAccountStatus)
	group.DELETE("/queue", a.DeleteQueue)
}
