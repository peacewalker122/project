package post

import (
	"context"
	"errors"
	"log"

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

	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		res.FileURL, err = arg.GcpFunc(ctx, arg.FileRequest)
		if err != nil {
			log.Println("gcp error: ", err)
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		res.Post, err = p.CreatePost(ctx, db.CreatePostParams{
			AccountID:          arg.AccountID,
			IsRetweet:          false,
			PictureDescription: arg.PictureDescription,
			PhotoDir:           util.InputSqlString(res.FileURL),
		})
		if err != nil {
			return err
		}

		res.PostFeature, err = p.CreatePost_feature(ctx, res.Post.PostID)
		if err != nil {
			return err
		}

		return err
	})
	return res, err
}
