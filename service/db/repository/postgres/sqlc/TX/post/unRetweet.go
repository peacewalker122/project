package post

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostTx) UnRetweet(ctx context.Context, PostID uuid.UUID, AccountID int64) error {
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		retweetPostRelated, err := q.GetPostidretweetJoin(ctx, db.GetPostidretweetJoinParams{
			PostID: PostID, FromAccountID: AccountID,
		})
		if err != nil {
			return err
		}
		post, err := q.GetPost_feature_Update(ctx, PostID)
		if err != nil {
			return err
		}
		_, err = q.GetRetweet(ctx, db.GetRetweetParams{FromAccountID: AccountID, PostID: PostID})
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("no specify qoute-retweet in database")
			}
			return err
		}

		_, err = q.CreateEntries(ctx, db.CreateEntriesParams{
			FromAccountID: AccountID,
			PostID:        PostID,
			TypeEntries:   db.UR,
		})
		if err != nil {
			return err
		}

		err = q.DeleteRetweet(ctx, db.DeleteRetweetParams{PostID: PostID, FromAccountID: AccountID})
		if err != nil {
			return err
		}

		// Delete first then decrement
		post.SumRetweet--
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

		err = q.DeletePostFeature(ctx, retweetPostRelated.PostID)
		if err != nil {
			return err
		}
		err = q.DeletePost(ctx, retweetPostRelated.PostID)
		if err != nil {
			return err
		}
		return err
	})
	return err
}
