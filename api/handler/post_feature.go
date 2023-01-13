package api

import (
	"context"
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

	db "github.com/peacewalker122/project/db/repository/postgres/sqlc"
	apiutil "github.com/peacewalker122/project/api/util"
	"github.com/peacewalker122/project/util"
)

type (
	PostResponses struct {
		Post        db.Post
		PostFeature db.PostFeature
		CommentList []db.ListCommentRow
	}
	QouteRetweetResponse struct {
		Post        db.Post
		PostFeature db.PostFeature
		ErrNum      int
	}
)

const (
	postTag      = "post"
	profilephoto = "profilephoto"
)

var (
	filePath string
	isShow   bool
	FileName string
	//errChan      = make(chan error, 1)
	//chanIsShow   = make(chan bool, 1)
	//chanFilePath = make(chan string, 1)
)

func (s *Handler) CreateLike(ctx context.Context, arg *LikePostRequest) (int, error) {
	err = s.store.CreateLike_feature(ctx, db.CreateLike_featureParams{
		FromAccountID: arg.FromAccountID,
		IsLike:        false,
		PostID:        arg.PostID,
	})
	if err != nil {
		return 500, err
	}
	return 0, nil
}

func (s *Handler) CreateQouteRetweetPost(c context.Context, param *QouteRetweetPostRequest) (QouteRetweetResponse, error) {
	var result QouteRetweetResponse
	res, err := s.store.CreateQouteRetweetPostTX(c, db.CreateQRetweetParams{
		FromAccountID: param.FromAccountID,
		PostID:        param.PostID,
		Qoute:         param.Qoute,
	})
	if err != nil {
		result.ErrNum = res.ErrCode
		return result, err
	}

	errNum, err = s.store.CreateQouteRetweet(c, db.CreateQRetweetParams{
		FromAccountID: param.FromAccountID,
		PostID:        param.PostID,
	})
	if err != nil {
		result.ErrNum = errNum
		return result, err
	}

	return QouteRetweetResponse{Post: res.Post, PostFeature: res.PostFeature}, nil
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
		filePath, err, isShow = s.SaveFile(c, postTag, arg.AccountID)
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

func (s *Handler) SaveFile(c echo.Context, PhotoType string, AccountID int64) (string, error, bool) {
	// the bool return to indicate a error that will viewed by the client side.
	// True = client will see and vice versa.

	// PhotoType indicate in what folder this will save.
	// only accept ProfilePhoto & PostPhoto
	// will updated soon

	// validate photoype if it doesn't recognise then throw error
	if PhotoType != postTag && PhotoType != profilephoto {
		return "", fmt.Errorf("must be either %v or %v", postTag, profilephoto), true
	}

	var wg sync.WaitGroup
	var fileName string
	folderPath := fmt.Sprintf("%s/%v/", s.config.UserDir, AccountID)

	// here we invoke if it's a profile photo then create a new folder if it's doesn't exist.
	if PhotoType == profilephoto {
		folderPath = fmt.Sprintf("%s/%v/%v/", s.config.UserDir, "profile", AccountID)
	}

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

	// here check the file folder already exist or not
	// if not then create the directory
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

	errchan := make(chan error)
	done := make(chan bool)
	go func() {
		err = apiutil.ValidateFileType(src)
		if err == nil {
			done <- true
			return
		}
		errchan <- err
	}()
	select {
	case err = <-errchan:
		return "", err, true
	case <-time.After(5 * time.Second):
		return "", errors.New("timeout"), true
	case <-done:
	}

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

	// destination
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

	if err != nil {
		return "", err, false
	}

	FileName = filePath
	wg.Wait()
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
