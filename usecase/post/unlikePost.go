package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostUsecase) UnLikePost(ctx context.Context, param *LikeRequest) error {
	if param == nil {
		return errors.New("param is nil")
	}

	likeCount, err := p.postgre.GetLikeRows(ctx, db.GetLikeRowsParams{
		Fromaccountid: param.AccountID,
		Postid:        param.PostID,
	})
	if err != nil {
		return err
	}

	if likeCount != 1 {
		return errors.New("not like")
	}

	isLike, err := p.postgre.GetLikejoin(ctx, param.PostID)
	if err != nil {
		return err
	}

	if !isLike {
		return errors.New("not like")
	}

	_, err = p.postgre.UnlikeTX(ctx, param.PostID, param.AccountID)
	if err != nil {
		return err
	}

	return nil
}
