package user

import (
	"github.com/labstack/echo/v4"
	handler "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/contract/auth"
	"github.com/peacewalker122/project/contract/user"
)

type UserHandler struct {
	auth auth.AuthContract
	user user.UserContract
	handler.Helper
}

func NewUserHandler(auth auth.AuthContract, user user.UserContract, helper handler.Helper) *UserHandler {
	return &UserHandler{
		auth:   auth,
		user:   user,
		Helper: helper,
	}
}

func (u *UserHandler) Router(c *echo.Group) {
	c.POST("", u.CreateUserRequest) // "/user"
	c.POST("/change-password", u.ChangePasswordRequest)
	c.POST("/change-password/:uid", u.ChangePassword)
	c.POST("/login", u.Login)
	c.POST("/signup/:uuid", u.CreateUser)
}

func (u *UserHandler) AuthRouter(c *echo.Group) {
	c.PATCH("/update", u.UpdateUser)
}
