package db

import (
	"context"

	"cloud.google.com/go/storage"
)

type StorageModel struct {
	Storage    Image
	Bucketname string
}

type Image interface {
	Bucket(name string) *storage.BucketHandle
	Buckets(ctx context.Context, projectID string) *storage.BucketIterator
	Close() error
}

func NewStorageModel(storage Image,bucketname string) *StorageModel {
	return &StorageModel{
		Storage: storage,
		Bucketname: bucketname,
	}
}

