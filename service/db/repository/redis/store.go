package redis

import (
	"context"
	"time"
)

type Store interface {
	// in set we marshal the value to json
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	// in setOne we don't marshal the value to json
	SetOne(ctx context.Context, key string, value interface{}, duration time.Duration) error
	// need to unmarshal the value from json
	Get(ctx context.Context, key string) (string, error)
	// append to the key
	Append(ctx context.Context, key string, value interface{}) error

	// get the value and delete the key
	GetRedisPayload(ctx context.Context, key string, payload interface{}) error

	Del(ctx context.Context, key string) error
}
