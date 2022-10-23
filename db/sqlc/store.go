package db

import (
	"database/sql"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func Newstore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// func (s *SQLStore) execCtx(ctx context.Context, fn func(q *Queries) error) error {
// 	tx, err := s.db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rBerr := tx.Rollback(); rBerr != nil {
// 			return fmt.Errorf("tx error %v, rb error %v", err, rBerr)
// 		}
// 	}
// 	return tx.Commit()
// }

// const (
// 	L  = "like"
// 	R  = "retweet"
// 	C  = "comment"
// 	QR = "qoute-retweet"
// )

// type PostTXParam struct {
// 	PostID        int64  `json:"post_id"`
// 	FeatureType   string `json:"feature_type"`
// }

// type PostTXResult struct {
// 	PostFeature   PostFeature `json:"post_feature"`
// 	PostID        int64       `json:"post_id"`
// 	FeatureType   string      `json:"feature_type"`
// 	Entry         Entry       `json:"entry"`
// }

// func (s *SQLStore) PostTX(ctx context.Context, arg PostTXParam) (PostTXResult, error) {
// 	var result PostTXResult

// 	err := s.execCtx(ctx, func(q *Queries) error {
// 		var err error

// 		numcomment := 0
// 		numlike := 0
// 		numretweet := 0
// 		numQretweet := 0

// 		switch arg.FeatureType {
// 		case L:
// 			numlike++
// 		case R:
// 			numretweet++
// 		case C:
// 			numcomment++
// 		case QR:
// 			numQretweet++
// 		}

// 		result.PostFeature, err = q.CreatePost_feature(ctx, CreatePost_featureParams{
// 			PostID:          arg.PostID,
// 			SumComment:      int64(numcomment),
// 			SumLike:         int64(numlike),
// 			SumRetweet:      int64(numretweet),
// 			SumQouteRetweet: int64(numQretweet),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		result.Entry, err = q.CreateEntries(ctx, CreateEntriesParams{
// 			PostID:        arg.PostID,
// 			TypeEntries:   arg.FeatureType,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// 	return result, err
// }
