package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	db "github.com/peacewalker122/project/db/sqlc"
)

type CreateAccountParams struct {
	Owner       string `json:"owner" validate:"require"`
	AccountType bool   `json:"account_type" validate:"require"`
}

func (s *Server) createAccount(c echo.Context) error {
	req := new(CreateAccountParams)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	arg := db.CreateAccountsParams{
		Owner:       req.Owner,
		AccountType: req.AccountType,
	}

	res, err := s.store.CreateAccounts(c.Request().Context(), arg)
	if err != nil {
		if pqerr, ok := err.(*pq.Error); ok {
			switch pqerr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err.Error())
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	output := AccountResponse(res)
	return c.JSON(http.StatusOK, output)
}
