package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *AccountHandler) UpdateAccountStatus(c echo.Context) error {
	ctx := c.Request().Context()

	errNum, payload, err := a.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err.Error())
	}

	status, err := a.account.PrivateAccount(ctx, payload.AccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, status)
}
