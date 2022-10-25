package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/token"
)

type (
	CreatePostParams struct {
		AccountID          int64  `json:"account_id" validate:"required"`
		PictureDescription string `json:"picture_description" validate:"required"`
	}
	GetPostParam struct {
		ID     int   `uri:"id" validate:"required,min=1"`
		Offset int32 `json:"offset" form:"offset" validate:"required,min=0"`
	}
	LikePostRequest struct {
		FromAccountID int64 `json:"from_account_id" validate:"required"`
		IsLike        bool  `json:"like"`
		PostID        int64 `json:"post_id" validate:"required"`
	}
	CommentPostRequest struct {
		FromAccountID int64  `json:"from_account_id" validate:"required"`
		Comment       string `json:"comment" form:"comment" validate:"required"`
		PostID        int64  `json:"post_id" validate:"required"`
	}
	RetweetPostRequest struct {
		FromAccountID int64 `json:"from_account_id" validate:"required"`
		IsRetweet     bool  `json:"retweet"`
		PostID        int64 `json:"post_id" validate:"required"`
	}
	QouteRetweetPostRequest struct {
		FromAccountID int64  `json:"from_account_id" validate:"required"`
		IsRetweet     bool   `json:"retweet"`
		Qoute         string `json:"qoute" form:"qoute" validate:"required"`
		PostID        int64  `json:"post_id" validate:"required"`
	}
)

func (s *Server) createPost(c echo.Context) error {
	req := new(CreatePostParams)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := ValidateString(req.PictureDescription); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	account, err := s.store.GetAccounts(c.Request().Context(), req.AccountID)
	if err := GetErrorValidator(c, err, accountag); err != nil {
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
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}

	post, err := s.store.GetPost(c.Request().Context(), int64(req.ID))
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}
	postFeature, err := s.store.GetPost_feature(c.Request().Context(), int64(req.ID))
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}

	arg := db.ListCommentParams{PostID: int64(req.ID), Limit: int32(10), Offset: (req.Offset - 1) * 10}
	comment, err := s.store.ListComment(c.Request().Context(), arg)
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
	return c.JSON(http.StatusOK, GetPostResponse(post, postFeature, comment))
}

func (s *Server) likePost(c echo.Context) error {
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
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}
	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized Username for this account")
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	_, err = s.store.GetLikeInfo(c.Request().Context(), db.GetLikeInfoParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
	if err != nil {
		if err == sql.ErrNoRows {
			err = s.store.CreateLike_feature(c.Request().Context(), db.CreateLike_featureParams{
				FromAccountID: req.FromAccountID,
				PostID:        req.PostID,
			})
			if err := CreateErrorValidator(c, err); err != nil {
				return err
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	post, err := s.store.GetPost_feature_Update(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}

	ok, err = s.store.GetLikejoin(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err, like); err != nil {
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

	err = s.store.UpdateLike(c.Request().Context(), db.UpdateLikeParams{IsLike: req.IsLike, PostID: req.PostID, FromAccountID: req.FromAccountID})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	entries := like
	if !req.IsLike {
		entries = unlike
	}
	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{
		FromAccountID: req.FromAccountID,
		PostID:        post.PostID,
		TypeEntries:   entries,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := s.store.UpdatePost_feature(c.Request().Context(), db.UpdatePost_featureParams{
		PostID:          post.PostID,
		SumComment:      post.SumComment,
		SumLike:         post.SumLike,
		SumRetweet:      post.SumRetweet,
		SumQouteRetweet: post.SumQouteRetweet,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, likeResponse(result))
}

func (s *Server) commentPost(c echo.Context) error {
	req := new(CommentPostRequest)

	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	err := ValidateString(req.Comment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ValidateError("comment", err.Error()))
	}

	acc, err := s.store.GetAccounts(c.Request().Context(), req.FromAccountID)
	if err := GetErrorValidator(c, err, accountag); err != nil {
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
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}

	post.SumComment++

	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{FromAccountID: req.FromAccountID, PostID: req.PostID, TypeEntries: comment})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	result, err := s.store.UpdatePost_feature(c.Request().Context(), db.UpdatePost_featureParams{
		PostID:          req.PostID,
		SumComment:      post.SumComment,
		SumLike:         post.SumLike,
		SumRetweet:      post.SumRetweet,
		SumQouteRetweet: post.SumQouteRetweet,
	})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, commentResponse(commentResult, result))
}

func (s *Server) retweetPost(c echo.Context) error {
	req := new(RetweetPostRequest)
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
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}
	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized Username for this account")
		c.JSON(http.StatusUnauthorized, err.Error())
		return c.Redirect(http.StatusPermanentRedirect, s.config.AuthErrorAddres)
	}

	_, err = s.store.GetRetweet(c.Request().Context(), db.GetRetweetParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
	if err != nil {
		if err == sql.ErrNoRows {
			err = s.store.CreateRetweet_feature(c.Request().Context(), db.CreateRetweet_featureParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
			if err := CreateErrorValidator(c, err); err != nil {
				return err
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	post, err := s.store.GetPost_feature_Update(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}

	ok, err = s.store.GetRetweetJoin(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err, retweet); err != nil {
		return err
	}

	if !ok {
		if req.IsRetweet {
			post.SumRetweet++
		}
	}
	if ok {
		if !req.IsRetweet {
			post.SumRetweet--
		}
	}

	err = s.store.UpdateRetweet(c.Request().Context(), db.UpdateRetweetParams{Retweet: req.IsRetweet, PostID: req.PostID, FromAccountID: req.FromAccountID})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	args := db.CreateEntriesParams{
		FromAccountID: req.FromAccountID,
		PostID:        post.PostID,
		TypeEntries:   retweet,
	}
	if !req.IsRetweet {
		args.TypeEntries = unretweet
	}
	_, err = s.store.CreateEntries(c.Request().Context(), args)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	arg := db.UpdatePost_featureParams{
		PostID:          req.PostID,
		SumComment:      post.SumComment,
		SumLike:         post.SumLike,
		SumRetweet:      post.SumRetweet,
		SumQouteRetweet: post.SumQouteRetweet,
	}
	update, err := s.store.UpdatePost_feature(c.Request().Context(), arg)
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, retweetResponse(update))
}

func(s *Server) qouteretweetPost(c echo.Context) error{
	req := new(QouteRetweetPostRequest)
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
	if err := GetErrorValidator(c, err, accountag); err != nil {
		return err
	}
	if acc.Owner != authParam.Username {
		err := errors.New("unauthorized Username for this account")
		c.JSON(http.StatusUnauthorized, err.Error())
		return c.Redirect(http.StatusPermanentRedirect, s.config.AuthErrorAddres)
	}

	
}