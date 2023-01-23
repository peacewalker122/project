package request

import (
	"errors"
	"mime/multipart"
)

type UploadFilesRequest struct {
	File          multipart.File
	FileHeader    *multipart.FileHeader
	MaxUploadSize int64
}

func (r *UploadFilesRequest) Validate() error {
	if r.File == nil {
		return errors.New("file is nil")
	}
	if r.FileHeader == nil {
		return errors.New("file header is nil")
	}
	if r.MaxUploadSize == 0 {
		return errors.New("max upload size is 0")
	}
	return nil
}
