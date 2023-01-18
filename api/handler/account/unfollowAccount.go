package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *AccountHandler) UnfollowAccount(c echo.Context) error {
	req := new(GetAccountParams)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	errNum, payload, err := a.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err.Error())
	}

	status, err := a.account.UnFollowAccount(ctx, payload.AccountID, req.ToAccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, status)
}
