package post

import (
	"context"
	"database/sql"
	"github.com/peacewalker122/project/service/db/repository/postgres/payload/model"
	"github.com/peacewalker122/project/util"

	"github.com/google/uuid"
	db "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/generate"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
)

type PostDBTX interface {
	CreatePostGCPTx(ctx context.Context, arg *request.CreatePostParams) (result.PostTXResult, error)
	CreateLike(ctx context.Context, postID uuid.UUID, accountID int64) (result.LikeTXResult, error)
	UnlikeTX(ctx context.Context, postID uuid.UUID, accountID int64) (result.LikeTXResult, error)
	CreateRetweetTX(ctx context.Context, arg *request.CreateRetweetParams) (result.RetweetTXResult, error)
	DeleteRetweetTX(ctx context.Context, PostID uuid.UUID, AccountID int64) error
	CreateQouteRetweet(ctx context.Context, arg *request.CreateQouteRetweetParams) (*uuid.UUID, error)
	DeleteQouteRetweetTX(ctx context.Context, AccountID int64, PostID uuid.UUID) error
	CreateComment(ctx context.Context, param *request.CreateCommentParams) *util.MultiError
	PurgePost(ctx context.Context, PostID uuid.UUID) error
}

type PostTx struct {
	*db.SQLStore
	ent model.Model
}

func NewPostTx(db *db.SQLStore, projectDB *sql.DB) *PostTx {
	return &PostTx{
		SQLStore: db,
		ent:      model.NewModel(projectDB),
	}
}
