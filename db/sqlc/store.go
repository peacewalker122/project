package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/peacewalker122/project/db/redis"
)

type Store interface {
	Querier
	Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error)
	UnFollowtx(ctx context.Context, arg UnfollowTXParam) (UnFollowTXResult, error)
	GetDirectory(path string) (string, error)
	CreateFileIndex(path, filename string) (string, error)
	CreatePostTx(ctx context.Context, arg CreatePostParams) (PostTXResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

type NoSQLStore struct {
	redis.Store
}

func Newstore(db *sql.DB, RedisURL string) (Store, redis.Store) {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}, &NoSQLStore{Store: redis.NewRedis(RedisURL)}
}

func (s *SQLStore) execCtx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rBerr := tx.Rollback(); rBerr != nil {
			return fmt.Errorf("tx error %v, rb error %v", err, rBerr)
		}
	}
	return tx.Commit()
}

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
		IsQueue   bool  `json:"is_queue"`
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
	PostTXResult struct {
		Post        Post        `json:"post"`
		PostFeature PostFeature `json:"post_feature"`
	}
)

func (s *SQLStore) CreatePostTx(ctx context.Context, arg CreatePostParams) (PostTXResult, error) {
	var result PostTXResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		result.Post, err = s.CreatePost(ctx, arg)
		if err != nil {
			return err
		}

		result.PostFeature, err = s.CreatePost_feature(ctx, result.Post.PostID)
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

func (s *SQLStore) Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error) {
	var result FollowTXResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error
		result.FeatureType = F

		result.Follow, result.ToAcc, result.FromAcc, err = s.UpdateFollowing(ctx, arg.Fromaccid, arg.Toaccid, int64(1))
		if err != nil {
			return err
		}

		if arg.IsQueue {
			err = s.DeleteAccountsFollow(ctx, DeleteAccountsFollowParams{
				Fromid: arg.Fromaccid,
				Toid:   arg.Toaccid,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	return result, err
}
func (s *SQLStore) UnFollowtx(ctx context.Context, arg UnfollowTXParam) (UnFollowTXResult, error) {
	var result UnFollowTXResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error
		result.FeatureType = UF

		result.ToAcc, result.FromAcc, result.Status, err = s.DeleteFollowing(ctx, arg.Fromaccid, arg.Toaccid, -int64(1))
		if err != nil {
			return err
		}

		return nil
	})
	return result, err
}

func (q *SQLStore) UpdateFollowing(
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

func (q *SQLStore) DeleteFollowing(
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
