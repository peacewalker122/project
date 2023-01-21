package gcp

import (
	"context"
	"fmt"
	"github.com/peacewalker122/project/api/util"
	util2 "github.com/peacewalker122/project/util"

	"github.com/peacewalker122/project/service/gcp/request"
	"io"
	"net/url"
	"time"
)

func (g *gcpService) UploadPhoto(ctx context.Context, req *request.UploadFilesRequest) (string, error) {

	var err error

	err = req.Validate()
	if err != nil {
		return "", err
	}

	req.FileHeader.Filename, err = util2.RandomFileName(req.FileHeader.Filename)
	if err != nil {
		return "", err
	}

	if req.FileHeader.Size > req.MaxUploadSize {
		return "", fmt.Errorf("file size is too large")
	}

	err = util.ValidateFileType(req.File)
	if err != nil {
		return "", err
	}

	upload := g.Client.Bucket(req.BucketName).Object(req.FileHeader.Filename).NewWriter(ctx)

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	if _, err := io.Copy(upload, req.File); err != nil {
		return "", err
	}

	if err := upload.Close(); err != nil {
		return "", err
	}

	URL := fmt.Sprintf("/%s/%v", req.BucketName, upload.Attrs())

	fileLink, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	return fileLink.String(), nil
}
