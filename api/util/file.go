package util

import (
	"errors"
	"fmt"
	file2 "github.com/peacewalker122/project/util/file"
	"mime/multipart"
	"net/http"
)

func ValidateFileType(input multipart.File) error {
	byte := make([]byte, 512)
	if _, err := input.Read(byte); err != nil {
		return err
	}
	file := http.DetectContentType(byte)

	for _, v := range file2.MIMEAUTH {
		if file == v.String() {
			return nil
		}
	}

	var err []error
	for _, v := range file2.MIMEAUTH {
		err = append(err, errors.New(v.String()))
	}

	return fmt.Errorf("invalid type! must be %v", err)
}
