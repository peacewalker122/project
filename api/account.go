package api

import (
	"database/sql"
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
	if errors := ValidateCreateAccount(req); errors != nil {
		return c.JSON(http.StatusBadRequest, errors)
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

func ValidateCreateAccount(input *CreateAccountParams) (errors []string) {
	if err := ValidateAlphanum(input.Owner); err != nil {
		errors = append(errors, ValidateError("owner", err.Error()))
	}
	return errors
}

type GetAccountsParam struct {
	ID int `uri:"id" validate:"required,min=1"`
}

func (s *Server) GetAccounts(c echo.Context) error {
	req := new(GetAccountsParam)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := ValidateURI(req, "id"); err != nil {
		return err
	}

	account, err := s.store.GetAccounts(c.Request().Context(), int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "no such account")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, AccountResponse(account))
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" validate:"required,min=1"`
	PageSize int32 `form:"page_size" validate:"required,min=5,max=50"`
}

func (server *Server) listAccount(c echo.Context) error {
	req := new(listAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, accounts)
}
