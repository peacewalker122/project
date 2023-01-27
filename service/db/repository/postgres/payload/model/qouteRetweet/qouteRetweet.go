package qouteRetweet

import (
	"context"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/db/repository/postgres/ent"
	qRetweet "github.com/peacewalker122/project/service/db/repository/postgres/ent/qoute_retweet_feature"
)

type QouteRetweetQuery interface {
	GetQouteRetweetByPost(ctx context.Context, postID uuid.UUID) ([]*ent.Qoute_retweet_feature, error)
}

type QouteRetweetQueries struct {
	*ent.Client
}

func NewQouteRetweetQuery(client *ent.Client) *QouteRetweetQueries {
	return &QouteRetweetQueries{
		Client: client,
	}
}

func (q *QouteRetweetQueries) GetQouteRetweetByPost(ctx context.Context, postID uuid.UUID) ([]*ent.Qoute_retweet_feature, error) {
	qouteRetweet, err := q.Qoute_retweet_feature.
		Query().
		Where(qRetweet.PostID(postID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return qouteRetweet, nil
}
