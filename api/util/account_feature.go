package api

import (
	"context"
	"errors"

	"github.com/peacewalker122/project/db/redis"
	db "github.com/peacewalker122/project/db/sqlc"
)

type UtilTools struct {
	store db.Store
	redis redis.Store
}

func NewApiUtil(store db.Store, redis redis.Store) *UtilTools {
	return &UtilTools{
		store: store,
		redis: redis,
	}
}

type (
	CreateQueue struct {
		FromAccountID int64
		ToAccountID   int64
	}
)

func (s *UtilTools) CreateAccountsQueue(ctx context.Context, req *CreateQueue) error {

	ok, err := s.store.CreateAccountsQueueTX(ctx, db.CreateAccountQueueParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
	})
	if err != nil || !ok {
		if !ok {
			err = errors.New("can't proceed queue")
		}
		return err
	}
	return err
}
