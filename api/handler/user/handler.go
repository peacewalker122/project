package user

import (
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/contract/auth"
	"github.com/peacewalker122/project/contract/user"
)

type UserHandler struct {
	auth auth.AuthContract
	user user.UserContract
}

func NewUserHandler(auth auth.AuthContract, user user.UserContract) *UserHandler {
	return &UserHandler{
		auth: auth,
		user: user,
	}
}

func (u *UserHandler) Router(c *echo.Group) {
	c.POST("", u.CreateUserRequest)
	c.POST("/change-password", u.ChangePasswordRequest)
	c.POST("/change-password/:uid", u.ChangePassword)
	c.POST("/login", u.Login)
	c.POST("/signup/:uuid", u.CreateUser)
}
