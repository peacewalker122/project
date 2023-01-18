package gcp

import "cloud.google.com/go/storage"

type GCPService interface {
	UploadPhoto(fileName string, file []byte) (string, error)
}

type gcpService struct {
	*storage.Client
}

func NewGCPService(client *storage.Client) GCPService {
	return &gcpService{client}
}
