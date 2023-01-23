package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostUsecase) UnRetweetPost(ctx context.Context, param *RetweetRequest) error {
	if param == nil {
		return errors.New("param is nil")
	}

	postCount, err := p.postgre.GetRetweetRows(ctx, db.GetRetweetRowsParams{
		FromAccountID: param.AccountID,
		PostID:        param.PostID,
	})
	if err != nil {
		return err
	}

	if postCount != 1 {
		return errors.New("not found retweet")
	}

	err = p.postgre.DeleteRetweetTX(ctx, param.PostID, param.AccountID)
	if err != nil {
		return err
	}

	return nil
}
