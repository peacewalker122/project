package api

import (
	"errors"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
)

type utilTools struct {
	store db.Store
	redis redis.Store
}

type UtilTools interface {
	accountFeature
}

func NewApiUtil(store db.Store, redis redis.Store) UtilTools {
	return &utilTools{
		store: store,
		redis: redis,
	}
}

func ValidateFileType(input multipart.File) error {
	byte := make([]byte, 512)
	if _, err := input.Read(byte); err != nil {
		return err
	}
	file := http.DetectContentType(byte)

	s := log.Default()
	s.Print(file)

	if file == "image/jpg" || file == "image/jpeg" || file == "image/gif" || file == "image/png" || file == "image/webp" || file == "video/mp4" {
		return nil
	}
	return errors.New("invalid type! must be jpeg/jpg/gif/png/mp4")
}

