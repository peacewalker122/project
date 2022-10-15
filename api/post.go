package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
)

type CreatePostParams struct {
	AccountID          int64  `json:"account_id" validate:"required"`
	PictureDescription string `json:"picture_description" validate:"required"`
	PictureID          int64  `json:"pictureid" validate:"required"`
}

func (s *Server) createPost(c echo.Context) error {
	req := new(CreatePostParams)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := ValidatePostRequest(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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

	arg := db.CreatePostParams{
		AccountID:          account.ID,
		PictureDescription: strings,
		PictureID:          req.PictureID,
	}

	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, PostResponse(post))
}

func ValidatePostRequest(req *CreatePostParams) (errors []string) {
	if err := ValidateNum(int(req.PictureID)); err != nil {
		errors = append(errors, ValidateError("pictureid", err.Error()))
	}
	if err := ValidateNum(int(req.AccountID)); err != nil {
		errors = append(errors, ValidateError("account_id", err.Error()))
	}
	return errors
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
	account, err := s.store.GetPost(c.Request().Context(), int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, PostResponse(account))
}
