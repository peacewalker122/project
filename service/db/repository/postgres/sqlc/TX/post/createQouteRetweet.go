package post

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	"github.com/peacewalker122/project/util"
)

func (p *PostTx) CreateQouteRetweet(ctx context.Context, arg *request.CreateQouteRetweetParams) (*uuid.UUID, error) {
	if arg == nil {
		return nil, errors.New("arg is nil")
	}
	var postID uuid.UUID
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		_, err = q.CreateQouteRetweet_feature(ctx, db.CreateQouteRetweet_featureParams{
			FromAccountID: arg.AccountID,
			PostID:        arg.PostID,
			Qoute:         arg.Qoute,
		})
		if err != nil {
			return err
		}

		isExist, err := q.GetPostQRetweetJoin(ctx, db.GetPostQRetweetJoinParams{
			FromAccountID: arg.AccountID,
			PostID:        arg.PostID,
		})
		if err != nil {
			return err
		}

		if isExist.QouteRetweet {
			return errors.New("already created")
		}

		postData, err := q.GetPost(ctx, arg.PostID)
		if err != nil {
			return err
		}

		postArg := db.CreatePostParams{
			AccountID: arg.AccountID,
			IsRetweet: true,
			PhotoDir:  util.InputSqlString(postData.PhotoDir.String),
		}
		postResult, err := q.CreatePost(ctx, postArg)
		if err != nil {
			return err
		}

		_, err = q.CreatePost_feature(ctx, postResult.PostID)
		if err != nil {
			return err
		}

		postUpdate, err := q.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			return err
		}
		postUpdate.SumQouteRetweet++

		err = q.UpdateQouteRetweet(ctx, db.UpdateQouteRetweetParams{
			QouteRetweet:  true,
			PostID:        arg.PostID,
			FromAccountID: arg.AccountID,
		})
		if err != nil {
			return err
		}

		argEntry := db.CreateEntriesParams{
			FromAccountID: arg.AccountID,
			PostID:        arg.PostID,
			TypeEntries:   db.QR,
		}

		_, err = q.CreateEntries(ctx, argEntry)
		if err != nil {
			return err
		}

		argPostUpdate := db.UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      postUpdate.SumComment,
			SumLike:         postUpdate.SumLike,
			SumRetweet:      postUpdate.SumRetweet,
			SumQouteRetweet: postUpdate.SumQouteRetweet,
		}

		_, err = q.UpdatePost_feature(ctx, argPostUpdate)
		if err != nil {
			return err
		}

		postID = postResult.PostID

		return err
	})
	return &postID, err
}
