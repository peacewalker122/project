package post

import (
	"errors"
	"github.com/google/uuid"
	"mime/multipart"
)

type CreatePostRequest struct {
	File               multipart.File
	FileHeader         *multipart.FileHeader
	AccountID          int64
	PictureDescription string
}
type LikeRequest struct {
	AccountID int64
	PostID    uuid.UUID
}
type RetweetRequest struct {
	AccountID int64
	PostID    uuid.UUID
}
type CommentRequest struct {
	AccountID int64
	PostID    uuid.UUID
	Comment   string
}

func (c *CommentRequest) Validate() error {
	if c.AccountID == 0 {
		return errors.New("account id is required")
	}
	if c.PostID == uuid.Nil {
		return errors.New("post id is required")
	}
	if c.Comment == "" {
		return errors.New("comment is required")
	}
	return nil
}

type QouteRetweetRequest struct {
	AccountID int64
	PostID    uuid.UUID
	Qoute     string
}
