package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
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
	if err := ValidateAlphanum(req.PictureDescription); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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
		AccountID:          account.AccountsID,
		PictureDescription: req.PictureDescription,
	}

	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	post2, err := s.store.CreatePost_feature(c.Request().Context(), post.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, PostResponse(post, post2))
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
	postFeature, err := s.store.GetPost_feature(c.Request().Context(), int64(req.ID))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if acc.AccountsID != post.AccountID {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, PostResponse(post, postFeature))
}

type LikePostRequest struct {
	FromAccountID int64 `json:"from_account_id" validate:"required"`
	IsLike        bool  `json:"is_like" validate:"required"`
	PostID        int64 `json:"post_id" validate:"required"`
}

func (s *Server) LikePost(c echo.Context) error {
	req := new(LikePostRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}

	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	acc, err := s.store.GetAccounts(c.Request().Context(), req.FromAccountID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	post, err := s.store.GetPost_feature_Update(c.Request().Context(), req.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if req.IsLike {
		post.SumLike++
	}
	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{
		FromAccountID: req.FromAccountID,
		PostID:        post.PostID,
		TypeEntries:   "Like",
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	result, err := s.store.UpdatePost_feature(c.Request().Context(), db.UpdatePost_featureParams{
		PostID:  post.PostID,
		SumLike: post.SumLike,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, likeResponse(result))
}
