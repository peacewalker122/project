package gcp

import (
	"context"
	"fmt"
	"github.com/peacewalker122/project/api/util"
	util2 "github.com/peacewalker122/project/util"
	"time"

	"github.com/peacewalker122/project/service/gcp/request"
	"io"
	"net/url"
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

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	upload := g.Client.Bucket(g.BUCKET).Object(req.FileHeader.Filename).NewWriter(ctx)

	clientFile, err := req.FileHeader.Open()
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(upload, clientFile); err != nil {
		return "", err
	}
	defer clientFile.Close()

	if err := upload.Close(); err != nil {
		return "", err
	}

	baseURl := "https://storage.cloud.google.com"
	URL := fmt.Sprintf("%s/%s/%v", baseURl, g.BUCKET, upload.Attrs().Name)

	fileLink, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	return fileLink.String(), nil
}
