package post

import (
	"context"
	"errors"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
)

func (p *PostUsecase) DeletePost(ctx context.Context, AccountID int64, PostID uuid.UUID) error {

	qRetweetRow, err := p.postgre.GetQouteRetweetRows(ctx, db.GetQouteRetweetRowsParams{
		FromAccountID: AccountID,
		PostID:        PostID,
	})
	if err != nil {
		return err
	}
	if qRetweetRow == 1 {
		return p.UnQouteRetweet(ctx, &RetweetRequest{
			AccountID: AccountID,
			PostID:    PostID,
		})
	}

	postCount, err := p.postgre.GetRetweetRows(ctx, db.GetRetweetRowsParams{
		FromAccountID: AccountID,
		PostID:        PostID,
	})
	if err != nil {
		return err
	}
	if postCount == 1 {
		return p.UnRetweetPost(ctx, &RetweetRequest{
			AccountID: AccountID,
			PostID:    PostID,
		})
	}

	postInfo, err := p.postgre.GetPost(ctx, PostID)

	if postInfo.AccountID != AccountID {
		return errors.New("delete post must be owner")
	}

	err = p.postgre.PurgePost(ctx, PostID)
	if err != nil {
		return err
	}

	return nil
}
