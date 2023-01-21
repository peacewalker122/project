package request

import (
	"errors"
	"mime/multipart"
)

type UploadFilesRequest struct {
	File          multipart.File
	FileHeader    *multipart.FileHeader
	BucketName    string
	MaxUploadSize int64
}

func (r *UploadFilesRequest) Validate() error {
	if r.File == nil {
		return errors.New("file is nil")
	}
	if r.FileHeader == nil {
		return errors.New("file header is nil")
	}
	if r.BucketName == "" {
		return errors.New("bucket name is empty")
	}
	if r.MaxUploadSize == 0 {
		return errors.New("max upload size is 0")
	}
	return nil
}
