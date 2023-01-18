package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *AccountHandler) DeleteQueue(c echo.Context) error {
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

	queueData, err := a.account.DeleteQueue(ctx, req.ToAccountID, payload.AccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, queueData)
}
