package api

//
//import (
//	"database/sql"
//	"errors"
//	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
//	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
//	"net/http"
//	"time"
//
//	"github.com/labstack/echo/v4"
//)
//
//type postService interface {
//	CreatePost(c echo.Context) error
//	GetPost(c echo.Context) error
//	GetPostImage(c echo.Context) error
//	LikePost(c echo.Context) error
//	CommentPost(c echo.Context) error
//	RetweetPost(c echo.Context) error
//	QouteretweetPost(c echo.Context) error
//}
//
//type (
//	GetImageParam struct {
//		FromAccountID int64 `json:"from_account_id" form:"account_id" validate:"required"`
//		PostID        int64 `uri:"id" validate:"required,min=1"`
//	}
//	GetPostParam struct {
//		PostID        int64 `uri:"id" validate:"required,min=1"`
//		Offset        int32 `json:"offset" form:"offset" query:"offset" validate:"required,min=0"`
//		FromAccountID int64 `json:"from_account_id" query:"accid" validate:"required,min=1"`
//	}
//
//	LikePostRequest struct {
//		FromAccountID int64 `json:"from_account_id" validate:"required"`
//		IsLike        bool  `json:"like"`
//		PostID        int64 `json:"post_id" validate:"required"`
//	}
//	CommentPostRequest struct {
//		FromAccountID int64  `json:"from_account_id" validate:"required"`
//		Comment       string `json:"comment" form:"comment" validate:"required"`
//		PostID        int64  `json:"post_id" validate:"required"`
//	}
//	RetweetPostRequest struct {
//		FromAccountID int64 `json:"from_account_id" validate:"required"`
//		IsRetweet     bool  `json:"retweet"`
//		PostID        int64 `json:"post_id" validate:"required"`
//	}
//	QouteRetweetPostRequest struct {
//		FromAccountID int64  `json:"from_account_id" validate:"required"`
//		IsRetweet     bool   `json:"retweet"`
//		Qoute         string `json:"qoute" form:"qoute" validate:"required"`
//		PostID        int64  `json:"post_id" validate:"required"`
//	}
//)
//
//func (s *Handler) CreatePost(c echo.Context) error {
//	req := new(CreatePostParams)
//
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	if err = ValidateString(req.PictureDescription, 1, 70); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	return s.CreatingPost(c, req)
//}
//
//func (s *Handler) GetPost(c echo.Context) error {
//	req := new(GetPostParam)
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	err = req.ValidateURIPost(c, "id")
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//	return s.GettingPost(c, req)
//}
//
//func (s *Handler) GetPostImage(c echo.Context) error {
//	req := new(GetImageParam)
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	err := req.ValidateURIPost(c, "id")
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	return s.GettingImage(c, req.PostID)
//}
//
//func (s *Handler) LikePost(c echo.Context) error {
//	var result db.CreateLikeTXResult
//	req := new(LikePostRequest)
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	num, err := s.store.GetLikeRows(c.Request().Context(), db2.GetLikeRowsParams{
//		Fromaccountid: req.FromAccountID,
//		Postid:        req.PostID,
//	})
//	if errNum, err = GetErrorValidator(c, err, Like); err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//
//	if num == 0 {
//		errNum, err = s.CreateLike(c.Request().Context(), req)
//		if err != nil {
//			return c.JSON(errNum, err.Error())
//		}
//	}
//
//	ok, err := s.store.GetLikejoin(c.Request().Context(), req.PostID)
//	if errNum, err = GetErrorValidator(c, err, Like); err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//
//	if req.IsLike {
//		if ok {
//			return c.JSON(http.StatusBadRequest, "already like")
//		}
//		result, err = s.store.CreateLikeTX(c.Request().Context(), db.CreateLikeParams{
//			FromAccountID: req.FromAccountID,
//			PostID:        req.PostID,
//		})
//		if err != nil {
//			return c.JSON(result.ErrCode, err.Error())
//		}
//	}
//	if !req.IsLike {
//		if !ok {
//			return c.JSON(http.StatusBadRequest, "never like")
//		}
//		result, err = s.store.UnlikeTX(c.Request().Context(), db.CreateLikeParams{
//			FromAccountID: req.FromAccountID,
//			PostID:        req.PostID,
//		})
//		if err != nil {
//			return c.JSON(result.ErrCode, err.Error())
//		}
//		return c.JSON(http.StatusOK, echo.Map{
//			"Status":    "Deleted",
//			"DeletedAt": time.Now().Unix(),
//		})
//	}
//
//	return c.JSON(http.StatusOK, likeResponse(result.PostFeature))
//}
//
//func (s *Handler) CommentPost(c echo.Context) error {
//	req := new(CommentPostRequest)
//
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	err := ValidateString(req.Comment, 1, 70)
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, ValidateError("comment", err.Error()))
//	}
//
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	result, err := s.store.CreateCommentTX(c.Request().Context(), db.CreateCommentParams{
//		FromAccountID: req.FromAccountID,
//		PostID:        req.PostID,
//		Comment:       req.Comment,
//	})
//	if err != nil {
//		return c.JSON(result.ErrCode, err.Error())
//	}
//
//	return c.JSON(http.StatusOK, commentResponse(result.Comment, result.PostFeature))
//}
//
//func (s *Handler) RetweetPost(c echo.Context) error {
//	ctx := c.Request().Context()
//	var (
//		errNum int
//		err    error
//		num    int64
//		Result RetweetResponse
//	)
//
//	req := new(RetweetPostRequest)
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	num, err = s.store.GetRetweetRows(ctx, db2.GetRetweetRowsParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
//	if err != nil {
//		if err == sql.ErrNoRows {
//			errNum = http.StatusNotFound
//		}
//		return c.JSON(errNum, err.Error())
//	}
//
//	if req.IsRetweet {
//		if num == 0 {
//			Result, errNum, err = s.CreateRetweetPost(c, req)
//			if err != nil {
//				return c.JSON(errNum, err.Error())
//			}
//			return c.JSON(http.StatusOK, retweetResponse(Result.Feature, Result.Post))
//		}
//		return c.JSON(http.StatusBadRequest, errors.New("already exist").Error())
//	}
//
//	if !req.IsRetweet {
//		if num == 0 {
//			return c.JSON(http.StatusNotFound, "not found")
//		}
//		return s.DeleteRetweetpost(req, c)
//
//	}
//
//	return c.JSON(http.StatusOK, retweetResponse(Result.Feature, Result.Post))
//}
//func (s *Handler) QouteretweetPost(c echo.Context) error {
//	var (
//		err    error
//		num    int64
//		Result QouteRetweetResponse
//	)
//
//	req := new(QouteRetweetPostRequest)
//	if err = c.Bind(req); err != nil {
//		return err
//	}
//	if err = ValidateString(req.Qoute, 1, 70); err != nil {
//		return c.JSON(http.StatusBadRequest, err.Error())
//	}
//	if err = c.Validate(req); err != nil {
//		return err
//	}
//	if errNum, _, err = s.AuthAccount(c); err != nil {
//		return c.JSON(errNum, err)
//	}
//
//	num, err = s.store.GetQouteRetweetRows(c.Request().Context(), db2.GetQouteRetweetRowsParams{FromAccountID: req.FromAccountID, PostID: req.PostID})
//	if errNum, err = GetErrorValidator(c, err, Qretweet); err != nil {
//		return c.JSON(errNum, err.Error())
//	}
//
//	if req.IsRetweet {
//		if num != 0 {
//			return c.JSON(http.StatusBadRequest, "already create")
//		}
//
//		Result, err = s.CreateQouteRetweetPost(c.Request().Context(), req)
//		if err != nil {
//			return c.JSON(Result.ErrNum, err.Error())
//		}
//		return c.JSON(http.StatusOK, qouteretweetResponse(Result.Post, Result.PostFeature, req.Qoute))
//	}
//
//	if !req.IsRetweet {
//		if num == 0 {
//			return c.JSON(http.StatusBadRequest, "no retweet")
//		}
//		errNum, err = s.store.DeleteQouteRetweetTX(c.Request().Context(), db.UnRetweetTXParam{
//			FromAccountID: req.FromAccountID,
//			PostID:        req.PostID,
//		})
//		if err != nil {
//			return c.JSON(errNum, err.Error())
//		}
//		return c.JSON(http.StatusCreated, echo.Map{
//			"Delete":    true,
//			"DeletedAt": time.Now().Unix(),
//		})
//	}
//
//	return err
//}
