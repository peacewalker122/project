package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/lib/pq"
	"github.com/peacewalker122/project/db/payload"
	"github.com/peacewalker122/project/db/redis"
	"github.com/peacewalker122/project/util"
)

type Store interface {
	Querier
	Model
	payload.Payload // we using this due tx not needed right now
}

type SQLStore struct {
	*Queries
	db *sql.DB
	payload.Payload
}

type NoSQLStore struct {
	redis.Store
}

func newTeststore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
		Payload: payload.NewPayload(db),
	}
}

func Newstore(db *sql.DB, Notif, RedisURL string) (Store, redis.Store) {
	return &SQLStore{
		Queries: New(db),
		db:      db,
		Payload:   payload.NewPayload(db),
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
	L   = "like"
	R   = "retweet"
	C   = "comment"
	QR  = "qoute-retweet"
	F   = "follow"
	UF  = "unfollow"
	UR  = "unretweet"
	UQR = "unqoute-retweet"
)

type Model interface {
	CreateUserTx(ctx context.Context, arg CreateUserParamsTx) (CreateUserTXResult, error)
	CreateAccountsQueueTX(ctx context.Context, arg CreateAccountQueueParams) (bool, error)
	DeleteQouteRetweetTX(ctx context.Context, arg UnRetweetTXParam) (int, error)
	CreateQouteRetweet(ctx context.Context, arg CreateQRetweetParams) (int, error)
	CreateQouteRetweetPostTX(ctx context.Context, arg CreateQRetweetParams) (CreateQRetweetResult, error)
	CreateCommentTX(ctx context.Context, arg CreateCommentParams) (CreateCommentTXResult, error)
	UnlikeTX(ctx context.Context, arg CreateLikeParams) (CreateLikeTXResult, error)
	CreateLikeTX(ctx context.Context, arg CreateLikeParams) (CreateLikeTXResult, error)
	DeleteRetweetTX(ctx context.Context, arg DeleteRetweetParams) error
	CreateRetweetPost(ctx context.Context, arg CreateRetweetParams) (CreateRetweetResult, error)
	CreateRetweetTX(ctx context.Context, arg CreateRetweetParams) (CreateRetweetTXResult, error)
	Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error)
	UnFollowtx(ctx context.Context, arg UnfollowTXParam) (UnFollowTXResult, error)
	GetDirectory(path string) (string, error)
	CreateFileIndex(path, filename string) (string, error)
	CreatePostTx(ctx context.Context, arg CreatePostParams) (PostTXResult, error)
}

type (
	CreateUserParamsTx struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}
	CreateQRetweetParams struct {
		FromAccountID int64  `json:"from_acc_id"`
		PostID        int64  `json:"post_id"`
		Qoute         string `json:"qoute"`
	}
	CreateCommentParams struct {
		FromAccountID int64  `json:"from_acc_id"`
		PostID        int64  `json:"post_id"`
		Comment       string `json:"comment"`
	}
	CreateLikeParams struct {
		FromAccountID int64 `json:"from_acc_id"`
		PostID        int64 `json:"post_id"`
	}
	CreateRetweetParams struct {
		FromAccountID int64 `json:"from_acc_id"`
		PostID        int64 `json:"post_id"`
		IsRetweet     bool  `json:"is_retweet"`
	}
	CreateRetweetResult struct {
		ErrCode int          `json:"err_code"`
		Post    PostTXResult `json:"post"`
	}
	CreateRetweetTXResult struct {
		ok      bool
		Err     error
		ErrCode int                 `json:"err_code"`
		Retweet CreateRetweetResult `json:"retweet_result"`
	}
	CreateAccountsQueueParams struct {
		Fromaccid int64 `json:"from_acc_id"`
		Toaccid   int64 `json:"to_acc_id"`
	}
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
	UnRetweetTXParam struct {
		FromAccountID int64 `json:"from_acc_id"`
		PostID        int64 `json:"post_id"`
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
	CreateLikeTXResult struct {
		PostFeature PostFeature `json:"post_feature"`
		ErrCode     int         `json:"err_code"`
	}
	CreateCommentTXResult struct {
		Comment     string      `json:"Comment"`
		PostFeature PostFeature `json:"post_feature"`
		ErrCode     int         `json:"err_code"`
	}
	CreateQRetweetResult struct {
		Qoute       string      `json:"qoute"`
		Post        Post        `json:"post"`
		PostFeature PostFeature `json:"post_feature"`
		ErrCode     int         `json:"err_code"`
	}
	CreateUserTXResult struct {
		User    User    `json:"user"`
		Account Account `json:"account"`
		Error   error   `json:"error"`
		ErrCode int     `json:"err_code"`
	}
)

func (s *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserParamsTx) (CreateUserTXResult, error) {
	var res CreateUserTXResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		res.User, err = q.CreateUser(ctx, CreateUserParams{
			Username:       arg.Username,
			HashedPassword: arg.Password,
			Email:          arg.Email,
			FullName:       arg.FullName,
		})
		if err != nil {
			if pqerr, ok := err.(*pq.Error); ok {
				switch pqerr.Code.Name() {
				case "unique_violation":
					res.ErrCode = http.StatusForbidden
					res.Error = errors.New("username or email already exists")
				}
			}
			res.ErrCode = http.StatusInternalServerError
			res.Error = err
		}

		res.Account, err = q.CreateAccounts(ctx, CreateAccountsParams{
			Owner: arg.Username,
		})
		if err != nil {
			if pqerr, ok := err.(*pq.Error); ok {
				switch pqerr.Code.Name() {
				case "unique_violation", "foreign_key_violation":
					res.ErrCode = http.StatusForbidden
					res.Error = err
				}
			}
			res.ErrCode = http.StatusInternalServerError
			res.Error = err
		}
		return err
	})
	return res, err
}

func (s *SQLStore) CreateAccountsQueueTX(ctx context.Context, arg CreateAccountQueueParams) (bool, error) {
	var result bool
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error
		// here we getting info did account is private or no
		ok, err := s.GetAccountsInfo(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		if ok.IsPrivate {
			res, err := s.CreatePrivateQueue(ctx, CreatePrivateQueueParams{
				Fromaccountid: arg.FromAccountID,
				Toaccountid:   arg.ToAccountID,
			})
			if err != nil {
				return err
			}
			result = res.Queue
			return err
		}
		result = false
		return err
	})

	return result, err
}

func (s *SQLStore) DeleteQouteRetweetTX(ctx context.Context, arg UnRetweetTXParam) (int, error) {
	var ErrCode int
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		num, err := s.GetPostQRetweetJoin(ctx, GetPostQRetweetJoinParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
		if err != nil {
			if err == sql.ErrNoRows {
				ErrCode = http.StatusNotFound
			}
			ErrCode = 500
			return err
		}

		_, err = s.GetQouteRetweet(ctx, GetQouteRetweetParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
		if err != nil {
			if err == sql.ErrNoRows {
				ErrCode = http.StatusNotFound
			}
			ErrCode = 500
			return err
		}

		_, err = s.CreateEntries(ctx, CreateEntriesParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			TypeEntries:   UQR,
		})
		if err != nil {
			return err
		}

		err = s.DeleteQouteRetweet(ctx, DeleteQouteRetweetParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
		if err != nil {
			return err
		}

		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				ErrCode = http.StatusNotFound
			}
			ErrCode = 500
			return err
		}

		// Delete first then decrement
		post.SumQouteRetweet--
		_, err = s.UpdatePost_feature(ctx, UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		err = s.DeletePostFeature(ctx, num.PostID)
		if err != nil {
			return err
		}
		err = s.DeletePost(ctx, num.PostID)
		if err != nil {
			return err
		}

		return err
	})
	return ErrCode, err
}

func (s *SQLStore) CreateCommentTX(ctx context.Context, arg CreateCommentParams) (CreateCommentTXResult, error) {
	var result CreateCommentTXResult

	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		result.Comment, err = s.CreateComment_feature(ctx, CreateComment_featureParams{FromAccountID: arg.FromAccountID, Comment: arg.Comment, PostID: arg.PostID})
		if err != nil {
			result.ErrCode = 500
			return err
		}

		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}
		post.SumComment++

		result.ErrCode = 500
		_, err = s.CreateEntries(ctx, CreateEntriesParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID, TypeEntries: arg.Comment})
		if err != nil {
			return err
		}

		result.PostFeature, err = s.UpdatePost_feature(ctx, UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		return err
	})
	return result, err
}

func (s *SQLStore) CreateQouteRetweet(ctx context.Context, arg CreateQRetweetParams) (int, error) {
	var ErrCode int
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				ErrCode = http.StatusNotFound
			}
			ErrCode = 500
			return err
		}
		post.SumQouteRetweet++

		err = s.UpdateQouteRetweet(ctx, UpdateQouteRetweetParams{
			QouteRetweet:  true,
			PostID:        arg.PostID,
			FromAccountID: arg.FromAccountID,
		})
		if err != nil {
			return err
		}

		args := CreateEntriesParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			TypeEntries:   UR,
		}

		_, err = s.CreateEntries(ctx, args)
		if err != nil {
			ErrCode = 500
			return err
		}

		arg := UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		}

		_, err = s.UpdatePost_feature(ctx, arg)
		if err != nil {
			ErrCode = 500
			return err
		}

		return err
	})
	return ErrCode, err
}

func (s *SQLStore) CreateQouteRetweetPostTX(ctx context.Context, arg CreateQRetweetParams) (CreateQRetweetResult, error) {
	var result CreateQRetweetResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error
		result.ErrCode = http.StatusInternalServerError

		result.Qoute, err = s.CreateQouteRetweet_feature(ctx, CreateQouteRetweet_featureParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			Qoute:         arg.Qoute,
		})
		if err != nil {
			return err
		}

		ok, err := s.GetPostQRetweetJoin(ctx, GetPostQRetweetJoinParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		if ok.QouteRetweet {
			return errors.New("already created")
		}

		post, err := s.GetPost(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		arg := CreatePostParams{
			AccountID: arg.FromAccountID,
			IsRetweet: true,
			PhotoDir:  util.InputSqlString(post.PhotoDir.String),
		}
		result.Post, err = s.CreatePost(ctx, arg)
		if err != nil {
			return err
		}

		result.PostFeature, err = s.CreatePost_feature(ctx, result.Post.PostID)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (s *SQLStore) CreateRetweetTX(ctx context.Context, arg CreateRetweetParams) (CreateRetweetTXResult, error) {
	var result CreateRetweetTXResult
	result.ErrCode = http.StatusInternalServerError
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		result.ok, err = s.GetRetweetJoin(ctx, GetRetweetJoinParams{Postid: arg.PostID, Fromaccountid: arg.FromAccountID})
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		if result.ok && arg.IsRetweet {
			result.ErrCode = http.StatusNotFound
			result.Err = errors.New("already retweet")
			return result.Err
		}

		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			result.ErrCode = 500
			return err
		}
		post.SumRetweet++

		err = s.UpdateRetweet(ctx, UpdateRetweetParams{Retweet: true, PostID: arg.PostID, FromAccountID: arg.FromAccountID})
		if err != nil {
			result.ErrCode = 500
			return err
		}

		args := CreateEntriesParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			TypeEntries:   R,
		}

		_, err = s.CreateEntries(ctx, args)
		if err != nil {
			result.ErrCode = 500
			return err
		}

		arg := UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		}

		_, err = s.UpdatePost_feature(ctx, arg)
		if err != nil {
			result.ErrCode = 500
			return err
		}
		return err
	})
	// if result.Err != nil {
	// 	err = result.Err
	// }
	return result, err
}

func (s *SQLStore) UnlikeTX(ctx context.Context, arg CreateLikeParams) (CreateLikeTXResult, error) {
	var res CreateLikeTXResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err != sql.ErrNoRows {
				res.ErrCode = http.StatusNotFound
			}
			return err
		}
		post.SumLike--

		err = s.UpdateLike(ctx, UpdateLikeParams{
			IsLike:        false,
			PostID:        arg.PostID,
			FromAccountID: arg.FromAccountID,
		})
		if err != nil {
			res.ErrCode = 500
			return err
		}

		_, err = s.CreateEntries(ctx, CreateEntriesParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			TypeEntries:   L,
		})
		if err != nil {
			return err
		}

		res.PostFeature, err = s.UpdatePost_feature(ctx, UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		return err

	})
	return res, err
}

func (s *SQLStore) CreateLikeTX(ctx context.Context, arg CreateLikeParams) (CreateLikeTXResult, error) {
	var res CreateLikeTXResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			if err != sql.ErrNoRows {
				res.ErrCode = http.StatusNotFound
			}
			return err
		}
		post.SumLike++

		err = s.UpdateLike(ctx, UpdateLikeParams{
			IsLike:        true,
			PostID:        arg.PostID,
			FromAccountID: arg.FromAccountID,
		})
		if err != nil {
			res.ErrCode = 500
			return err
		}

		_, err = s.CreateEntries(ctx, CreateEntriesParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			TypeEntries:   L,
		})
		if err != nil {
			return err
		}

		res.PostFeature, err = s.UpdatePost_feature(ctx, UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		return err
	})
	return res, err
}

func (s *SQLStore) DeleteRetweetTX(ctx context.Context, arg DeleteRetweetParams) error {
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error

		res, err := s.GetPostidretweetJoin(ctx, GetPostidretweetJoinParams{
			PostID: arg.PostID, FromAccountID: arg.FromAccountID,
		})
		if err != nil {
			return err
		}
		post, err := s.GetPost_feature_Update(ctx, arg.PostID)
		if err != nil {
			return err
		}
		_, err = s.GetRetweet(ctx, GetRetweetParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("no specify qoute-retweet in database")
			}
			return err
		}

		_, err = s.CreateEntries(ctx, CreateEntriesParams{
			FromAccountID: arg.FromAccountID,
			PostID:        arg.PostID,
			TypeEntries:   UR,
		})
		if err != nil {
			return err
		}

		err = s.DeleteRetweet(ctx, DeleteRetweetParams{PostID: arg.PostID, FromAccountID: arg.FromAccountID})
		if err != nil {
			return err
		}

		// Delete first then decrement
		post.SumRetweet--
		_, err = s.UpdatePost_feature(ctx, UpdatePost_featureParams{
			PostID:          arg.PostID,
			SumComment:      post.SumComment,
			SumLike:         post.SumLike,
			SumRetweet:      post.SumRetweet,
			SumQouteRetweet: post.SumQouteRetweet,
		})
		if err != nil {
			return err
		}

		err = s.DeletePostFeature(ctx, res.PostID)
		if err != nil {
			return err
		}
		err = s.DeletePost(ctx, res.PostID)
		if err != nil {
			return err
		}
		return err
	})
	return err
}

func (s *SQLStore) CreateRetweetPost(ctx context.Context, arg CreateRetweetParams) (CreateRetweetResult, error) {
	var result CreateRetweetResult
	err := s.execCtx(ctx, func(q *Queries) error {
		var err error
		result.ErrCode = http.StatusInternalServerError

		err = s.CreateRetweet_feature(ctx, CreateRetweet_featureParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
		if err != nil {
			return err
		}

		ok, err := s.GetPostidretweetJoin(ctx, GetPostidretweetJoinParams{FromAccountID: arg.FromAccountID, PostID: arg.PostID})
		if err != nil {
			if err == sql.ErrNoRows {
				result.ErrCode = http.StatusNotFound
			}
			return err
		}

		if !ok.Retweet {

			post, err := s.GetPost(ctx, arg.PostID)
			if err != nil {
				if err == sql.ErrNoRows {
					result.ErrCode = http.StatusNotFound
				}
				return err
			}

			arg := CreatePostParams{
				AccountID: arg.FromAccountID,
				IsRetweet: true,
				PhotoDir:  util.InputSqlString(post.PhotoDir.String),
			}
			result.Post.Post, err = s.CreatePost(ctx, arg)
			if err != nil {
				return err
			}

			result.Post.PostFeature, err = s.CreatePost_feature(ctx, result.Post.Post.PostID)
			if err != nil {
				return err
			}

		}
		return err
	})

	return result, err
}

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

		if arg.IsQueue {
			err = s.DeleteAccountQueue(ctx, DeleteAccountQueueParams{
				Fromaccountid: arg.Fromaccid,
				Toaccountid:   arg.Toaccid,
			})
			if err != nil {
				return err
			}
		}

		result.Follow, result.ToAcc, result.FromAcc, err = s.UpdateFollowing(ctx, arg.Fromaccid, arg.Toaccid, int64(1))
		if err != nil {
			return err
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
