package user

import (
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/user"
)

func (u *UserHandler) UpdateUser(c echo.Context) error {
	req := new(UpdateUserParam)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	errNum, _, err := u.Helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err)
	}

	errUpdate := u.user.UpdateUser(c.Request().Context(), &user.UpdateUserParam{
		Username: req.Username,
		FullName: req.Fullname,
		Email:    req.Email,
	})
	if errUpdate.HasError() {
		return c.JSON(500, errUpdate.Errors)
	}

	return c.JSON(201, nil)
}
