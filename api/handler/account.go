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

type accountService interface {
	GetAccounts(c echo.Context) error
	ListAccount(c echo.Context) error
	FollowAccount(c echo.Context) error
	AcceptFollower(c echo.Context) error
}

var (
	errNum int
	err    error
)

type (
	GetAccountsParams struct {
		FromAccountID int `uri:"from_account_id" query:"from_account_id" validate:"required,min=1"`
		ToAccountID   int `uri:"to_account_id" query:"to_account_id" validate:"required,min=1"`
	}
	listAccountRequest struct {
		FromAccountID int   `uri:"from_account_id" query:"from_account_id" validate:"required,min=1"`
		PageID        int32 `form:"page_id" query:"page_id" validate:"required,min=1"`
		PageSize      int32 `form:"page_size" query:"page_size" validate:"required,min=5,max=50"`
	}
	FollowAccountRequest struct {
		FromAccountID int64 `json:"from_account_id" query:"from_account_id" validate:"required"`
		ToAccountID   int64 `json:"to_account_id" query:"to_account_id" validate:"required"`
		Follow        bool  `json:"follow" query:"follow" `
	}
	AcceptAccountRequest struct {
		FromAccountID int64 `json:"from_account_id" query:"from_account_id" validate:"required"`
		ToAccountID   int64 `json:"to_account_id" query:"to_account_id" validate:"required"`
	}
)

func (s *Handler) GetAccounts(c echo.Context) error {
	req := new(GetAccountsParams)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	num, err := ValidateURIAccount(req, c, "to_account_id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err, ok := ValidationGetUser(num); !ok {
		return c.JSON(http.StatusBadRequest, err)
	}

	account, err := s.store.GetAccounts(c.Request().Context(), int64(num.ToAccountID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "no such account")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if errNum, err = s.AuthAccount(c, int64(req.FromAccountID)); err != nil {
		return c.JSON(errNum, err)
	}

	return c.JSON(http.StatusOK, AccountResponse(account))
}

func (s *Handler) ListAccount(c echo.Context) error {
	req := new(listAccountRequest)

	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if errs := ValidateCreateListAccount(req); errs != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	authParam, ok := c.Get(AuthPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if errNum, err = s.AuthAccount(c, int64(req.FromAccountID)); err != nil {
		return c.JSON(errNum, err)
	}
	arg := db.ListAccountsParams{
		Owner:  authParam.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := s.store.ListAccounts(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, accounts)
}

func (s *Handler) FollowAccount(c echo.Context) error {
	var (
		result db.FollowTXResult
		delete db.UnFollowTXResult
		num    int64
		err    error
		errNum int
	)
	req := new(FollowAccountRequest)
	if err = c.Bind(req); err != nil {
		return err
	}
	if err = c.Validate(req); err != nil {
		return err
	}
	if errNum, err = s.AuthAccount(c, req.FromAccountID); err != nil {
		return c.JSON(errNum, err.Error())
	}

	if !req.Follow {
		num, err = s.store.GetAccountsFollowRows(c.Request().Context(), db.GetAccountsFollowRowsParams{Fromid: req.FromAccountID, Toid: req.ToAccountID})
		if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
			return c.JSON(errNum, err)
		}
		if num != 1 {
			return c.JSON(http.StatusNotFound, "not follow yet")
		}
		delete, err = s.store.UnFollowtx(c.Request().Context(), db.UnfollowTXParam{
			Fromaccid: req.FromAccountID,
			Toaccid:   req.ToAccountID,
		})
		if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
			return c.JSON(errNum, err)
		}
		return c.JSON(http.StatusOK, echo.Map{
			"Status": delete.Status,
			"Type":   delete.FeatureType,
		})
	}

	num, err = s.store.GetAccountsFollowRows(c.Request().Context(), db.GetAccountsFollowRowsParams{Fromid: req.FromAccountID, Toid: req.ToAccountID})
	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
		return c.JSON(errNum, err)
	}

	// here we getting info did account is private or no
	ok, err := s.store.GetAccountsInfo(c.Request().Context(), req.FromAccountID)
	if err != nil {
		errNum, err = GetErrorValidator(c, err, "account")
		return c.JSON(errNum, err.Error())
	}
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
		IsQueue:   false,
	})
	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
		return c.JSON(errNum, err)
	}

	return c.JSON(http.StatusOK, followResponse(result))
}

func (s *Handler) AcceptFollower(c echo.Context) error {
	req := &AcceptAccountRequest{}

	if err = c.Bind(req); err != nil {
		return err
	}
	if errNum, err = s.AuthAccount(c, req.FromAccountID); err != nil {
		return c.JSONPretty(errNum, err.Error(), "\t")
	}

	_, err = s.store.Followtx(c.Request().Context(), db.FollowTXParam{
		Fromaccid: req.FromAccountID,
		Toaccid:   req.ToAccountID,
		IsQueue:   true,
	})
	
	if err != nil {
		return CreateErrorValidator(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Status": true,
	})
}
