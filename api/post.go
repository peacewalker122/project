package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
)

type CreatePostParams struct {
	AccountID   int64  `json:"account_id"`
	PostWord    string `json:"post_word"`
	PostPicture []byte `json:"post_picture"`
}

func (s *Server) createPost(c echo.Context) error {
	req := new(CreatePostParams)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	strings, err := util.InputSqlString(req.PostWord, 3)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ValidateError("post_word", err.Error()))
	}
	account, err := s.store.GetAccounts(c.Request().Context(), req.AccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, fmt.Errorf("no such account %v", err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	arg := db.CreatePostParams{
		AccountID:   account.ID,
		PostWord:    strings,
		PostPicture: req.PostPicture,
	}

	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, PostResponse(post))
}
