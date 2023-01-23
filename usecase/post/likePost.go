package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostUsecase) LikePost(ctx context.Context, param *LikeRequest) error {
	if param == nil {
		return errors.New("param is nil")
	}

	likeCount, err := p.postgre.GetLikeRows(ctx, db.GetLikeRowsParams{
		Fromaccountid: param.AccountID,
		Postid:        param.PostID,
	})

	if likeCount == 0 {
		err = p.postgre.CreateLike_feature(ctx, db.CreateLike_featureParams{
			PostID:        param.PostID,
			FromAccountID: param.AccountID,
			IsLike:        true,
		})
		if err != nil {
			return err
		}
	}

	isLike, err := p.postgre.GetLikejoin(ctx, param.PostID)
	if err != nil {
		return err
	}

	if isLike {
		return errors.New("already like")
	}

	_, err = p.postgre.CreateLike(ctx, param.PostID, param.AccountID)
	if err != nil {
		return err
	}

	return nil
}
