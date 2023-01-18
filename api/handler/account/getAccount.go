package account

import (
	"github.com/labstack/echo/v4"
	api "github.com/peacewalker122/project/api/handler"
	"github.com/peacewalker122/project/usecase/account"
	"net/http"
)

func (a *AccountHandler) GetAccount(c echo.Context) error {
	req := new(GetAccountParams)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	errNum, _, err := a.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err.Error())
	}

	accountData, err := a.account.GetAccount(ctx, req.ToAccountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := api.AccountResponse(*accountData)

	return c.JSON(http.StatusOK, res)
}

func (a *AccountHandler) ListAccount(c echo.Context) error {
	req := new(GetAccountsParams)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()

	errNum, _, err := a.helper.AuthAccount(c)
	if err != nil {
		return c.JSON(errNum, err.Error())
	}

	accountsData, err := a.account.ListAccount(ctx, &account.GetAccountParams{
		Offset:      req.Page,
		ToAccountID: req.ToAccountID,
		Limit:       req.Limit,
	})

	pageInfo := map[string]interface{}{
		"page":  req.Page,
		"limit": req.Limit,
	}

	res := api.FuncPublicAccountsResp(*accountsData, pageInfo)

	return c.JSON(http.StatusOK, res)
}
