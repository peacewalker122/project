package util

import (
	"context"
	"errors"
	"fmt"
	db2 "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/peacewalker122/project/util"
)

const (
	postTag      = "post"
	profilephoto = "profilephoto"
)

type accountFeature interface {
	CreateAccountsQueue(ctx context.Context, req *CreateQueue) error
	UpdateProfilePhoto(c echo.Context, accountid int64) (int, error)
}

type (
	CreateQueue struct {
		FromAccountID int64
		ToAccountID   int64
	}
)

func (s *utilTools) CreateAccountsQueue(ctx context.Context, req *CreateQueue) error {

	ok, err := s.store.CreateAccountsQueueTX(ctx, db2.CreateAccountQueueParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
	})
	if err != nil || !ok {
		if !ok {
			err = errors.New("can't proceed queue")
		}
		return err
	}
	return err
}

func (s *utilTools) UpdateProfilePhoto(c echo.Context, accountid int64) (int, error) {
	ctx := c.Request().Context()
	file, err, ok := s.SaveFile(c, profilephoto, accountid)
	if err != nil {
		if !ok {
			return http.StatusInternalServerError, err
		}
		return http.StatusBadRequest, err
	}

	err = s.store.UpdatePhoto(ctx, db2.UpdatePhotoParams{
		Filedirectory: util.InputSqlString(file),
		Accountid:     accountid,
	})
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, err
}

func (s *utilTools) SaveFile(c echo.Context, PhotoType string, AccountID int64) (string, error, bool) {
	// the bool return to indicate an error that will be view by the client side.
	// True = client will see and vice versa.

	// PhotoType indicate in what folder this will save.
	// only accept ProfilePhoto & PostPhoto
	// will update soon

	// validate photoype if it doesn't recognise then throw error
	if PhotoType != postTag && PhotoType != profilephoto {
		return "", fmt.Errorf("must be either %v or %v", postTag, profilephoto), true
	}

	var wg sync.WaitGroup
	var fileName string
	folderPath := fmt.Sprintf("/home/danielputra/Pictures/Project/%v/", AccountID)

	// here we invoke if it's a profile photo then create a new folder if it's doesn't exist.
	if PhotoType == profilephoto {
		folderPath = fmt.Sprintf("/home/danielputra/Pictures/Project/%v/%v/", AccountID, "profile")
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
		fileName, err = util.RandomFileName(file.Filename)
		if err != nil {
			return "", err, true
		}
		file.Filename = fileName
	}

	// here check the file folder already exist or not
	// if not then create the directory
	if _, err = os.Stat(folderPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return "", err, false
		}
	}

	src, err := file.Open()
	if err != nil {
		return "", err, false
	}

	done := make(chan bool)
	errchan := make(chan error)
	go func() {
		err = ValidateFileType(src)
		if err != nil {
			errchan <- err
			return
		}
		done <- true
	}()
	select {
	case <-done:
	case err = <-errchan:
		return "", err, true
	}

	// maximum size 2MB
	if file.Size > 2000000 {
		return "", errors.New("maximum size: 2MB"), true
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
		os.Remove(filePath)
		return "", err, false
	}
	wg.Wait()

	return filePath, nil, false
}
