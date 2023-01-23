package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
)

func (p *PostUsecase) CreateQouteRetweet(ctx context.Context, param *QouteRetweetRequest) error {
	if param == nil {
		return errors.New("param is nil")
	}

	qRetweetRow, err := p.postgre.GetQouteRetweetRows(ctx, db.GetQouteRetweetRowsParams{
		FromAccountID: param.AccountID,
		PostID:        param.PostID,
	})
	if err != nil {
		return err
	}

	if qRetweetRow != 0 {
		return errors.New("already exist")
	}

	postID, err := p.postgre.CreateQouteRetweet(ctx, &request.CreateQouteRetweetParams{
		AccountID: param.AccountID,
		PostID:    param.PostID,
		Qoute:     param.Qoute,
	})
	if err != nil || postID == nil {
		return err
	}

	return nil
}
