package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
)

func (p *PostTx) CreateComment(ctx context.Context, param *request.CreateCommentParams) error {
	if param == nil {
		return errors.New("param is nil")
	}
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		_, err = q.CreateComment_feature(ctx, db.CreateComment_featureParams{FromAccountID: param.AccountID, Comment: param.Comment, PostID: param.PostID})
		if err != nil {
			return err
		}

		post, err := q.GetPost_feature_Update(ctx, param.PostID)
		if err != nil {
			return err
		}
		post.SumComment++

		_, err = q.CreateEntries(ctx, db.CreateEntriesParams{FromAccountID: param.AccountID, PostID: param.PostID, TypeEntries: db.C})
		if err != nil {
			return err
		}

		_, err = q.UpdatePost_feature(ctx, db.UpdatePost_featureParams{
			PostID:          param.PostID,
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
	return err
}
