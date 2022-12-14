package redis

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
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
	Del(ctx context.Context, key string) error
}

type RedisStore struct {
	redis *redis.Client
}

func NewRedis(URL string) Store {
	opt, err := redis.ParseURL(URL)
	if err != nil {
		log.Fatal(err)
	}
	rdb := redis.NewClient(opt)

	return &RedisStore{rdb}
}

func (r *RedisStore) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.redis.Set(ctx, key, json, duration).Err()
	return err
}
func (r *RedisStore) Get(ctx context.Context, key string) (string, error) {
	res, err := r.redis.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return "", errors.New("no key found")
		}
		return "", err
	}

	return res, nil
}

func (r *RedisStore) Del(ctx context.Context, key string) error {
	_, err := r.redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisStore) SetOne(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	err := r.redis.Set(ctx, key, value, duration).Err()
	return err
}

func (r *RedisStore) Append(ctx context.Context, key string, value interface{}) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.redis.Append(ctx, key, string(json)).Err()
	return err
}
