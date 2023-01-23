package post

import (
	"context"
	"errors"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	"github.com/peacewalker122/project/util"
)

func (p *PostUsecase) CreateComment(ctx context.Context, param *CommentRequest) (err *util.MultiError) {
	if param == nil {
		err.Add(errors.New("param is nil"))
		return
	}

	err = p.postgre.CreateComment(ctx, &request.CreateCommentParams{
		AccountID: param.AccountID,
		PostID:    param.PostID,
		Comment:   param.Comment,
	})
	if err != nil {
		return err
	}

	return nil
}
