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
	if err := GetErrorValidator(c, err); err != nil {
		return err
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
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}

	post, err := s.store.GetPost(c.Request().Context(), int64(req.ID))
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}
	postFeature, err := s.store.GetPost_feature(c.Request().Context(), int64(req.ID))
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}
	if acc.AccountsID != post.AccountID {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	return c.JSON(http.StatusOK, PostResponse(post, postFeature))
}

type LikePostRequest struct {
	FromAccountID int64 `json:"from_account_id" validate:"required"`
	IsLike        bool  `json:"is_like"`
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
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}

	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	_, err = s.store.GetLikeInfo(c.Request().Context(), db.GetLikeInfoParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = s.store.CreateLike_feature(c.Request().Context(), db.CreateLike_featureParams{
				FromAccountID: req.FromAccountID,
				IsLike:        req.IsLike,
				PostID:        req.PostID,
			})
			if err := CreateErrorValidator(c, err); err != nil {
				return err
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	post, err := s.store.GetPost_feature_Update(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}

	ok, err = s.store.GetLikejoin(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}
	if !ok {
		if req.IsLike {
			post.SumLike++
		}
	}
	if ok {
		if !req.IsLike {
			post.SumLike--
		}
	}
	_, err = s.store.UpdateLike(c.Request().Context(), db.UpdateLikeParams{IsLike: req.IsLike, PostID: req.PostID, FromAccountID: req.FromAccountID})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}
	_, err = s.store.GetEntriesFull(c.Request().Context(), db.GetEntriesFullParams{PostID: req.PostID, FromAccountID: req.FromAccountID, TypeEntries: like})
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{
				FromAccountID: req.FromAccountID,
				PostID:        post.PostID,
				TypeEntries:   like,
			})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
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

type CommentPostRequest struct {
	FromAccountID int64  `json:"from_account_id" validate:"required"`
	Comment       string `json:"comment" validate:"required"`
	PostID        int64  `json:"post_id" validate:"required"`
}

func (s *Server) CommentPost(c echo.Context) error {
	req := new(CommentPostRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	err := ValidateAlphanum(req.Comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ValidateError("comment", err.Error()))
	}
	acc, err := s.store.GetAccounts(c.Request().Context(), req.FromAccountID)
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}
	authParam, ok := c.Get(authPayload).(*token.Payload)
	if !ok {
		err := errors.New("failed conversion")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized username")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	commentResult, err := s.store.CreateComment_feature(c.Request().Context(), db.CreateComment_featureParams{FromAccountID: req.FromAccountID, Comment: req.Comment, PostID: req.PostID})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	post, err := s.store.GetPost_feature_Update(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err); err != nil {
		return err
	}

	post.SumComment++

	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{FromAccountID: req.FromAccountID, PostID: req.PostID, TypeEntries: comment})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	result, err := s.store.UpdatePost_feature(c.Request().Context(), db.UpdatePost_featureParams{
		PostID:     req.PostID,
		SumComment: post.SumComment,
	})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, commentResponse(commentResult, result))
}
