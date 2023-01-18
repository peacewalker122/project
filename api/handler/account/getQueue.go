package account

import (
	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/usecase/account"
	"net/http"
)

func (a *AccountHandler) GetAllQueue(c echo.Context) error {
	req := new(GetAccountsParams)
	ctx := c.Request().Context()

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	errNum, payload, err := a.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err.Error())
	}

	queueData, err := a.account.ListQueuedAccount(ctx, &account.GetAccountParams{
		Offset:      req.Page,
		ToAccountID: payload.AccountID,
		Limit:       req.Limit,
	})
	result := make([]map[string]interface{}, len(*queueData))

	for _, data := range *queueData {
		result = append(result, map[string]interface{}{
			"username":   data.Owner,
			"account_id": data.FromAccountID,
		})
	}

	return c.JSON(http.StatusOK, result)
}
