package likefeature

import (
	"context"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/likefeature"
)

type LikeQuery interface {
	GetLikesByPost(ctx context.Context, postID uuid.UUID) ([]*ent.LikeFeature, error)
	DeleteLike(ctx context.Context, postID uuid.UUID, fromAccountID int64) error
}

type LikeQueries struct {
	*ent.Client
}

func NewLikeQuery(client *ent.Client) *LikeQueries {
	return &LikeQueries{
		client,
	}
}

func (l *LikeQueries) GetLikesByPost(ctx context.Context, postID uuid.UUID) ([]*ent.LikeFeature, error) {
	likeData, err := l.LikeFeature.
		Query().
		Where(likefeature.PostID(postID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return likeData, nil
}

func (l *LikeQueries) DeleteLike(ctx context.Context, postID uuid.UUID, fromAccountID int64) error {
	_, err := l.LikeFeature.
		Delete().
		Where(likefeature.PostID(postID), likefeature.FromAccountID(fromAccountID)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
