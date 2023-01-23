package post

import (
	"context"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
	post2 "github.com/peacewalker122/project/usecase/post"
	"github.com/peacewalker122/project/util"
)

type PostContract interface {
	CreatePost(ctx context.Context, param *post2.CreatePostRequest) (*result.PostTXResult, error)
	LikePost(ctx context.Context, param *post2.LikeRequest) error
	UnLikePost(ctx context.Context, param *post2.LikeRequest) error
	RetweetPost(ctx context.Context, param *post2.RetweetRequest) error
	CreateQouteRetweet(ctx context.Context, param *post2.QouteRetweetRequest) error
	CreateComment(ctx context.Context, param *post2.CommentRequest) (err *util.MultiError)
}
