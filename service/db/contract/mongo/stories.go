package contract

import (
	"context"
	"github.com/peacewalker122/project/service/db/model/mongo"
)

type StoriesRepoitory interface {
	Get(ctx context.Context, id string) (*model.Story, error)
	GetAll(ctx context.Context, usersID []string) ([]*model.Story, error)
	Create(ctx context.Context, story *model.Story) error
	Delete(ctx context.Context, id string) error
}
