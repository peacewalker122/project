package post

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	"github.com/peacewalker122/project/util"
	log "github.com/sirupsen/logrus"
)

func (p *PostTx) CreateComment(ctx context.Context, param *request.CreateCommentParams) *util.MultiError {
	multiErr := &util.MultiError{Errors: make([]*util.Error, 0)}
	if param == nil {
		multiErr.Add(errors.New("param is nil"))
		return multiErr
	}
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var err error

		_, err = q.CreateComment_feature(ctx, db.CreateComment_featureParams{
			CommentID:     uuid.New(),
			FromAccountID: param.AccountID,
			Comment:       param.Comment,
			PostID:        param.PostID,
		})
		if err != nil {
			if pqerr, ok := err.(*pq.Error); ok {
				switch pqerr.Code.Name() {
				// case for the accountID diminished due comment can happen when acc exist
				case "foreign_key_violation":
					fKeyErr := errors.New("postID or accountID not exist")
					log.Println("error in CreateComment: ", fKeyErr)
					multiErr.Add(fKeyErr)
				}
			}
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
	if err != nil {
		log.Println("error in CreateComment: ", err)
		multiErr.Add(err)
	}

	log.Println("CreateComment: ", multiErr.Error())
	return multiErr
}
