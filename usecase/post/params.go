package post

import "mime/multipart"

type CreatePostRequest struct {
	File               multipart.File
	FileHeader         *multipart.FileHeader
	AccountID          int64
	PictureDescription string
}
