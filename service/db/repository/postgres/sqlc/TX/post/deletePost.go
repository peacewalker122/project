package post

import (
	"context"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"log"
	"sync"
)

func (p *PostTx) DeletePostTX(ctx context.Context, PostID uuid.UUID) error {
	var err error

	err = p.DeletePostFeature(ctx, PostID)
	if err != nil {
		return err
	}

	err = p.DeletePost(ctx, PostID)
	if err != nil {
		return err
	}

	return err
}

func (p *PostTx) PurgePost(ctx context.Context, PostID uuid.UUID) error {
	err := p.DBTx(ctx, func(q *db.Queries) error {
		var (
			err error
			wg  sync.WaitGroup
		)
		chanerr := make(chan error, 1)

		retweets, err := p.ent.GetRetweetByPost(ctx, PostID)
		if err != nil {
			return err
		}

		qretweets, err := p.ent.GetQouteRetweetByPost(ctx, PostID)
		if err != nil {
			return err
		}

		likes, err := p.ent.GetLikesByPost(ctx, PostID)
		if err != nil {
			return err
		}

		wg.Add(3)
		go func() {
			defer wg.Done()
			for _, like := range likes {
				err = p.ent.DeleteLike(ctx, like.PostID, like.FromAccountID)
				if err != nil {
					chanerr <- err
					return
				}
			}
		}()

		go func() {
			defer wg.Done()
			for _, retweet := range retweets {
				err = q.DeleteRetweet(ctx, db.DeleteRetweetParams{
					PostID:        PostID,
					FromAccountID: retweet.FromAccountID,
				})
				if err != nil {
					chanerr <- err
					return
				}
				err = p.DeletePost(ctx, retweet.PostID)
				if err != nil {
					chanerr <- err
				}
			}
		}()

		go func() {
			defer wg.Done()
			for _, qretweet := range qretweets {
				err = q.DeleteQouteRetweet(ctx, db.DeleteQouteRetweetParams{
					PostID:        PostID,
					FromAccountID: qretweet.FromAccountID,
				})
				if err != nil {
					chanerr <- err
				}
				err = p.DeletePost(ctx, qretweet.PostID)
				if err != nil {
					chanerr <- err
				}
			}
		}()
		wg.Wait()
		select {
		case err = <-chanerr:
			log.Println("Error: ", err)
			return err
		default:
		}

		err = p.DeletePost(ctx, PostID)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
