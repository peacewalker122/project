package post

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	. "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
	"github.com/peacewalker122/project/util"
)

func (p *PostTx) CreatePostGCPTx(ctx context.Context, arg *request.CreatePostParams) (PostTXResult, error) {
	var (
		res PostTXResult
	)

	if arg == nil {
		return res, errors.New("arg is nil")
	}

	err := p.DBnTx(ctx, arg.DelFunc, func(q *db.Queries) (string, error) {
		var err error

		res.FileURL, err = arg.GcpFunc(ctx, arg.FileRequest)
		if err != nil {
			log.Println("gcp error: ", err)
			res.Err = err
			return res.FileURL, err
		}
		select {
		case <-ctx.Done():
			res.Err = ctx.Err()
			return res.FileURL, ctx.Err()
		default:
		}
		log.Println("file url: ", res.FileURL)

		res.Post, err = p.CreatePost(ctx, db.CreatePostParams{
			ID:                 uuid.New(),
			AccountID:          arg.AccountID,
			IsRetweet:          false,
			PictureDescription: arg.PictureDescription,
			PhotoDir:           util.InputSqlString(res.FileURL),
		})
		if err != nil {
			log.Println(err)
			return res.FileURL, err
		}

		res.PostFeature, err = p.CreatePost_feature(ctx, res.Post.ID)
		if err != nil {
			log.Println(err)
			return res.FileURL, err
		}

		return "", err
	})

	if res.Err != nil {
		err = res.Err
	}

	return res, err
}
