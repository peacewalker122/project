package post

import (
	"context"
	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	"github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
)

type PostDBTX interface {
	CreatePostGCPTx(ctx context.Context, arg *request.CreatePostParams) (result.PostTXResult, error)
	CreateLike(ctx context.Context, postID uuid.UUID, accountID int64) (result.LikeTXResult, error)
	UnlikeTX(ctx context.Context, postID uuid.UUID, accountID int64) (result.LikeTXResult, error)
	CreateRetweetTX(ctx context.Context, arg *request.CreateRetweetParams) (result.RetweetTXResult, error)
	UnRetweet(ctx context.Context, PostID uuid.UUID, AccountID int64) error
	DeleteQouteRetweetTX(ctx context.Context, AccountID int64, PostID uuid.UUID) error
}

type PostTx struct {
	*db.SQLStore
}

func NewPostTx(db *db.SQLStore) *PostTx {
	return &PostTx{
		SQLStore: db,
	}
}
