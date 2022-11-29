package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
)

type (
	GetAccountsParams struct {
		ID int `uri:"id" validate:"required,min=1"`
	}
	listAccountRequest struct {
		PageID   int32 `form:"page_id" query:"page_id" validate:"required,min=1"`
		PageSize int32 `form:"page_size" query:"page_size" validate:"required,min=5,max=50"`
	}
	FollowAccountRequest struct {
		FromAccountID int64 `json:"from_account_id" query:"from_account_id" validate:"required"`
		ToAccountID   int64 `json:"to_account_id" query:"to_account_id" validate:"required"`
		Follow        bool  `json:"follow" query:"follow" `
	}
	AcceptAccountRequest struct {
		FromAccountID int64 `json:"from_account_id" query:"from_account_id" validate:"required"`
		ToAccountID   int64 `json:"to_account_id" query:"to_account_id" validate:"required"`
		accept        bool  `json:"accept" query:"accept" `
	}
)

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

func (s *Server) followAccount(c echo.Context) error {
	var (
		result db.FollowTXResult
		delete db.UnFollowTXResult
		num    int64
		err    error
	)
	req := new(FollowAccountRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return err
	}

	if !req.Follow {
		num, err = s.store.GetAccountsFollowRows(c.Request().Context(), db.GetAccountsFollowRowsParams{Fromid: req.FromAccountID, Toid: req.ToAccountID})
		if err := GetErrorValidator(c, err, accountag); err != nil {
			return err
		}
		if num != 1 {
			return c.JSON(http.StatusNotFound, "not follow yet")
		}
		delete, err = s.store.UnFollowtx(c.Request().Context(), db.UnfollowTXParam{
			Fromaccid: req.FromAccountID,
			Toaccid:   req.ToAccountID,
		})
		if err := GetErrorValidator(c, err, accountag); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"Status": delete.Status,
			"Type":   delete.FeatureType,
		})
	}

	num, err = s.store.GetAccountsFollowRows(c.Request().Context(), db.GetAccountsFollowRowsParams{Fromid: req.FromAccountID, Toid: req.ToAccountID})
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}

	// here we getting info did account is private or no
	ok, err := s.store.GetAccountsInfo(c.Request().Context(), req.FromAccountID)

	if ok.IsPrivate {
		res, err := s.store.CreatePrivateQueue(c.Request().Context(), db.CreatePrivateQueueParams{
			Fromaccountid: req.FromAccountID,
			Toaccountid:   req.ToAccountID,
		})
		if err = CreateErrorValidator(c, err); err != nil {
			return err
		}
		val := fmt.Sprintf("queue status: %v", res.Queue)
		return c.JSON(http.StatusOK, val)
	}

	if num != 0 {
		return c.JSON(http.StatusBadRequest, "already follow")
	}

	result, err = s.store.Followtx(c.Request().Context(), db.FollowTXParam{
		Fromaccid: req.FromAccountID,
		Toaccid:   req.ToAccountID,
	})
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, followResponse(result))
}

func (s *Server) acceptFollower(c echo.Context) error {
	req := &AcceptAccountRequest{}

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return c.JSONPretty(http.StatusUnauthorized, err.Error(), "\t")
	}

}
