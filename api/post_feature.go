package api

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	db "github.com/peacewalker122/project/db/sqlc"
	"github.com/peacewalker122/project/util"
)

var (
	filePath string
	err      error
	isShow   bool
	//errChan      = make(chan error, 1)
	//chanIsShow   = make(chan bool, 1)
	//chanFilePath = make(chan string, 1)
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
	// in this block we created a goroutine to minimize the respond time
	// first we declare our sync using waitgroup
	var wg sync.WaitGroup
	var post db.Post
	var post2 db.PostFeature

	benchTime := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		filePath, err, isShow = s.saveFile(c, arg)
	}()
	wg.Wait()
	if err != nil {
		if !isShow {
			log.Errorf("error in here due: ", err.Error())
			log.Print("benchmark-creating-post: ", time.Since(benchTime))
			return err
		}
		log.Print("benchmark-creating-post: ", time.Since(benchTime))
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	dbArg := db.CreatePostParams{
		AccountID:          arg.AccountID,
		PictureDescription: arg.PictureDescription,
		PhotoDir:           util.InputSqlString(filePath),
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		post, err = s.store.CreatePost(c.Request().Context(), dbArg)
	}()
	wg.Wait()
	log.Print("out: ", post.PostID)
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Print("in: ", post.PostID)
		post2, err = s.store.CreatePost_feature(c.Request().Context(), post.PostID)
	}()
	wg.Wait()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Print("benchmark-creating-post: ", time.Since(benchTime))
	return c.JSON(http.StatusOK, PostResponse(post, post2))
}

// the bool return to indicate a error that will viewed by the client side.
// True = client will see and vice versa.
func (s *Server) saveFile(c echo.Context, arg *CreatePostParams) (string, error, bool) {
	var wg sync.WaitGroup
	var fileName string
	folderPath := fmt.Sprintf("/home/servumtopia/Pictures/Project/%v/", arg.AccountID)
	timeMark := time.Now()
	file, err := c.FormFile("photo")
	if err != nil {
		if err == http.ErrMissingFile {
			return "", nil, false
		}
		return "", err, false
	}

	// maximum size 100MB
	if file.Size > 100000000 {
		log.Print("benchmark: ", time.Since(timeMark))
		return "", errors.New("maximum size: 100MB"), true
	}

	// Here we validate the file is it already in our directory or not
	if _, err = os.Stat(folderPath + file.Filename); err == nil {
		fileName, err = s.store.CreateFileIndex(folderPath, file.Filename)
		if err != nil {
			return "", err, true
		}
		file.Filename = fileName
	}

	if _, err = os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			return "", err, false
		}
	}

	src, err := file.Open()
	if err != nil {
		return "", err, false
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = ValidateFileType(src)
	}()
	wg.Wait()
	if err != nil {
		return "", err, true
	}

	// here we create another file opener to avoid error in copying file
	src.Close()
	src, err = file.Open()
	if err != nil {
		return "", err, false
	}
	defer src.Close()

	filePath := folderPath + file.Filename

	// Destination
	dst, err := os.Create(filepath.Join(folderPath, filepath.Base(file.Filename)))
	if err != nil {
		return "", err, false
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err, false
	}
	log.Print("benchmark: ", time.Since(timeMark))
	return filePath, nil, false
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
