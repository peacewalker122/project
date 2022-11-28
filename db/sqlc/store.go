package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Store interface {
	Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error)
	UnFollowtx(ctx context.Context, arg UnfollowTXParam) (UnFollowTXResult, error)
	GetDirectory(path string) (string, error)
	CreateFileIndex(path, filename string) (string, error)
	Querier
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func Newstore(db *sql.DB, bucketType string) Store {
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
	F  = "follow"
	UF = "unfollow"
)

type (
	FollowTXParam struct {
		Fromaccid int64 `json:"from_acc_id"`
		Toaccid   int64 `json:"to_acc_id"`
	}
	FollowTXResult struct {
		Follow      AccountsFollow `json:"account_follow"`
		FeatureType string         `json:"feature_type"`
		FromAcc     Account        `json:"from_acc"`
		ToAcc       Account        `json:"to_acc"`
	}
	UnfollowTXParam struct {
		Fromaccid int64 `json:"from_acc_id"`
		Toaccid   int64 `json:"to_acc_id"`
	}
	UnFollowTXResult struct {
		Status      bool    `json:"status"`
		FeatureType string  `json:"feature_type"`
		FromAcc     Account `json:"from_acc"`
		ToAcc       Account `json:"to_acc"`
	}
)

func (q *Queries) Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error) {
	var result FollowTXResult
	var err error
	result.FeatureType = F

	result.Follow, result.ToAcc, result.FromAcc, err = q.UpdateFollowing(ctx, arg.Fromaccid, arg.Toaccid, int64(1))
	if err != nil {
		return result, err
	}

	return result, nil
}
func (q *Queries) UnFollowtx(ctx context.Context, arg UnfollowTXParam) (UnFollowTXResult, error) {
	var result UnFollowTXResult
	var err error
	result.FeatureType = UF

	result.ToAcc, result.FromAcc, result.Status, err = q.DeleteFollowing(ctx, arg.Fromaccid, arg.Toaccid, -int64(1))
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *Queries) UpdateFollowing(
	ctx context.Context, fromAccount, toAccount, num int64,
) (
	acc AccountsFollow, Toacc, Fromacc Account, err error,
) {
	acc, err = q.CreateAccountsFollow(ctx, CreateAccountsFollowParams{
		FromAccountID: fromAccount,
		ToAccountID:   toAccount,
		Follow:        true,
	})
	if err != nil {
		return
	}
	Toacc, err = q.UpdateAccountFollower(ctx, UpdateAccountFollowerParams{Num: num, AccountsID: toAccount})
	if err != nil {
		return
	}
	Fromacc, err = q.UpdateAccountFollowing(ctx, UpdateAccountFollowingParams{Num: num, AccountsID: fromAccount})
	if err != nil {
		return
	}
	return
}

func (q *Queries) DeleteFollowing(
	ctx context.Context, fromAccount, toAccount, num int64,
) (
	Toacc, Fromacc Account, status bool, err error,
) {
	Toacc, err = q.UpdateAccountFollower(ctx, UpdateAccountFollowerParams{Num: num, AccountsID: toAccount})
	if err != nil {
		return
	}
	Fromacc, err = q.UpdateAccountFollowing(ctx, UpdateAccountFollowingParams{Num: num, AccountsID: fromAccount})
	if err != nil {
		return
	}

	err = q.DeleteAccountsFollow(ctx, DeleteAccountsFollowParams{
		Fromid: fromAccount,
		Toid:   toAccount,
	})
	if err != nil {
		return
	}
	status = true

	return
}

func (q *Queries) GetDirectory(path string) (string, error) {
	res, err := exec.Command(path).Output()
	if err != nil {
		return err.Error(), err
	}
	return string(res), err
}

// creating new file index if its already exist.
// in linux using pwd command in your terminal, then paste it to the path parameter.
func (q *Queries) CreateFileIndex(path, filename string) (string, error) {
	if path == "" {
		return "", errors.New("empty string")
	}
	// Before for the strings before the separator strings and vice versa
	// add before with (n) n start from 1
	n := 1
	before, after, _ := strings.Cut(filename, ".")
	before = before + fmt.Sprintf("(%v)", n)
	result := before + "." + after
	if _, err := os.Stat(path + result); err == nil {
		result, _ = validatingfile(result)
	}
	return result, nil
}

func validatingfile(filename string) (string, error) {
	n := 1
	before, after, _ := strings.Cut(filename, ".")
	before = before + fmt.Sprintf("(%v)", n)

	s := before + "." + after

	return s, nil
}
