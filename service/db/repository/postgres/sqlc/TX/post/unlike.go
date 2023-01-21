package post

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
	"net/http"
)

func (p *PostTx) UnlikeTX(ctx context.Context, postID uuid.UUID, accountID int64) (result.LikeTXResult, error) {
	var res result.LikeTXResult
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		post, err := q.GetPost_feature_Update(ctx, postID)
		if err != nil {
			if err != sql.ErrNoRows {
				res.ErrCode = http.StatusNotFound
			}
			return err
		}
		post.SumLike--

		err = q.UpdateLike(ctx, db.UpdateLikeParams{
			IsLike:        false,
			PostID:        postID,
			FromAccountID: accountID,
		})
		if err != nil {
			res.ErrCode = 500
			return err
		}

		_, err = q.CreateEntries(ctx, db.CreateEntriesParams{
			FromAccountID: accountID,
			PostID:        postID,
			TypeEntries:   db.UL,
		})
		if err != nil {
			return err
		}

		res.PostFeature, err = q.UpdatePost_feature(ctx, db.UpdatePost_featureParams{
			PostID:          postID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		return err
	})
	return res, err
}
