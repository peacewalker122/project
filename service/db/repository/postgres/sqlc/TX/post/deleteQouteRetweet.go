package post

import (
	"context"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostTx) DeleteQouteRetweetTX(ctx context.Context, AccountID int64, PostID uuid.UUID) error {
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		qRetweetInfo, err := q.GetPostQRetweetJoin(ctx, db.GetPostQRetweetJoinParams{
			PostID:        PostID,
			FromAccountID: AccountID,
		})
		if err != nil {
			return err
		}

		_, err = q.GetQouteRetweet(ctx, db.GetQouteRetweetParams{
			FromAccountID: AccountID,
			PostID:        PostID,
		})
		if err != nil {
			return err
		}

		_, err = q.CreateEntries(ctx, db.CreateEntriesParams{
			FromAccountID: AccountID,
			PostID:        PostID,
			TypeEntries:   db.UQR,
		})
		if err != nil {
			return err
		}

		err = q.DeleteQouteRetweet(ctx, db.DeleteQouteRetweetParams{
			PostID:        PostID,
			FromAccountID: AccountID,
		})
		if err != nil {
			return err
		}

		post, err := q.GetPost_feature_Update(ctx, PostID)
		if err != nil {
			return err
		}

		// Delete first then decrement
		post.SumQouteRetweet--
		_, err = q.UpdatePost_feature(ctx, db.UpdatePost_featureParams{
			PostID:          PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		err = q.DeletePostFeature(ctx, qRetweetInfo.PostID)
		if err != nil {
			return err
		}
		err = q.DeletePost(ctx, qRetweetInfo.PostID)
		if err != nil {
			return err
		}

		return err
	})
	return err
}
