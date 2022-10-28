package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/peacewalker122/project/db/sqlc"
)

func (s *Server) deleteQouteRetweet(arg *QouteRetweetPostRequest, c echo.Context, post db.PostFeature) error {
	num, err := s.store.GetPostQRetweetJoin(c.Request().Context(), db.GetPostQRetweetJoinParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
	if err := GetErrorValidator(c, err, qretweet); err != nil {
		return err
	}

	_, err = s.store.GetQouteRetweet(c.Request().Context(), db.GetQouteRetweetParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "no specify qoute-retweet in database")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{
		FromAccountID: arg.FromAccountID,
		PostID:        arg.PostID,
		TypeEntries:   unqretweet,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.store.DeleteQouteRetweet(c.Request().Context(), db.DeleteQouteRetweetParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Delete first then decrement
	post.SumQouteRetweet--
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

	err = s.store.DeletePostFeature(c.Request().Context(), num.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = s.store.DeletePost(c.Request().Context(), num.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"Delete":    true,
		"DeletedAt": time.Now().Unix(),
	})
}

func (s *Server) createQouteRetweetPost(param *QouteRetweetPostRequest, c echo.Context) (db.Post, db.PostFeature, error) {
	arg := db.CreatePostParams{
		AccountID:          param.FromAccountID,
		PictureDescription: param.Qoute,
		IsRetweet:          true,
	}

	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return db.Post{}, db.PostFeature{}, c.JSON(http.StatusInternalServerError, err.Error())
	}

	postfeat, err := s.store.CreatePost_feature(c.Request().Context(), post.PostID)
	if err != nil {
		return db.Post{}, db.PostFeature{}, c.JSON(http.StatusInternalServerError, err.Error())
	}

	return post, postfeat, nil
}

func (s *Server) createRetweetPost(param *RetweetPostRequest, c echo.Context) (db.Post, db.PostFeature, error) {
	arg := db.CreatePostParams{
		AccountID: param.FromAccountID,
		IsRetweet: true,
	}

	post, err := s.store.CreatePost(c.Request().Context(), arg)
	if err != nil {
		return db.Post{}, db.PostFeature{}, c.JSON(http.StatusInternalServerError, err.Error())
	}

	postfeat, err := s.store.CreatePost_feature(c.Request().Context(), post.PostID)
	if err != nil {
		return db.Post{}, db.PostFeature{}, c.JSON(http.StatusInternalServerError, err.Error())
	}

	return post, postfeat, nil
}

func (s *Server) deleteRetweetpost(arg *RetweetPostRequest, c echo.Context, post db.PostFeature) error {
	res, err := s.store.GetPostidretweetJoin(c.Request().Context(), db.GetPostidretweetJoinParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
	if err := GetErrorValidator(c, err, retweet); err != nil {
		return err
	}
	_, err = s.store.GetRetweet(c.Request().Context(), db.GetRetweetParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "no specify qoute-retweet in database")
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = s.store.CreateEntries(c.Request().Context(), db.CreateEntriesParams{
		FromAccountID: arg.FromAccountID,
		PostID:        arg.PostID,
		TypeEntries:   unretweet,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = s.store.DeleteRetweet(c.Request().Context(), db.DeleteRetweetParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Delete first then decrement
	post.SumQouteRetweet--
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

	err = s.store.DeletePostFeature(c.Request().Context(), res.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = s.store.DeletePost(c.Request().Context(), res.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"Delete":    true,
		"DeletedAt": time.Now().Unix(),
	})
}
