package post

import (
	"context"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
	post2 "github.com/peacewalker122/project/usecase/post"
)

type PostContract interface {
	CreatePost(ctx context.Context, param *post2.CreatePostRequest) (*result.PostTXResult, error)
}
