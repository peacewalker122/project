package request

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/peacewalker122/project/service/gcp/request"
)

type CreatePostParams struct {
	AccountID          int64  `json:"account_id"`
	IsRetweet          bool   `json:"is_retweet"`
	PictureDescription string `json:"picture_description"`
	FileRequest        *request.UploadFilesRequest
	GcpFunc            func(ctx context.Context, req *request.UploadFilesRequest) (string, error)
}

func (p *CreatePostParams) Validate() error {
	if p.AccountID == 0 {
		return errors.New("account_id is required")
	}
	if p.FileRequest == nil {
		return errors.New("file_request is required")
	}
	if err := p.FileRequest.Validate(); err != nil {
		return err
	}
	if p.GcpFunc == nil {
		return errors.New("gcp_func is required")
	}
	return nil
}

type CreateRetweetParams struct {
	AccountID int64     `json:"account_id"`
	PostID    uuid.UUID `json:"post_id"`
}

type CreateCommentParams struct {
	AccountID int64     `json:"account_id"`
	PostID    uuid.UUID `json:"post_id"`
	Comment   string    `json:"comment"`
}

type CreateQouteRetweetParams struct {
	AccountID int64     `json:"account_id"`
	PostID    uuid.UUID `json:"post_id"`
	Qoute     string    `json:"qoute"`
}
