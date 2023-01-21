package api

//
//import (
//	"context"
//	"errors"
//	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
//	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
//	"log"
//	"net/http"
//	"time"
//
//	"github.com/labstack/echo/v4"
//	api "github.com/peacewalker122/project/api/util"
//)
//
//type accountService interface {
//}
//
//var (
//	errNum int
//	err    error
//)
//
//type (
//	listAccountRequest struct {
//		PageID   int32 `form:"page_id" query:"page_id" validate:"required,min=1"`
//		PageSize int32 `form:"page_size" query:"page_size" validate:"required,min=5,max=50"`
//	}
//	FollowAccountRequest struct {
//		ToAccountID int64 `json:"to_account_id" query:"to_account_id" validate:"required"`
//		Follow      bool  `json:"follow" query:"follow"`
//	}
//	AcceptAccountRequest struct {
//		ToAccountID int64 `json:"to_account_id" query:"to_account_id" validate:"required,min=1"`
//	}
//)
//
//func (s *Handler) ListAccountQueue(c echo.Context) error {
//	errNum, _, err := s.AuthAccount(c)
//	if err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	accounts, err := s.store.ListQueue(c.Request().Context(), db2.ListQueueParams{})
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, accounts)
//}
//
//func (s *Handler) ListAccount(c echo.Context) error {
//	req := new(listAccountRequest)
//
//	if err = c.Bind(req); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if err = c.Validate(req); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	errNum, authParam, err := s.AuthAccount(c)
//	if err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	arg := db2.ListAccountsParams{
//		Owner:  authParam.Username,
//		Limit:  req.PageSize,
//		Offset: (req.PageID - 1) * req.PageSize,
//	}
//
//	accounts, err := s.store.ListAccounts(c.Request().Context(), arg)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, FuncPublicAccountResponse(accounts))
//}
//
//func (s *Handler) FollowAccount(c echo.Context) error {
//	var (
//		result db.FollowTXResult
//		delete db.UnFollowTXResult
//		num    int64
//		err    error
//		errNum int
//	)
//	req := new(FollowAccountRequest)
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	errNum, authParam, err := s.AuthAccount(c)
//	if err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	if authParam.AccountID == req.ToAccountID {
//		return c.JSON(http.StatusBadRequest, "INVALID")
//	}
//
//	IsPrivate, err := s.store.GetAccountsInfo(c.Request().Context(), req.ToAccountID)
//	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//		return c.JSON(errNum, err)
//	}
//	isQueue, err := s.store.GetQueueRows(c.Request().Context(), db2.GetQueueRowsParams{
//		Fromaccountid: authParam.AccountID,
//		Toaccountid:   req.ToAccountID,
//	})
//	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	if !req.Follow {
//		if IsPrivate.IsPrivate {
//
//			if isQueue == 1 {
//				err = s.store.DeleteAccountQueue(c.Request().Context(), db2.DeleteAccountQueueParams{
//					Fromaccountid: authParam.AccountID,
//					Toaccountid:   req.ToAccountID,
//				})
//				if errNum, err = GetErrorValidator(c, err, "account-queue"); err != nil {
//					return c.JSON(errNum, err)
//				}
//				return c.JSON(http.StatusOK, echo.Map{
//					"Status":    "OK",
//					"DeletedAt": time.Now().Unix(),
//				})
//			}
//		}
//
//		num, err = s.store.GetAccountsFollowRows(c.Request().Context(), db2.GetAccountsFollowRowsParams{Fromid: authParam.AccountID, Toid: req.ToAccountID})
//		if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//			return c.JSON(errNum, err)
//		}
//		if num != 1 {
//			return c.JSON(http.StatusNotFound, "not follow yet")
//		}
//		delete, err = s.store.UnFollowtx(c.Request().Context(), db.UnfollowTXParam{
//			Fromaccid: authParam.AccountID,
//			Toaccid:   req.ToAccountID,
//		})
//		if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//			return c.JSON(errNum, err)
//		}
//		return c.JSON(http.StatusOK, echo.Map{
//			"Status": delete.Status,
//			"Type":   delete.FeatureType,
//		})
//	}
//
//	num, err = s.store.GetAccountsFollowRows(c.Request().Context(), db2.GetAccountsFollowRowsParams{Fromid: authParam.AccountID, Toid: req.ToAccountID})
//	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	if IsPrivate.IsPrivate {
//
//		if num == 1 {
//			return c.JSON(http.StatusBadRequest, errors.New("in queue").Error())
//		}
//
//		err = s.util.CreateAccountsQueue(c.Request().Context(), &api.CreateQueue{
//			FromAccountID: authParam.AccountID,
//			ToAccountID:   req.ToAccountID,
//		})
//		if errNum, err = GetErrorValidator(c, err, "account-queue"); err != nil {
//			return c.JSON(errNum, err.Error())
//		}
//
//		return c.JSON(http.StatusOK, echo.Map{
//			"Status":    "OK",
//			"CreatedAt": time.Now().Unix(),
//		})
//	}
//
//	if num != 0 {
//		return c.JSON(http.StatusBadRequest, "already follow")
//	}
//
//	result, err = s.store.Followtx(c.Request().Context(), db.FollowTXParam{
//		Fromaccid: authParam.AccountID,
//		Toaccid:   req.ToAccountID,
//		IsQueue:   false,
//	})
//	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	return c.JSON(http.StatusOK, followResponse(result))
//}
//
//func (s *Handler) AcceptFollower(c echo.Context) error {
//	req := new(AcceptAccountRequest)
//
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//
//	log.Print(req)
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//
//	errNum, authParam, err := s.AuthAccount(c)
//	if err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	if authParam.AccountID == req.ToAccountID {
//		return c.JSON(http.StatusBadRequest, "can't accept urself!")
//	}
//	account, err := s.store.GetAccountsQueue(c.Request().Context(), db2.GetAccountsQueueParams{
//		Fromid: req.ToAccountID,
//		Toid:   authParam.AccountID,
//	})
//	if errNum, err = GetErrorValidator(c, err, "accounts-queue"); err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//
//	_, err = s.store.Followtx(c.Request().Context(), db.FollowTXParam{
//		Fromaccid: req.ToAccountID,
//		Toaccid:   authParam.AccountID,
//		IsQueue:   account,
//	})
//	if err != nil {
//		return CreateErrorValidator(c, err)
//	}
//
//	return c.JSON(http.StatusOK, echo.Map{
//		"Status": true,
//	})
//}
//
//func (s *Handler) UpdatePrivate(c echo.Context) error {
//	errNum, authParam, err := s.AuthAccount(c)
//	if err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	account, err := s.store.GetAccounts(c.Request().Context(), authParam.AccountID)
//	if errNum, err = GetErrorValidator(c, err, Accountag); err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//
//	err = s.store.PrivateAccount(context.Background(), db2.PrivateAccountParams{
//		IsPrivate: !account.IsPrivate,
//		Username:  account.Owner,
//	})
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, "status: OK")
//}
//
//func (s *Handler) UpdatePhotoProfile(c echo.Context) error {
//	// add form with photo param as a file bridge to sent into api
//
//	errNum, authParam, err := s.AuthAccount(c)
//	if err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	errNum, err = s.util.UpdateProfilePhoto(c, authParam.AccountID)
//	if err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, echo.Map{
//		"Status":    "Updated",
//		"UpdatedAt": time.Now().Unix(),
//	})
//}
