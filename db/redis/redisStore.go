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
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys string) error
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

func (r *RedisStore) Set(ctx context.Context, key string, value interface{}) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.redis.Set(ctx, key, json, 15*time.Minute).Err()
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

func (r *RedisStore) Del(ctx context.Context, keys string) error {
	_, err := r.redis.Del(ctx, keys).Result()
	if err != nil {
		return err
	}
	return nil
}
