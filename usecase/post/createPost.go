package post

import (
	"context"
	"errors"
	request "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/request/post"
	result "github.com/peacewalker122/project/service/db/repository/postgres/sqlc/result/post"
	request2 "github.com/peacewalker122/project/service/gcp/request"
)

func (p *PostUsecase) CreatePost(ctx context.Context, param *CreatePostRequest) (*result.PostTXResult, error) {
	if param == nil {
		return nil, errors.New("param is nil")
	}

	postData, err := p.postgre.CreatePostGCPTx(ctx, &request.CreatePostParams{
		AccountID:          param.AccountID,
		IsRetweet:          false,
		PictureDescription: param.PictureDescription,
		FileRequest: &request2.UploadFilesRequest{
			File:          param.File,
			FileHeader:    param.FileHeader,
			BucketName:    p.config.BucketName,
			MaxUploadSize: p.config.MaxUploadSizeInBytes(),
		},
		GcpFunc: p.gcp.UploadPhoto,
	})
	if err != nil {
		return nil, err
	}

	return &postData, nil
}
