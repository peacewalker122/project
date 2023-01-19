package user

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/auth"
	"net/http"
)

func (u *UserHandler) ChangePasswordRequest(c echo.Context) error {
	req := new(ValidateChangePassRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := uuid.New()

	errReq := u.auth.CreateRequest(c.Request().Context(), auth.AuthParams{
		Email: req.Email,
		UUID:  id,

		ClientIp: c.RealIP(),
	})
	if errReq != nil {
		return c.JSON(http.StatusBadRequest, errReq.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (u *UserHandler) ChangePassword(c echo.Context) error {
	req := new(ChangePasswordRequest)
	ctx := c.Request().Context()

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	uid := c.Param("uid")

	errChange := u.auth.ChangePasswordAuth(ctx, auth.ChangePassParams{
		Password: req.Password,
		UUID:     uid,
	})
	if errChange != nil {
		return c.JSON(http.StatusBadRequest, errChange.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
