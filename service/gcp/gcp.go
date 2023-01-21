package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/peacewalker122/project/service/gcp/request"
)

type GCPService interface {
	UploadPhoto(ctx context.Context, req *request.UploadFilesRequest) (string, error)
}

type gcpService struct {
	*storage.Client
}

func NewGCPService(client *storage.Client) GCPService {
	return &gcpService{client}
}
