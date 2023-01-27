package retweet_feature

import (
	"context"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent/retweet_feature"
)

type RetweetQuery interface {
	GetRetweetByPost(ctx context.Context, postID uuid.UUID) ([]*ent.Retweet_feature, error)
}

type RetweetQueries struct {
	*ent.Client
}

func NewRetweetQuery(client *ent.Client) *RetweetQueries {
	return &RetweetQueries{
		Client: client,
	}
}

func (r *RetweetQueries) GetRetweetByPost(ctx context.Context, postID uuid.UUID) ([]*ent.Retweet_feature, error) {
	retweetFeature, err := r.Retweet_feature.
		Query().
		Where(retweet_feature.PostID(postID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return retweetFeature, nil
}
