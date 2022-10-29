package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error)
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

// TO BE IMPLEMENTED IF TX NEEDED
// func (s *SQLStore) execCtx(ctx context.Context, fn func(q *Queries) error) error {
// 	tx, err := s.db.BeginTx(ctx, nil)
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

const (
	L  = "like"
	R  = "retweet"
	C  = "comment"
	QR = "qoute-retweet"
	F  = "Follow"
)

type FollowTXParam struct {
	Fromaccid int64 `json:"from_acc_id"`
	Toaccid   int64 `json:"to_acc_id"`
}

type FollowTXResult struct {
	Follow      AccountsFollow `json:"account_follow"`
	FeatureType string         `json:"feature_type"`
	FromAcc     Account        `json:"from_acc"`
	ToAcc       Account        `json:"to_acc"`
}

func (q *Queries) Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error) {
	var result FollowTXResult
	var err error
	result.FeatureType = F
	
	result.Follow, result.ToAcc, result.FromAcc, err = q.AddFollowing(ctx, arg.Fromaccid, arg.Toaccid)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *Queries) AddFollowing(ctx context.Context, fromAccount, toAccount int64) (acc AccountsFollow, Toacc, Fromacc Account, err error) {
	acc, err = q.CreateAccountsFollow(ctx, CreateAccountsFollowParams{
		FromAccountID: fromAccount,
		ToAccountID:   toAccount,
		Follow:        true,
	})
	if err != nil {
		return
	}
	Toacc, err = q.AddAccountFollower(ctx, toAccount)
	if err != nil {
		return
	}
	Fromacc, err = q.AddAccountFollowing(ctx, fromAccount)
	if err != nil {
		return
	}
	return
}
