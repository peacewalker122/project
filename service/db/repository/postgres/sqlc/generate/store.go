package db

import (
	"context"
	"database/sql"
	"fmt"
)

type SQLStore struct {
	*Queries
	DB *sql.DB
}

func newTeststore(db *sql.DB) *SQLStore {
	return &SQLStore{
		Queries: New(db),
		DB:      db,
	}
}

func NewStore(projectDB *sql.DB) *SQLStore {
	sqlStore := &SQLStore{
		Queries: New(projectDB),
		DB:      projectDB, // the first db is for notif and the second is for the main db
	}

	return sqlStore
}

func (s *SQLStore) execCtx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := s.DB.BeginTx(ctx, nil)
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

func (s *SQLStore) DBTx(ctx context.Context, fn func(q *Queries) error) error {
	tx, err := s.DB.BeginTx(ctx, nil)
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
	UL  = "unlike"
	R   = "retweet"
	C   = "comment"
	QR  = "qoute-retweet"
	F   = "follow"
	UF  = "unfollow"
	UR  = "unretweet"
	UQR = "unqoute-retweet"
)

type Model interface {
	CreateAccountsQueueTX(ctx context.Context, arg CreateAccountQueueParams) (bool, error)
	//CreateUserTx(ctx context.Context, arg CreateUserParamsTx) (CreateUserTXResult, error)
	//Followtx(ctx context.Context, arg FollowTXParam) (FollowTXResult, error)
	//UnFollowtx(ctx context.Context, arg UnfollowTXParam) (UnFollowTXResult, error)
	//DeleteQouteRetweetTX(ctx context.Context, arg UnRetweetTXParam) (int, error)
	//CreateQouteRetweet(ctx context.Context, arg CreateQRetweetParams) (int, error)
	//CreateQouteRetweetPostTX(ctx context.Context, arg CreateQRetweetParams) (CreateQRetweetResult, error)
	//CreateCommentTX(ctx context.Context, arg CreateCommentParams) (CreateCommentTXResult, error)
	//UnlikeTX(ctx context.Context, arg CreateLikeParams) (CreateLikeTXResult, error)
	//CreateLikeTX(ctx context.Context, arg CreateLikeParams) (CreateLikeTXResult, error)
	//CreatePostTx(ctx context.Context, arg CreatePostParams) (PostTXResult, error)
	//DeleteRetweetTX(ctx context.Context, arg DeleteRetweetParams) error
	//CreateRetweetPost(ctx context.Context, arg CreateRetweetParams) (CreateRetweetResult, error)
	//CreateRetweetTX(ctx context.Context, arg CreateRetweetParams) (CreateRetweetTXResult, error)
}


type (
	CreateRetweetResult struct {
		ErrCode int          `json:"err_code"`
		Post    PostTXResult `json:"post"`
	}
	PostTXResult struct {
		Post        Post        `json:"post"`
		PostFeature PostFeature `json:"post_feature"`
	}
)

func (s *SQLStore) CreateAccountsQueueTX(ctx context.Context, arg CreateAccountQueueParams) (bool, error) {
	var result bool

	var err error

	ok, err := s.GetAccountsInfo(ctx, arg.ToAccountID)
	if err != nil {
		return result, err
	}

	if ok.IsPrivate {
		res, err := s.CreatePrivateQueue(ctx, CreatePrivateQueueParams{
			Fromaccountid: arg.FromAccountID,
			Toaccountid:   arg.ToAccountID,
		})
		if err != nil {
			return result, err
		}
		result = res.Queue
		return result, err
	}
	result = false

	return result, err
}
