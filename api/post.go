package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
	"github.com/peacewalker122/project/util"
)

type CreatePostParams struct {
	AccountID          int64  `json:"account_id" validate:"required"`
	PictureDescription string `json:"picture_description"`
}

func (s *Server) createPost(c echo.Context) error {
	req := new(CreatePostParams)

	if err := c.Bind(req); err != nil {
		return err
	}
	strings, err := util.InputSqlString(req.PictureDescription, 3)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ValidateError("post_word", err.Error()))
	}
	account, err := s.store.GetAccounts(c.Request().Context(), req.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if account.Owner != authParam.Username {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	arg := db.CreatePostParams{
		AccountID:          account.ID,
		PictureDescription: strings,
	}

	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, PostResponse(post))
}

type GetPostParam struct {
	ID int `uri:"id" validate:"required,min=1"`
}

func (s *Server) getPost(c echo.Context) error {
	req := new(GetPostParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := ValidateURIPost(req, c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	acc, err := s.store.GetAccountsOwner(c.Request().Context(), authParam.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	post, err := s.store.GetPost(c.Request().Context(), int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if acc.ID != post.AccountID {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, PostResponse(post))
}
