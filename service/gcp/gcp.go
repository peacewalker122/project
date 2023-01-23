package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/peacewalker122/project/service/gcp/request"
)

type GCPService interface {
	UploadPhoto(ctx context.Context, req *request.UploadFilesRequest) (string, error)
	DeleteFile(ctx context.Context, Object string) error
}

type gcpService struct {
	*storage.Client
	BUCKET string
}

func NewGCPService(client *storage.Client, bucket string) GCPService {
	return &gcpService{
		Client: client,
		BUCKET: bucket,
	}
}
