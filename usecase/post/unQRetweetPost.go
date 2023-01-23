package post

import (
	"context"
	"errors"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostUsecase) UnQouteRetweet(ctx context.Context, param *RetweetRequest) error {
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

	err = p.postgre.DeleteQouteRetweetTX(ctx, param.AccountID, param.PostID)
	if err != nil {
		return err
	}

	return nil
}
