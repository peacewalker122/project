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
		AccountID          int64  `json:"account_id" form:"account_id" validate:"required"`
		PictureDescription string `json:"picture_description" form:"picture_description" validate:"required"`
	}
	GetImageParam struct {
		FromAccountID int64 `json:"from_account_id" form:"account_id" validate:"required"`
		postID        int64 `uri:"id" validate:"required,min=1"`
	}
	GetPostParam struct {
		postID        int64 `uri:"id" validate:"required,min=1"`
		Offset        int32 `json:"offset" form:"offset" query:"offset" validate:"required,min=0"`
		FromAccountID int64 `json:"from_account_id" query:"accid" validate:"required,min=1"`
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
	if err := ValidateString(req.PictureDescription, 1, 70); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := s.AuthAccount(c, req.AccountID); err != nil {
		return err
	}

	return s.creatingPost(c, req)
}

func (s *Server) getPost(c echo.Context) error {
	req := new(GetPostParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := req.ValidateURIPost(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return err
	}
	return s.GettingPost(c, req)
}

func (s *Server) getPostImage(c echo.Context) error {
	req := new(GetImageParam)
	if err := c.Bind(req); err != nil {
		return err
	}
	err := req.ValidateURIPost(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return err
	}

	return s.gettingImage(c, req.postID)
}

func (s *Server) likePost(c echo.Context) error {
	req := new(LikePostRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return err
	}

	ok, err := s.store.GetLikejoin(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err, like); err != nil {
		return err
	}

	if ok && req.IsLike {
		return c.JSON(http.StatusBadRequest, "already like")
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
	err := ValidateString(req.Comment, 1, 70)
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
	var (
		err   error
		num   int64
		ok    bool
		Cpost db.Post
		Fpost db.PostFeature
	)

	req := new(RetweetPostRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return err
	}

	if req.IsRetweet {
		num, err = s.store.GetRetweetRows(c.Request().Context(), db.GetRetweetRowsParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
		if err := GetErrorValidator(c, err, retweet); err != nil {
			return err
		}
		if num == 0 {
			Cpost, Fpost, err = s.createRetweetPost(req, c)
			if err != nil {
				return err
			}
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

	if ok && req.IsRetweet {
		return c.JSON(http.StatusBadRequest, "already retweet")
	}

	if !ok {
		if req.IsRetweet {
			post.SumRetweet++
		}
	}

	if !req.IsRetweet {
		return s.deleteRetweetpost(req, c, post)
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
	_, err = s.store.UpdatePost_feature(c.Request().Context(), arg)
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, retweetResponse(Fpost, Cpost))
}

func (s *Server) qouteretweetPost(c echo.Context) error {
	var (
		err   error
		ok    bool
		num   int64
		Cpost db.Post
		Fpost db.PostFeature
	)

	req := new(QouteRetweetPostRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := ValidateString(req.Qoute, 1, 70); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	if err := s.AuthAccount(c, req.FromAccountID); err != nil {
		return err
	}

	if req.IsRetweet {
		num, err = s.store.GetQouteRetweetRows(c.Request().Context(), db.GetQouteRetweetRowsParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
		if err := GetErrorValidator(c, err, qretweet); err != nil {
			return err
		}
		if num == 0 {
			_, err = s.store.CreateQouteRetweet_feature(c.Request().Context(), db.CreateQouteRetweet_featureParams{FromAccountID: req.FromAccountID, PostID: req.PostID, Qoute: req.Qoute})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		res, err := s.store.GetPostQRetweetJoin(c.Request().Context(), db.GetPostQRetweetJoinParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
		if err := GetErrorValidator(c, err, qretweet); err != nil {
			return err
		}
		if !res.QouteRetweet {
			// to validate if the Retweet&qoute_retweet is false then execute below.
			Cpost, Fpost, err = s.createQouteRetweetPost(req, c)
			if err != nil {
				return err
			}
		}
	}

	post, err := s.store.GetPost_feature_Update(c.Request().Context(), req.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ok, err = s.store.GetQouteRetweetJoin(c.Request().Context(), req.PostID)
	if err := GetErrorValidator(c, err, qretweet); err != nil {
		return err
	}
	if ok && req.IsRetweet {
		return c.JSON(http.StatusBadRequest, "already created")
	}

	if !ok {
		if req.IsRetweet {
			post.SumQouteRetweet++
		}
	}

	if !req.IsRetweet {
		err = s.deleteQouteRetweet(req, c, post)
		return err
	}

	err = s.store.UpdateQouteRetweet(c.Request().Context(), db.UpdateQouteRetweetParams{FromAccountID: req.FromAccountID, PostID: req.PostID, QouteRetweet: req.IsRetweet})
	if err := CreateErrorValidator(c, err); err != nil {
		return err
	}

	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{
		FromAccountID: req.FromAccountID,
		PostID:        post.PostID,
		TypeEntries:   qretweet,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = s.store.UpdatePost_feature(c.Request().Context(), db.UpdatePost_featureParams{
		PostID:          post.PostID,
		SumComment:      post.SumComment,
		SumLike:         post.SumLike,
		SumRetweet:      post.SumRetweet,
		SumQouteRetweet: post.SumQouteRetweet,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, qouteretweetResponse(Cpost, Fpost, req.Qoute))
}
