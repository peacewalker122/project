package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
)

func (p *PostUsecase) RetweetPost(ctx context.Context, param *RetweetRequest) error {
	if param == nil {
		return errors.New("param is nil")
	}

	// Check if the post exists
	postCount, err := p.postgre.GetRetweetRows(ctx, db.GetRetweetRowsParams{
		FromAccountID: param.AccountID,
		PostID:        param.PostID,
	})
	if err != nil {
		return err
	}

	if postCount != 0 {
		return errors.New("already retweet")
	}

	_, err = p.postgre.CreateRetweetTX(ctx, &request.CreateRetweetParams{
		AccountID: param.AccountID,
		PostID:    param.PostID,
	})

	if err != nil {
		return err
	}

	return nil
}
