package util

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
)

type MIME string

var (
	MIME_JPG  MIME = "image/jpg"
	MIME_JPEG MIME = "image/jpeg"
	MIME_GIF  MIME = "image/gif"
	MIME_PNG  MIME = "image/png"
	MIME_WEBP MIME = "image/webp"
	MIME_MP4  MIME = "video/mp4"

	MIMEAUTH = []MIME{
		MIME_JPEG,
		MIME_JPG,
		MIME_GIF,
		MIME_PNG,
		MIME_WEBP,
		MIME_MP4,
	}
)

func (m MIME) String() string {
	return string(m)
}

func ValidateFileType(input multipart.File) error {
	byte := make([]byte, 512)
	if _, err := input.Read(byte); err != nil {
		return err
	}
	file := http.DetectContentType(byte)

	for _, v := range MIMEAUTH {
		if file == v.String() {
			return nil
		}
	}

	var err []error
	for _, v := range MIMEAUTH {
		err = append(err, errors.New(v.String()))
	}

	return fmt.Errorf("invalid type! must be %v", err)
}
