package api

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	api "github.com/peacewalker122/project/api/util"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
)

type accountService interface {
	GetAccounts(c echo.Context) error
	ListAccount(c echo.Context) error
	FollowAccount(c echo.Context) error
	AcceptFollower(c echo.Context) error
	UpdatePrivate(c echo.Context) error
	UpdatePhotoProfile(c echo.Context) error
}

var (
	errNum int
	err    error
)

type (
	UpdateAccountParams struct {
		FromAccountID int64 `uri:"id" query:"from_account_id" validate:"required,min=1"`
	}
	UpdatePhotoProfileParams struct {
		FromAccountID int64 `uri:"id" query:"id" validate:"required,min=1"`
	}
	GetAccountsParams struct {
		Offset        int32 `json:"offset" form:"offset" query:"offset" validate:"required,min=0"`
		FromAccountID int   `uri:"from_account_id" query:"from_account_id" validate:"required,min=1"`
		ToAccountID   int   `uri:"id" query:"to_account_id" validate:"required,min=1"`
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
		FromAccountID int64 `json:"from_account_id" query:"from_account_id" validate:"required,min=1"`
		ToAccountID   int64 `json:"to_account_id" query:"to_account_id" validate:"required,min=1"`
	}
)

func (s *Handler) GetAccounts(c echo.Context) error {
	req := new(GetAccountsParams)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// if err = c.Validate(req); err != nil {
	// 	return err
	// }
	num, err := ValidateURIAccount(req, c, "id")
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

	if req.FromAccountID == req.ToAccountID {
		queue, err := s.store.ListQueue(c.Request().Context(), db.ListQueueParams{
			Limit:     10,
			Offset:    (req.Offset - 1) * 10,
			Accountid: int64(req.ToAccountID),
		})
		if errNum, err = GetErrorValidator(c, err, "accounts-queue"); err != nil {
			return c.JSON(errNum, err.Error())
		}

		return c.JSON(http.StatusOK, OwnerAccountResponse(account, queue...))
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

	if req.FromAccountID == req.ToAccountID {
		return c.JSON(http.StatusBadRequest, "INVALID")
	}

	IsPrivate, err := s.store.GetAccountsInfo(c.Request().Context(), req.ToAccountID)
	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
		return c.JSON(errNum, err)
	}
	isQueue, err := s.store.GetQueueRows(c.Request().Context(), db.GetQueueRowsParams{
		Fromaccountid: req.FromAccountID,
		Toaccountid:   req.ToAccountID,
	})
	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
		return c.JSON(errNum, err)
	}

	if !req.Follow {

		if IsPrivate.IsPrivate {
			if isQueue == 0 {
				return c.JSON(http.StatusBadRequest, "no queue")
			}
			if isQueue == 1 {
				err = s.store.DeleteAccountQueue(c.Request().Context(), db.DeleteAccountQueueParams{
					Fromaccountid: req.FromAccountID,
					Toaccountid:   req.ToAccountID,
				})
				if errNum, err = GetErrorValidator(c, err, "account-queue"); err != nil {
					return c.JSON(errNum, err)
				}
				return c.JSON(http.StatusOK, echo.Map{
					"Status":    "OK",
					"DeletedAt": time.Now().Unix(),
				})
			}
		}

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

	if IsPrivate.IsPrivate {

		if num == 1 {
			return c.JSON(http.StatusBadRequest, errors.New("please be patient").Error())
		}

		err = s.util.CreateAccountsQueue(c.Request().Context(), &api.CreateQueue{
			FromAccountID: req.FromAccountID,
			ToAccountID:   req.ToAccountID,
		})
		if errNum, err = GetErrorValidator(c, err, "account-queue"); err != nil {
			c.JSON(errNum, err.Error())
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Status":    "OK",
			"CreatedAt": time.Now().Unix(),
		})
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
	req := new(AcceptAccountRequest)

	if err = c.Bind(req); err != nil {
		return err
	}

	log.Print(req)
	if err = c.Validate(req); err != nil {
		return err
	}

	if errNum, err = s.AuthAccount(c, req.FromAccountID); err != nil {
		return c.JSONPretty(errNum, err.Error(), "\t")
	}
	if req.FromAccountID == req.ToAccountID {
		return c.JSON(http.StatusBadRequest, "can't accept urself!")
	}
	account, err := s.store.GetAccountsQueue(c.Request().Context(), db.GetAccountsQueueParams{
		Fromid: req.ToAccountID,
		Toid:   req.FromAccountID,
	})
	if errNum, err = GetErrorValidator(c, err, "accounts-queue"); err != nil {
		return c.JSON(errNum, err.Error())
	}

	_, err = s.store.Followtx(c.Request().Context(), db.FollowTXParam{
		Fromaccid: req.ToAccountID,
		Toaccid:   req.FromAccountID,
		IsQueue:   account,
	})
	if err != nil {
		return CreateErrorValidator(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Status": true,
	})
}

func (s *Handler) UpdatePrivate(c echo.Context) error {
	req := new(UpdateAccountParams)

	if err = c.Bind(req); err != nil {
		return err
	}
	req.FromAccountID, err = ValidateURI[int64](c, "id")
	if err != nil {
		return err
	}
	if errNum, err = s.AuthAccount(c, req.FromAccountID); err != nil {
		return c.JSONPretty(errNum, err.Error(), "\t")
	}

	account, err := s.store.GetAccounts(c.Request().Context(), req.FromAccountID)
	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
		return c.JSON(errNum, err.Error())
	}

	err = s.store.PrivateAccount(context.Background(), db.PrivateAccountParams{
		IsPrivate: !account.IsPrivate,
		Username:  account.Owner,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "status: OK")
}

func (s *Handler) UpdatePhotoProfile(c echo.Context) error {
	// add form with photo param as a file bridge to sent into api

	req := new(UpdatePhotoProfileParams)

	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	req.FromAccountID, err = ValidateURI[int64](c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if errNum, err = s.AuthAccount(c, req.FromAccountID); err != nil {
		return c.JSON(errNum, err.Error())
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	errNum, err = s.util.UpdateProfilePhoto(c, req.FromAccountID)
	if err != nil {
		return c.JSON(errNum, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"Status":    "Updated",
		"UpdatedAt": time.Now().Unix(),
	})
}
