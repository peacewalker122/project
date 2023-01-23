package file

import (
	"errors"
	"fmt"
	"github.com/h2non/filetype"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func ValidateFileType(src multipart.File) error {
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())
	_, err = io.CopyN(tempFile, src, 216)
	if err != nil {
		return err
	}
	src.Close()

	bytes := make([]byte, 216)

	_, err = tempFile.ReadAt(bytes, 0)
	if err != nil {
		return errors.New("failed to read file: " + err.Error())
	}

	mimeType, err := filetype.Match(bytes)
	log.Println("type: ", mimeType.MIME.Value)

	for _, mime := range MIMEAUTH {
		if mimeType.MIME.Value == mime.String() {
			return nil
		}
	}

	var errs []error
	for _, v := range MIMEAUTH {
		errs = append(errs, errors.New(v.String()))
	}

	return fmt.Errorf("invalid type! must be %v", errs)
}
