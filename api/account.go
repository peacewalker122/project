package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
)

type GetAccountsParams struct {
	ID int `uri:"id" validate:"required,min=1"`
}

func (s *Server) getAccounts(c echo.Context) error {
	req := new(GetAccountsParams)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	num, err := ValidateURIAccount(req, c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err, ok := ValidationGetUser(num); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}
	account, err := s.store.GetAccounts(c.Request().Context(), int64(num.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "no such account")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if account.Owner != authParam.Username {
		err := errors.New("unauthorized Username for this account")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, AccountResponse(account))
}

func ValidationGetUser(input *GetAccountsParams) (errors string, ok bool) {
	if err := ValidateNum(input.ID); err != nil {
		ok = false
		errors = ValidateError("full_name", err.Error())
	}
	ok = true
	return errors, ok
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" query:"page_id" validate:"required,min=1"`
	PageSize int32 `form:"page_size" query:"page_size" validate:"required,min=5,max=50"`
}

func (server *Server) listAccount(c echo.Context) error {
	req := new(listAccountRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := ValidateCreateListAccount(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	arg := db.ListAccountsParams{
		Owner:  authParam.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, accounts)
}

func ValidateCreateListAccount(req *listAccountRequest) (errors []string) {
	if err := ValidateNum(int(req.PageID)); err != nil {
		errors = append(errors, ValidateError("page_id", err.Error()))
	}
	if err := ValidateNum(int(req.PageSize)); err != nil {
		errors = append(errors, ValidateError("page_size", err.Error()))
	}
	return errors
}
