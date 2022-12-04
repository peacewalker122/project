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

type (
	PostResponses struct {
		Post        db.Post
		PostFeature db.PostFeature
		CommentList []db.ListCommentRow
	}
)

var (
	filePath string
	isShow   bool
	FileName string
	//errChan      = make(chan error, 1)
	//chanIsShow   = make(chan bool, 1)
	//chanFilePath = make(chan string, 1)
)

func (s *Handler) DeleteQouteRetweet(arg *QouteRetweetPostRequest, c echo.Context, post db.GetPost_feature_UpdateRow) error {

	num, err := s.store.GetPostQRetweetJoin(c.Request().Context(), db.GetPostQRetweetJoinParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
	if errNum, err = GetErrorValidator(c, err, Qretweet); err != nil {
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
		TypeEntries:   Unqretweet,
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
		PostID:          arg.PostID,
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

func (s *Handler) CreateQouteRetweetPost(param *QouteRetweetPostRequest, c echo.Context) (db.Post, db.PostFeature, error) {
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

func (s *Handler) CreateRetweetPost(c echo.Context, param *RetweetPostRequest) (RetweetResponse, int, error) {
	arg := db.CreateRetweetParams{
		FromAccountID: param.FromAccountID,
		PostID:        param.PostID,
		IsRetweet:     true}

	result, err := s.store.CreateRetweetPost(c.Request().Context(), arg)
	if err != nil {
		return RetweetResponse{}, result.ErrCode, err
	}
	res, err := s.store.CreateRetweetTX(c.Request().Context(), arg)
	if err != nil {
		return RetweetResponse{}, res.ErrCode, err
	}

	return RetweetResponse{Post: result.Post.Post, Feature: result.Post.PostFeature}, 0, err

}

func (s *Handler) DeleteRetweetpost(arg *RetweetPostRequest, c echo.Context) error {
	err := s.store.DeleteRetweetTX(c.Request().Context(), db.DeleteRetweetParams{
		PostID:        arg.PostID,
		FromAccountID: arg.FromAccountID,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"Delete":    true,
		"DeletedAt": time.Now().Unix(),
	})
}

func (s *Handler) CreatingPost(c echo.Context, arg *CreatePostParams) error {
	var wg sync.WaitGroup
	var post db.PostTXResult

	// in this block we created a goroutine to minimize the respond time
	// first step we need declare our sync using waitgroup
	benchTime := time.Now()
	wg.Add(1)
	go func() {
		defer wg.Done()
		filePath, err, isShow = s.SaveFile(c, arg)
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

	// goroutine to maximize our code response
	// here we invoke goroutine to minimize time for writing in the database.
	wg.Add(1)
	go func() {
		ctx := c.Request().Context()
		defer wg.Done()
		post, err = s.store.CreatePostTx(ctx, dbArg)
	}()
	wg.Wait()
	// goroutine waiting until creatingpost_feature is done

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Print("benchmark-creating-post: ", time.Since(benchTime))
	return c.JSON(http.StatusOK, PostResponse(post.Post, post.PostFeature))
}

func (s *Handler) SaveFile(c echo.Context, arg *CreatePostParams) (string, error, bool) {
	// the bool return to indicate a error that will viewed by the client side.
	// True = client will see and vice versa.
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

	// maximum size 100MB
	if file.Size > 100000000 {
		return "", errors.New("maximum size: 100MB"), true
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
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err = io.Copy(dst, src)
	}()
	wg.Wait()

	if err != nil {
		return "", err, false
	}
	log.Print("benchmark: ", time.Since(timeMark))
	FileName = filePath
	return filePath, nil, false
}

func (s *Handler) GettingPost(c echo.Context, req *GetPostParam) error {
	result := &PostResponses{}

	result.Post, err = s.store.GetPost(c.Request().Context(), int64(req.PostID))
	if errNum, err = GetErrorValidator(c, err, Posttag); err != nil {
		return c.JSON(errNum, err.Error())
	}

	result.PostFeature, err = s.store.GetPost_feature(c.Request().Context(), int64(req.PostID))
	if errNum, err = GetErrorValidator(c, err, Posttag); err != nil {
		return c.JSON(errNum, err.Error())
	}

	arg := db.ListCommentParams{PostID: int64(req.PostID), Limit: int32(10), Offset: (req.Offset - 1) * 10}
	result.CommentList, err = s.store.ListComment(c.Request().Context(), arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, err.Error())
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSONPretty(http.StatusOK, GetPostResponse(result.Post, result.PostFeature, result.CommentList), "\t")
}

func (s *Handler) GettingImage(c echo.Context, postID int64) error {
	post, err := s.store.GetPost(c.Request().Context(), int64(postID))
	if errNum, err = GetErrorValidator(c, err, Posttag); err != nil {
		return err
	}

	return c.File(post.PhotoDir.String)
}
