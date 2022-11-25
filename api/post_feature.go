package api

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
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

	err := s.store.CreateRetweet_feature(c.Request().Context(), db.CreateRetweet_featureParams{FromAccountID: param.FromAccountID, PostID: param.PostID})
	if err := CreateErrorValidator(c, err); err != nil {
		return db.Post{}, db.PostFeature{}, err
	}

	ok, err := s.store.GetPostidretweetJoin(c.Request().Context(), db.GetPostidretweetJoinParams{FromAccountID: param.FromAccountID, PostID: param.PostID})
	if err := GetErrorValidator(c, err, qretweet); err != nil {
		return db.Post{}, db.PostFeature{}, err
	}
	if !ok.Retweet {
		post, err := s.store.CreatePost(c.Request().Context(), arg)
		if err != nil {
			return db.Post{}, db.PostFeature{}, c.JSON(http.StatusInternalServerError, err.Error())
		}

		postfeat, err := s.store.CreatePost_feature(c.Request().Context(), post.PostID)
		if err != nil {
			return db.Post{}, db.PostFeature{}, c.JSON(http.StatusInternalServerError, err.Error())
		}

		return post, postfeat, nil
	} else {
		return db.Post{}, db.PostFeature{}, errors.New("already existing retweet")
	}

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

func (s *Server) creatingPost(c echo.Context, arg *CreatePostParams) error {
	filePath, err := s.saveFile(c)
	if err != nil {
		log.Errorf("error in here due: ", err.Error())
		return err
	}

	dbArg := db.CreatePostParams{
		AccountID:          arg.AccountID,
		PictureDescription: arg.PictureDescription,
		PhotoDir:           util.InputSqlString(filePath),
	}

	post, err := s.store.CreatePost(c.Request().Context(), dbArg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	post2, err := s.store.CreatePost_feature(c.Request().Context(), post.PostID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, PostResponse(post, post2))
}

func (s *Server) saveFile(c echo.Context) (string, error) {
	file, err := c.FormFile("photo")
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("error in open: %v", err.Error())
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join("/home/servumtopia/Pictures/Project", filepath.Base(file.Filename)))
	if err != nil {
		return "", fmt.Errorf("error in create: %v", err.Error())
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	filePath := fmt.Sprintf("/home/servumtopia/Pictures/Project/%v", file.Filename)
	if file.Filename == "" {
		// to ensure consistency of file.
		filePath = ""
	}
	return filePath, nil
}

func (s *Server) GettingPost(c echo.Context, req *GetPostParam) error {
	post, err := s.store.GetPost(c.Request().Context(), int64(req.postID))
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}

	postFeature, err := s.store.GetPost_feature(c.Request().Context(), int64(req.postID))
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}

	arg := db.ListCommentParams{PostID: int64(req.postID), Limit: int32(10), Offset: (req.Offset - 1) * 10}
	comment, err := s.store.ListComment(c.Request().Context(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(http.StatusOK, GetPostResponse(post, postFeature, comment), "\t")
}

func (s *Server) gettingImage(c echo.Context, postID int64) error {
	post, err := s.store.GetPost(c.Request().Context(), int64(postID))
	if err := GetErrorValidator(c, err, posttag); err != nil {
		return err
	}
	return c.File(post.PhotoDir.String)
}
