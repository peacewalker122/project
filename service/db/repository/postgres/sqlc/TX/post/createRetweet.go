package post

import (
	"context"
	"database/sql"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
	"github.com/peacewalker122/project/util"
	"net/http"
)

func (p *PostTx) CreateRetweetTX(ctx context.Context, arg *request.CreateRetweetParams) (result.RetweetTXResult, error) {
	var result result.RetweetTXResult

	if arg == nil {
		return result, errors.New("arg is nil")
	}

	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error
		result.ErrCode = http.StatusInternalServerError

		err = q.CreateRetweet_feature(ctx, db.CreateRetweet_featureParams{FromAccountID: arg.AccountID, PostID: arg.PostID})
		if err != nil {
			return err
		}

		ok, err := q.GetPostidretweetJoin(ctx, db.GetPostidretweetJoinParams{FromAccountID: arg.AccountID, PostID: arg.PostID})
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		if ok.Retweet {
			return errors.New("already created")
		}

		post, err := q.GetPost(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		postArg := db.CreatePostParams{
			AccountID: arg.AccountID,
			IsRetweet: true,
			PhotoDir:  util.InputSqlString(post.PhotoDir.String),
		}
		result.Post, err = q.CreatePost(ctx, postArg)
		if err != nil {
			return err
		}

		result.PostFeature, err = q.CreatePost_feature(ctx, result.Post.PostID)
		if err != nil {
			return err
		}

		isRetweetExist, err := q.GetRetweetJoin(ctx, db.GetRetweetJoinParams{Postid: arg.PostID, Fromaccountid: arg.AccountID})
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		if isRetweetExist {
			result.ErrCode = http.StatusNotFound
			result.Err = errors.New("already retweet")
			return result.Err
		}

		args := db.CreateEntriesParams{
			FromAccountID: arg.AccountID,
			PostID:        arg.PostID,
			TypeEntries:   db.R,
		}

		go q.CreateEntries(ctx, args)

		updatePost, err := q.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			result.ErrCode = 500
			return err
		}
		updatePost.SumRetweet++

		err = q.UpdateRetweet(ctx, db.UpdateRetweetParams{Retweet: true, PostID: arg.PostID, FromAccountID: arg.AccountID})
		if err != nil {
			result.ErrCode = 500
			return err
		}

		arg := db.UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      updatePost.SumComment,
			SumLike:         updatePost.SumLike,
			SumRetweet:      updatePost.SumRetweet,
			SumQouteRetweet: updatePost.SumQouteRetweet,
		}

		_, err = q.UpdatePost_feature(ctx, arg)
		if err != nil {
			result.ErrCode = 500
			return err
		}
		return err
	})

	return result, err
}
