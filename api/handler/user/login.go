package user

import (
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/user"
	"net/http"
)

func (u *UserHandler) Login(c echo.Context) error {
	req := new(LoginParams)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	loginRequest := &user.LoginParams{
		Username:  req.Username,
		Password:  req.Password,
		ClientIp:  c.RealIP(),
		UserAgent: c.Request().UserAgent(),
	}

	res, errLogin := u.user.Login(c.Request().Context(), loginRequest)
	if errLogin != nil {
		return c.JSON(errLogin.Error())
	}

	return c.JSON(http.StatusOK, res)
}
